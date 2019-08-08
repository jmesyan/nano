package game

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"net"
	"time"
)

const (
	gameserverStatusWorking  = 0
	gameserverStatusClosed   = 1
	gameserverStatusStarting = 2
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
	gid       int32
	rtype     int32
	ridx      int32
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
	data := p.Data
	cmd := p.Cmd
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
				g.initTables(tables)
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
		g.sendString(fmt.Sprintf("02BEAT%d", t))
	} else {
		g.sendString("02BEAT")
	}
}

func (g *GameServer) sendString(str string) bool {
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

}

func (g *GameServer) initMatchServers(tables []*ControlRoomUsersTableInfo) {

}
func (g *GameServer) initGoldServers(tables []*ControlRoomUsersTableInfo) {

}

func (g *GameServer) initTables(tables []*ControlRoomUsersTableInfo) {
	for _, table := range tables {
		g.addTable(table)
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

func (g *GameServer) handleConn() {
	defer g.conn.Close()
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
	g.gid, g.rtype, g.ridx = gid, rtype, ridx
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
