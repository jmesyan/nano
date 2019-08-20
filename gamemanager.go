package nano

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/game"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"net"
	"reflect"
	"sort"
	"time"
)

var (
	GMHander         *GameManager
	DefautListenGame = "0.0.0.0:20572"
)

type GameManager struct {
	listenaddrs      string
	natsaddrs        string
	Serversort       map[string]*game.GameServer
	Alltablesort     map[string]*game.GameTable
	enterMaxConnects int
	client           *nats.Conn
	msgch            chan *nats.Msg
	shut             chan struct{}
	supTopic         string
	sdownTopic       string
	tbregTopic       string
}

type GameManagerOpts func(g *GameManager)

func WithGameManagerAddrs(addrs string) GameManagerOpts {
	return func(g *GameManager) {
		g.listenaddrs = addrs
	}
}

func NewGameManager(opts ...GameManagerOpts) *GameManager {
	g := &GameManager{
		listenaddrs:      DefautListenGame,
		natsaddrs:        nats.DefaultURL,
		Serversort:       make(map[string]*game.GameServer),
		Alltablesort:     make(map[string]*game.GameTable),
		msgch:            make(chan *nats.Msg, 64),
		shut:             make(chan struct{}, 1),
		enterMaxConnects: 0,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(g)
		}
	}
	return g
}
func (g *GameManager) Init() {
	g.InitNats()
	go g.handleConn()
}
func (g *GameManager) handleConn() {
	listen, err := net.Listen("tcp", g.listenaddrs)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}

		// start a new goroutine to handle the new connection
		game.NewGameServer(conn, g)
	}
}

func (g *GameManager) InitNats() {
	var err error
	g.client, err = nats.Connect(g.natsaddrs)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = g.client.ChanSubscribe("GameMaster.>", g.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	//设置topic
	g.supTopic = "GameMaster.sup"
	g.sdownTopic = "GameMaster.sdown"
	g.tbregTopic = "GameMaster.tbreg"
	//监听消息
	go g.watcher()
}

func (g *GameManager) watcher() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			g.ServerMaintence()
		case msg := <-g.msgch:
			g.HandleMsg(msg)
		case <-g.shut:
			logger.Println("receive stop msg")
			close(g.msgch)
			return
		}
	}
}

func (g *GameManager) ServerMaintence() {
	list := dcm.GetGameServerNodes()
	nowtime := utils.Time()
	for nid, node := range list {
		gsid := node.Gsid
		if _, ok := g.Serversort[gsid]; !ok {
			logger.Println("register remote server:", gsid)
			gid, rtype, ridx := game.GetGameParamsByGsid(gsid)
			server := &game.GameServer{
				Node:       node,
				Natsaddrs:  nats.DefaultURL,
				Status:     2,
				Tablesort:  nil,
				Gsid:       gsid,
				Gid:        gid,
				Rtype:      rtype,
				Ridx:       ridx,
				StartTime:  nowtime,
				C2sTopic:   fmt.Sprintf("%s.c2s", nid),
				SsTopic:    fmt.Sprintf("%s.ss", nid),
				C2sDestory: fmt.Sprintf("%s.channel.destory", nid),
				IsRemote:   true,
			}
			server.SetClient(g.client)
			g.Serversort[gsid] = server
		}
	}
}

