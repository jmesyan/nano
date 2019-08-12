package nano

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/game"
	"github.com/jmesyan/nano/utils"
	"net"
	"reflect"
	"sort"
	"time"
)

var (
	GameManagerHander *GameManager
	DefautListenGame  = "0.0.0.0:20572"
)

type GameManager struct {
	listenaddrs      string
	Serversort       map[string]*game.GameServer
	enterMaxConnects int
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
		Serversort:       make(map[string]*game.GameServer),
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
	go g.watcher()
}
func (g *GameManager) watcher() {
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
func (g *GameManager) AfterInit() {

}
func (g *GameManager) BeforeShutdown() {

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
		mt.Method.Func.Call([]reflect.Value{body})
	} else {
		fmt.Printf("the is no srv handler, route is:%s, body is:%#v", route, body)
	}
}

func (g *GameManager) RegisterServer(gsid string, server *game.GameServer) {
	g.Serversort[gsid] = server
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
}

func (g *GameManager) GetCenterServerByBalance(ngid int) *game.GameServer {
	config, ok := gds.Configs[ngid]
	if !ok {
		return nil
	}
	gcid := config.Censerver
	gsids := make(map[int]string)
	for gsid, _ := range g.Serversort {
		gid, rtype, _ := game.GetGameParamsByGsid(gsid)
		grid := game.GetGrid(gid, rtype)
		if grid == gcid && !game.IsServerMaintence(gsid) {
			gsids[gds.Gcsu[gsid]] = gsid
		}
	}
	if len(gsids) > 0 {
		gsorts := make([]int, len(gsids))
		for k, _ := range gsids {
			gsorts = append(gsorts, k)
		}
		sort.Ints(gsorts)
		gsid := gsids[gsorts[0]]
		return g.Serversort[gsid]
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
		sess, err := game.ConnectorHandler.Member(uid)
		if err != nil {
			return nil, err
		}
		err = sess.Push("reconnect", serverdata)
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

func init() {
	GameManagerHander = NewGameManager()
	Register(GameManagerHander)
}
