package game

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jmesyan/nano/application/stores"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"net"
	"reflect"
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
	conn      net.Conn
	node      *nodes.Node
	natsaddrs string
	status    int32
	client    *nats.Conn
	msgch     chan *nats.Msg
	shut      chan struct{}
	gsid      string
	gid       int
	rtype     int
	ridx      int
	startTime int
	tablesort map[int32]*GameTable
}

type GameServerOpts func(g *GameServer)

func WithGameServerNatsaddrs(address string) GameServerOpts {
	return func(g *GameServer) {
		g.natsaddrs = address
	}
}

func NewGameServer(conn net.Conn, opts ...GameServerOpts) *GameServer {
	g := &GameServer{
		conn:      conn,
		tablesort: make(map[int32]*GameTable),
		status:    gameserverStatusStarting,
		natsaddrs: nats.DefaultURL,
		msgch:     make(chan *nats.Msg, 64),
		shut:      make(chan struct{}, 1),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(g)
		}
	}
	go g.handleConn()
	return g
}
func (g *GameServer) processPacket(p *Packet) error {
	fmt.Printf("processPacket:%#v\n", p)
	tick := p.T
	data := p.Data
	cmd := p.Cmd
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
			fmt.Printf("server start: gid:%d, the rtype is:%d, the ridx is:%d\n", register.GetGid(), register.GetRidx(), register.GetRtype())
			g.Init(register.GetGid(), register.GetRidx(), register.GetRtype())
		}
	case CMD.OGID_CONTROL_TABLES | CMD.ACK:
		//注册桌子
		desk := &ControlRoomUsers{}
		err := proto.Unmarshal(data, desk)
		if err != nil {
			return err
		} else {
			fmt.Printf("the desks info is  gid:%d, the rtype is:%d, the ridx is:%d \n", desk.GetGid(), desk.GetRtype(), desk.GetRidx())
			gid := desk.GetGid()
			tables := desk.Tableinfo
			if gid < 10 {
				time.AfterFunc(2*time.Second, func() {
					g.initMatchServers(tables)
				})
			} else if gid >= 10 && gid < 20 {
				time.AfterFunc(2*time.Second, func() {
					g.initGoldServers(tables)
				})
			} else {
				logger.Println("initTables:", g.gsid, tables)
				g.initTables(tables)
				TableManager.registerTables(g.gsid, tables)
			}
		}
	case CMD.OGID_CONTROL_HEART_BEAT | CMD.ACK:
		//心跳
		heart := &ControlHeartBeat{}
		err := proto.Unmarshal(data, heart)
		if err != nil {
			return err
		} else {
			fmt.Printf("the heart info is :%d\n", heart.GetNowstamp())
			g.sendHeartBeat(heart.GetNowstamp())
		}
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
	if g.conn != nil {
		_, err := g.conn.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}

func (g *GameServer) dispose() {
	logger.Printf("============服务器%s析构开始=====================\n", g.node.Nid)
	dcm.DeRegisterNode(g.node.Nid)
}

func (g *GameServer) initMatchServers(tables []*ControlRoomUsersTableInfo) {
	//暂无比赛
}
func (g *GameServer) initGoldServers(tables []*ControlRoomUsersTableInfo) {
	grid := fmt.Sprintf("%d_%d", g.gid, g.rtype)
	mgsids := []int{}
	for k, v := range gds.Configs {
		censerver := v.Censerver
		if len(censerver) > 0 && censerver == grid && !utils.InArray(k, mgsids) {
			mgsids = append(mgsids, k)
		}
	}

	logger.Println("initGoldServers", g.gsid, mgsids)
	for _, mgid := range mgsids {
		for gsid, server := range serversort {
			gid, rtype, ridx := GetGameParamsByGsid(gsid)
			if mgid == gid && !IsServerMaintence(gsid) && ridx%2 == g.ridx%2 {
				mtids := []int32{}
				for mtid, _ := range server.tablesort {
					mtids = append(mtids, mtid)
				}
				data, err := json.Marshal(mtids)
				if err != nil {
					logger.Println(err)
					continue
				}
				msg := string(data)
				server.N2S(gid, rtype, ridx, "01", msg)
			}
		}
	}
}

func (g *GameServer) initTables(tables []*ControlRoomUsersTableInfo) {
	for _, table := range tables {
		g.addTable(table)
	}
	//金币场
	if g.gid >= 1000 && g.gid < 5000 {
		if gd, ok := gds.Configs[g.gid]; ok {
			mtids := []int32{}
			for mtid, _ := range g.tablesort {
				mtids = append(mtids, mtid)
			}
			data, err := json.Marshal(mtids)
			if err != nil {
				logger.Println(err)
				return
			}
			msg := string(data)
			if len(gd.Censerver) > 0 {
				for gsid, server := range serversort {
					gid, rtype, ridx := GetGameParamsByGsid(gsid)
					grid := fmt.Sprintf("%d_%d", gid, rtype)
					if grid == gd.Censerver && !IsServerMaintence(gsid) && ridx%2 == g.ridx%2 {
						server.N2S(g.gid, g.rtype, g.ridx, "01", msg)
					}
				}
			}
		}
	}
}

func (g *GameServer) addTable(table *ControlRoomUsersTableInfo) *GameTable {
	gametable := g.getTable(table.GetTid())
	if gametable == nil {
		gametable = NewGameTable()
	}
	gametable.Init(g.gsid, table)
	gametable.gameserver = g
	g.tablesort[gametable.tableid] = gametable
	alltablesort[gametable.tableid] = gametable
	return gametable
}

func (g *GameServer) getTable(tableid int32) *GameTable {
	if table, ok := g.tablesort[tableid]; ok {
		return table
	}
	return nil
}

func (g *GameServer) N2S(gid, rtype, ridx int, cmd, msg string) string {
	if len(cmd) == 0 {
		cmd = "00"
	}
	mgid, mrtype, mridx := g.formatGsid(gid), g.formatGsid(rtype), g.formatGsid(ridx)
	data := fmt.Sprintf("04AAAA%s%s%s%s%s", mgid, cmd, mrtype, mridx, msg)
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
	return g.node.Nid
}

func (g *GameServer) Status() int32 {
	return g.status
}

func (g *GameServer) Init(gid, rtype, ridx int32) {
	g.gid, g.rtype, g.ridx = int(gid), int(rtype), int(ridx)
	g.gsid = fmt.Sprintf("%d_%d_%d", gid, rtype, ridx)
	if oldserver, ok := serversort[g.gsid]; ok {
		g.tablesort = oldserver.tablesort
		oldserver.tablesort = nil
		g.startTime = oldserver.startTime
		g.node = oldserver.node
		g.client = oldserver.client
	} else {
		serversort[g.gsid] = g
		g.InitNats()
	}
	g.startTime = utils.Time()
	g.status = gameserverStatusWorking
}

func (g *GameServer) InitNats() {
	var err error
	nid := utils.GenerateNodeId(nodes.NodeGameServer, g.gsid)
	n := nodes.NewNode("GameServer", nid, nodes.NodeGameServer)
	dcm.RegisterNode(nid, n)
	g.node = n
	g.client, err = nats.Connect(g.natsaddrs)
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
}