func (g *GameManager) HandleMsg(msg *nats.Msg) {
	logger.Printf("handle gamemanager nats msg:%#v\n", msg)
	switch msg.Subject {
	case g.supTopic:
		var server *game.GameServer
		err := json.Unmarshal(msg.Data, &server)
		if err != nil {
			logger.Println(err)
			return
		}
		logger.Println("register remote server2:", server.Gsid)
		if _, ok := g.Serversort[server.Gsid]; !ok {
			server.IsRemote = true
			g.Serversort[server.Gsid] = server
		}

	case g.sdownTopic:
		gsid := string(msg.Data)
		if _, ok := g.Serversort[gsid]; ok {
			delete(g.Serversort, gsid)
		}
	case g.tbregTopic:
		var table *game.GameTable
		err := json.Unmarshal(msg.Data, &table)
		if err != nil {
			logger.Println(err)
			return
		}
		if _, ok := g.Alltablesort[table.Gsidtid]; !ok {
			g.Alltablesort[table.Gsidtid] = table
		}
	}
}
func (g *GameManager) AfterInit() {

}
func (g *GameManager) BeforeShutdown() {
	g.shut <- struct{}{}
}
func (g *GameManager) Shutdown() {

}

func (g *GameManager) ProcessServer(route string, body reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			logger.Println(fmt.Sprintf("ProcessServer err: %v", err))
			println(utils.Stack())
		}
	}()
	if mt, ok := handler.srvhandlers[route]; ok {
		mt.Method.Func.Call([]reflect.Value{mt.Receiver, body})
	} else {
		fmt.Printf("the is no srv handler, route is:%s, body is:%#v", route, body)
	}
}

func (g *GameManager) RegisterTable(gsidtid string, table *game.GameTable) {
	g.Alltablesort[gsidtid] = table
	data, err := json.Marshal(table)
	if err != nil {
		logger.Println("RegisterTable marshal error:", err.Error())
		return
	}
	err = g.client.Publish(g.tbregTopic, data)
	if err != nil {
		logger.Println("RegisterTable publish error:", err.Error())
	}
}

func (g *GameManager) RegisterServer(gsid string, server *game.GameServer) {
	g.Serversort[gsid] = server
	data, err := json.Marshal(server)
	if err != nil {
		logger.Println("RegisterServer marshal error:", err.Error())
		return
	}
	err = g.client.Publish(g.supTopic, data)
	if err != nil {
		logger.Println("RegisterServer publish error:", err.Error())
	}
}

func (g *GameManager) GetServerSort() map[string]*game.GameServer {
	return g.Serversort
}

func (g *GameManager) GetServerByGSID(gsid string) *game.GameServer {
	if server, ok := g.Serversort[gsid]; ok {
		return server
	}
	return nil
}

func (g *GameManager) RemoveServerByGSID(gsid string) {
	delete(g.Serversort, gsid)
	err := g.client.Publish(g.sdownTopic, []byte(gsid))
	if err != nil {
		logger.Printf("RemoveServerByGSID,gsid:%s, err:%s", gsid, err.Error())
	}
}

func (g *GameManager) GetCenterServers(ngid int, ngc func(s *game.GameServer) bool) map[string]*game.GameServer {
	config, ok := gds.Configs[ngid]
	if !ok {
		return nil
	}
	gcid := config.Censerver
	gsids := make(map[string]*game.GameServer)
	for gsid, server := range g.Serversort {
		gid, rtype, _ := game.GetGameParamsByGsid(gsid)
		grid := game.GetGrid(gid, rtype)
		if grid == gcid && ngc(server) {
			gsids[gsid] = server
		}
	}
	return gsids
}

func (g *GameManager) GetCenterServerByBalance(ngid int) *game.GameServer {
	config, ok := gds.Configs[ngid]
	if !ok {
		return nil
	}
	gcid := config.Censerver
	gsids := make(map[string]int)
	for gsid, _ := range g.Serversort {
		gid, rtype, _ := game.GetGameParamsByGsid(gsid)
		grid := game.GetGrid(gid, rtype)
		if grid == gcid && !game.IsServerMaintence(gsid) {
			gsids[gsid] = gds.Gcsu[gsid]
		}
	}
	if len(gsids) > 0 {
		var i, s, n = 0, make([]string, len(gsids)), make([]int, len(gsids))
		for gsid, num := range gsids {
			s[i], n[i] = gsid, num
			i += 1
		}
		sort.Slice(s, func(i, j int) bool {
			return n[i] < n[j]
		})
		return g.Serversort[s[0]]
	}
	return nil
}

