package game

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jmesyan/nano/application/models"
	"github.com/jmesyan/nano/application/stores"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"net"
	"reflect"
	"sort"
	"time"
)

const (
	gameserverStatusWorking  = 0
	gameserverStatusClosed   = 1
	gameserverStatusStarting = 2
)

var (
	gds = stores.StoresHandler.Gds
	sys = stores.StoresHandler.Sys
)

type GameServer struct {
	Node       *nodes.Node          `json:"node"`
	Natsaddrs  string               `json:"natsaddrs"`
	Status     int32                `json:"status"`
	Tablesort  map[int32]*GameTable `json:"tablesort"`
	Gsid       string               `json:"gsid"`
	Gid        int                  `json:"gid"`
	Rtype      int                  `json:"rtype"`
	Ridx       int                  `json:"ridx"`
	StartTime  int                  `json:"startTime"`
	C2sTopic   string               `json:"c2sTopic"`
	SsTopic    string               `json:"SsTopic"`
	C2sDestory string               `json:"c2sDestory"`
	IsRemote   bool                 `json:"isRemote"`
	client     *nats.Conn
	conn       net.Conn
	msgch      chan *nats.Msg
	shut       chan struct{}
	service    GameService
}

type GameServerOpt func(g *GameServer)

func WithGameServerNatsaddrs(address string) GameServerOpt {
	return func(g *GameServer) {
		g.Natsaddrs = address
	}
}

func NewGameServer(conn net.Conn, service GameService, opts ...GameServerOpt) *GameServer {
	g := &GameServer{
		conn:      conn,
		Tablesort: make(map[int32]*GameTable),
		Status:    gameserverStatusStarting,
		Natsaddrs: nats.DefaultURL,
		msgch:     make(chan *nats.Msg, 64),
		shut:      make(chan struct{}, 1),
		service:   service,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(g)
		}
	}
	go g.handleConn()
	return g
}

func (g *GameServer) GetNode() *nodes.Node {
	return g.Node
}

func (g *GameServer) SetClient(client *nats.Conn) {
	g.client = client
}