func (g *GameManager) ReconnectToGame(uid int, connectServerdata *game.ServerData) (*game.ServerData, error) {
	var serverdata *game.ServerData
	channel := game.GetChannel(uid)
	if channel != nil {
		serverdata = channel.SrvData
		logger.Println("reconnectToGame logoutgame:", uid)
		err := channel.LogoutGame(true)
		if err != nil {
			return nil, err
		}
	}
	if serverdata != nil {
		if connectServerdata != nil && serverdata.Gsidtid == connectServerdata.Gsidtid {
			return nil, errors.New(fmt.Sprintf("same user request to game:%#v", connectServerdata))
		}
		u, err := game.UMHandler.Member(uid)
		if err != nil {
			return nil, err
		}
		err = u.Push("reconnect", serverdata)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (g *GameManager) EnterToGame(uid int, serverdata *game.ServerData, cb func(result *game.ControlUserEnterroom), isretry bool) error {
	isreconnect, err := g.ReconnectToGame(uid, serverdata)
	if err != nil {
		return err
	}
	if isreconnect != nil {
		serverdata = isreconnect
	}
	server := g.GetServerByGSID(serverdata.Gsid)
	if server == nil {
		return errors.New(fmt.Sprintf("serverNotOnline,uid:%d, sid:%s", uid, serverdata.Gsid))
	}
	channel := game.NewGameChannel(uid, game.ConnectorHandler.NID(), game.ConnectorHandler.GetClient(), g)
	channel.SetGameNid(server.NID(), server.GetNode().Address)
	t := time.Now()
	callback := func(data *game.ControlUserEnterroom) {
		logger.Printf("进桌返回,uid:%d, data:%#v", uid, data)
		dur := time.Now().Sub(t).Nanoseconds()
		if dur > 100 {
			logger.Printf("%d 登录 %s 消耗: %dms \n", uid, serverdata.Gsidtid, dur)
		}
		rel := data.GetRel()
		if rel == 0 {
			g.enterMaxConnects = 0
			channel.SetStatus(game.ChannelWorking)
			player := game.UMHandler.GetUser(uid)
			if player != nil {
				player.SetPlayerChannel(channel)
			}
			if cb != nil {
				cb(data)
			}
		} else {
			if rel == game.ControlUserEnterroom_ENTER_WRONG_RELOAD {
				err := channel.LogoutGame(true)
				fmt.Println(err)
			}
			err := channel.Destory(true)
			if err != nil {
				fmt.Println(err)
				return
			}
			if rel == game.ControlUserEnterroom_ENTER_WRONG_INGAME {
				if g.enterMaxConnects <= 10 {
					g.enterMaxConnects++
					time.AfterFunc(100*time.Nanosecond, func() {
						err = g.EnterToGame(uid, serverdata, cb, isretry)
						if err != nil {
							fmt.Println(err)
						}
					})
					return
				}
			}
			if cb != nil {
				cb(data)
			}
		}
	}
	err = channel.LoginGame(serverdata, game.TickHandler.GetTick(reflect.ValueOf(callback)), isretry)
	return err
}

func (g *GameManager) GetGameListState() map[string]interface{} {
	list := make(map[string]interface{})
	for gsid, server := range g.Serversort {
		if server != nil {
			use, nouse := game.TableManager.GetUseTableCount(gsid)
			list[gsid] = map[string]interface{}{
				"users":  server.GetUserCount(),
				"tables": server.GetTableCount(),
				"state":  map[string]interface{}{"use": use, "nouse": nouse},
			}
		}
	}
	return list
}

func (g *GameManager) GetTable(gsidtid string) *game.GameTable {
	if table, ok := g.Alltablesort[gsidtid]; ok {
		return table
	}
	return nil
}

func init() {
	GMHander = NewGameManager()
	Register(GMHander)
}