func (g *GameServer) processPacket(p *Packet) error {
	fmt.Printf("processPacket:%#v\n", p)
	cid := int(p.Cid)
	heart := p.N
	tick := p.T
	data := p.Data
	cmd := p.Cmd
	if cid > 0 {
		cn := GetChannel(cid)
		if cn != nil {
			err := cn.S2C(heart, cmd, data) //消息转发到客户端
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("can't find the channel:%#v", p)
		}
		return nil
	}
	if tick > 0 {
		switch cmd {
		case CMD.OGID_CONTROL_USER_SIGN | CMD.REQ:
			body := &ControlUserSign{}
			err := proto.Unmarshal(data, body)
			if err != nil {
				return err
			} else {
				TickHandler.ExecTick(tick, reflect.ValueOf(body))
			}
		case CMD.OGID_ROOMSVR_ENTERROOM | CMD.ACK:
			body := &ControlUserEnterroom{}
			err := proto.Unmarshal(data, body)
			if err != nil {
				return err
			} else {
				uid := int(body.GetUid())
				user := UMHandler.GetUser(uid)
				if user == nil {
					logger.Println("error enter user:", uid)
					return nil
				}
				if user.IsPeer() {
					TickHandler.ExecTick(tick, reflect.ValueOf(body))
				} else {
					channel := NewGameChannel(uid, user.ConnectorNid, user.client, g.service, WithFromGame(true))
					channel.SetGameNid(g.NID(), g.GetNode().Address)
					channel.SetStatus(ChannelWorking)
					user.SetPlayerChannel(channel)
					err = UMHandler.AddRemote(user)
					if err != nil {
						logger.Println(err)
						return err
					}
					body.Tick = &tick
					ndata, err := proto.Marshal(body)
					if err != nil {
						logger.Println(err)
						return err
					}
					user.NotifyConnector("game.enter", ndata)
				}

			}
		}
		return nil
	}
	switch cmd {
	case CMD.OGID_CONTROL_REGIS | CMD.ACK:
		//握手
		register := &RegisterServer{}
		err := proto.Unmarshal(data, register)
		if err != nil {
			return err
		} else {
			fmt.Printf("server start: Gid:%d, the Rtype is:%d, the Ridx is:%d\n", register.GetGid(), register.GetRtype(), register.GetRidx())
			g.Init(register.GetGid(), register.GetRtype(), register.GetRidx())
		}
	case CMD.OGID_CONTROL_TABLES | CMD.ACK:
		//注册桌子
		desk := &ControlRoomUsers{}
		err := proto.Unmarshal(data, desk)
		if err != nil {
			return err
		} else {
			fmt.Printf("the desks info is  Gid:%d, the Rtype is:%d, the Ridx is:%d \n", desk.GetGid(), desk.GetRtype(), desk.GetRidx())
			Gid := desk.GetGid()
			tables := desk.Tableinfo
			if Gid < 10 {
				time.AfterFunc(2*time.Second, func() {
					g.initMatchServers(tables)
				})
			} else if Gid >= 10 && Gid < 20 {
				time.AfterFunc(2*time.Second, func() {
					g.initGoldServers(tables)
				})
			} else {
				logger.Println("initTables:", g.Gsid, tables)
				g.initTables(tables)
				TableManager.registerTables(g.Gsid, tables)
			}
		}
	case CMD.OGID_CONTROL_HEART_BEAT | CMD.ACK:
		//心跳
		heart := &ControlHeartBeat{}
		err := proto.Unmarshal(data, heart)
		if err != nil {
			return err
		} else {
			fmt.Printf("the heart info is: gsid:%s, timestamp :%d\n", g.Gsid, heart.GetNowstamp())
			g.sendHeartBeat(heart.GetNowstamp())
		}
	case CMD.OGID_CONTROL_DISTRIBUTE_USER | CMD.ACK:
		body := &ControlDistributeUsers{}
		err := proto.Unmarshal(data, body)
		if err != nil {
			return err
		} else {
			g.service.ProcessServer("hall.user.goldEnterRoom", reflect.ValueOf(body))
		}
	case CMD.OGID_GAME_MSG | CMD.ACK:
		body := &ControlGameMsg{}
		err := proto.Unmarshal(data, body)
		if err != nil {
			return err
		}
		logger.Printf("control_game_msg, gsid:%s, body:%#v", g.Gsid, body)
		uid, mtype, mtid, mpos := body.GetUid(), body.GetType(), body.GetTid(), body.GetPos()
		if mtype == 0 { //玩家进入
			table := g.GetTable(mtid)
			if table == nil {
				table = g.addTable(&ControlRoomUsersTableInfo{
					Tid: body.Tid,
				})
			}
			table.AddPlayer(uid)
			models.AddUserOnline(map[string]interface{}{"userid": uid, "gid": g.Gid, "rtype": g.Rtype, "ridx": g.Ridx, "tid": mtid, "pos": mpos})
		} else if mtype == 2 { //离开房间
			table := g.GetTable(mtid)
			if table != nil {
				table.RemovePlayer(uid)
				models.RemoveUserOnline(int(uid))
			}
		}
	case CMD.OGID_CONTROL_CANCEL_TABLE | CMD.ACK:
		body := &ControlCancelTable{}
		err := proto.Unmarshal(data, body)
		if err != nil {
			return err
		}
		gid := int32(g.Gid)
		body.Gid = &gid
		logger.Printf("control_cancel_table, gsid:%s, body:%#v", g.Gsid, body)
		g.service.ProcessServer("hall.user.endGame", reflect.ValueOf(body))
	}
	return nil
}

func (g *GameServer) sendHeartBeat(t uint32) {
	if t > 0 {
		g.SendString("02BEAT%d", t)
	} else {
		g.SendString("02BEAT")
	}
}

func (g *GameServer) SendString(format string, args ...interface{}) bool {
	str := fmt.Sprintf(format, args...)
	if str == "B" {
		g.dispose()
		return false
	}
	if g.IsRemote {
		err := g.client.Publish(g.SsTopic, []byte(str))
		if err != nil {
			logger.Println("sendstring remote error:", err.Error())
		}
	} else {
		if g.conn != nil {
			str += "\x00"
			logger.Println("SendString:", str)
			_, err := g.conn.Write([]byte(str))
			if err != nil {
				fmt.Println(err)
			}
			return true
		}
	}
	return false
}

func (g *GameServer) dispose() {
	logger.Printf("============服务器%s析构开始=====================\n", g.Node.Nid)
	dcm.DeRegisterNode(g.Node.Nid)
	g.service.RemoveServerByGSID(g.Gsid)
}

func (g *GameServer) initMatchServers(tables []*ControlRoomUsersTableInfo) {
	//暂无比赛
}
func (g *GameServer) initGoldServers(tables []*ControlRoomUsersTableInfo) {
	grid := fmt.Sprintf("%d_%d", g.Gid, g.Rtype)
	mGsids := []int{}
	for k, v := range gds.Configs {
		censerver := v.Censerver
		if len(censerver) > 0 && censerver == grid && !utils.InArray(k, mGsids) {
			mGsids = append(mGsids, k)
		}
	}

	logger.Println("initGoldServers", g.Gsid, mGsids)
	serversort := g.service.GetServerSort()
	for _, mGid := range mGsids {
		for Gsid, server := range serversort {
			Gid, Rtype, Ridx := GetGameParamsByGsid(Gsid)
			if mGid == Gid && !IsServerMaintence(Gsid) && Ridx%2 == g.Ridx%2 {
				mtids := []int32{}
				for mtid, _ := range server.Tablesort {
					mtids = append(mtids, mtid)
				}
				sort.Slice(mtids, func(i, j int) bool {
					return mtids[i] < mtids[j]
				})
				data, err := json.Marshal(mtids)
				if err != nil {
					logger.Println(err)
					continue
				}
				msg := string(data)
				logger.Println("gold initTables2:", Gid, Rtype, Ridx, "01", msg)
				g.N2S(Gid, Rtype, Ridx, "01", msg)
			}
		}
	}
}

func (g *GameServer) initTables(tables []*ControlRoomUsersTableInfo) {
	for _, table := range tables {
		g.addTable(table)
	}
	//金币场
	if g.Gid >= 1000 && g.Gid < 5000 {
		if gd, ok := gds.Configs[g.Gid]; ok {
			mtids := []int32{}
			for mtid, _ := range g.Tablesort {
				mtids = append(mtids, mtid)
			}
			data, err := json.Marshal(mtids)
			if err != nil {
				logger.Println(err)
				return
			}
			msg := string(data)
			if len(gd.Censerver) > 0 {
				serversort := g.service.GetServerSort()
				for Gsid, server := range serversort {
					Gid, Rtype, Ridx := GetGameParamsByGsid(Gsid)
					grid := fmt.Sprintf("%d_%d", Gid, Rtype)
					if grid == gd.Censerver && !IsServerMaintence(Gsid) && Ridx%2 == g.Ridx%2 {
						server.N2S(g.Gid, g.Rtype, g.Ridx, "01", msg)
					}
				}
			}
		}
	}
}

func (g *GameServer) addTable(table *ControlRoomUsersTableInfo) *GameTable {
	gametable := g.GetTable(table.GetTid())
	if gametable == nil {
		gametable = NewGameTable()
	}
	gametable.Init(g.Gsid, table)
	g.Tablesort[gametable.Tableid] = gametable
	g.service.RegisterTable(gametable.Gsidtid, gametable)
	return gametable
}

func (g *GameServer) GetTable(tableid int32) *GameTable {
	if table, ok := g.Tablesort[tableid]; ok {
		return table
	}
	return nil
}

func (g *GameServer) N2S(Gid, Rtype, Ridx int, cmd, msg string) string {
	if len(cmd) == 0 {
		cmd = "00"
	}
	mGid, mRtype, mRidx := g.formatGsid(Gid), g.formatGsid(Rtype), g.formatGsid(Ridx)
	data := fmt.Sprintf("04AAAA%s%s%s%s%s", mGid, cmd, mRtype, mRidx, msg)
	logger.Println(data)
	g.SendString(data)
	return data
}

func (g *GameServer) formatGsid(id int) string {
	format := "0000000000%d"
	ids := fmt.Sprintf(format, id)
	return ids[len(ids)-10:]
}

func (g *GameServer) handleConn() {
	defer g.dispose()
	buf := make([]byte, 2048)
	decoder := NewDecoder()
	for {
		n, err := g.conn.Read(buf)
		if err != nil {
			fmt.Println(n)
			return
		}
		// TODO(warning): decoder use slice for performance, packet data should be copy before next Decode
		packets, err := decoder.Decode(buf[:n])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(packets) < 1 {
			continue
		}
		// process all packet
		for i := range packets {
			if err := g.processPacket(packets[i]); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (g *GameServer) NID() string {
	return g.Node.Nid
}
func (g *GameServer) Init(Gid, Rtype, Ridx int32) {
	g.Gid, g.Rtype, g.Ridx = int(Gid), int(Rtype), int(Ridx)
	g.Gsid = fmt.Sprintf("%d_%d_%d", Gid, Rtype, Ridx)
	oldserver := g.service.GetServerByGSID(g.Gsid)
	if oldserver != nil {
		g.Tablesort = oldserver.Tablesort
		oldserver.Tablesort = nil
		g.StartTime = oldserver.StartTime
		g.Node = oldserver.Node
		g.client = oldserver.client
	} else {
		g.service.RegisterServer(g.Gsid, g)
		g.InitNats()
	}
	g.StartTime = utils.Time()
	g.Status = gameserverStatusWorking
}

func (g *GameServer) InitNats() {
	var err error
	nid := utils.GenerateNodeId(nodes.NodeGameServer, g.Gsid)
	n := nodes.NewNode("GameServer", nid, nodes.NodeGameServer, nodes.WithNodeAddress(utils.GenerateLocalAddr()), nodes.WithNodeGsid(g.Gsid))
	dcm.RegisterNode(nid, n)
	g.Node = n
	g.client, err = nats.Connect(g.Natsaddrs)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = g.client.ChanQueueSubscribe("queue_GameServer.>", "queue_GameServer", g.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = g.client.ChanSubscribe("GameServer.>", g.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = g.client.ChanSubscribe(fmt.Sprintf("%s.>", n.Nid), g.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	//设置topic
	g.C2sTopic = fmt.Sprintf("%s.c2s", g.NID())
	g.SsTopic = fmt.Sprintf("%s.ss", g.NID())
	g.C2sDestory = fmt.Sprintf("%s.channel.destory", g.NID())
	//监听消息
	go g.watcher()
}
func (g *GameServer) watcher() {
	for {
		select {
		case msg := <-g.msgch:
			g.HandleMsg(msg)
		case <-g.shut:
			logger.Println("receive stop msg")
			close(g.msgch)
			g.Node.Status = nodes.NodeStoping
			return
		}
	}
}

func (g *GameServer) HandleMsg(msg *nats.Msg) {
	logger.Printf("handle gameserver nats msg:%#v\n", msg)
	switch msg.Subject {
	case g.C2sTopic:
		g.SendString(string(msg.Data))
	case g.C2sDestory:
		payload := make(map[string]interface{})
		err := utils.Serializer.Unmarshal(msg.Data, payload)
		if err != nil {
			logger.Println(err)
			return
		}
		uid := int(payload["uid"].(float64))
		cn := GetChannel(uid)
		if cn != nil {
			err = cn.Destory(true)
			if err != nil {
				fmt.Println(err)
			}
		}
	case g.SsTopic:
		msg := string(msg.Data)
		g.SendString(msg)
	}
}

func (g *GameServer) GetUserCount() int {
	users := 0
	for _, table := range g.Tablesort {
		if table != nil {
			users += table.GetPlayerCount()
		}
	}
	return users
}

func (g *GameServer) GetTableCount() int {
	return len(g.Tablesort)
}
