package nano

import (
	"fmt"
	"github.com/jmesyan/nano/game"
	"github.com/jmesyan/nano/utils"
	"net"
	"reflect"
	"sort"
)

var (
	GameManagerHander *GameManager
	DefautListenGame  = "0.0.0.0:20572"
)

type GameManager struct {
	listenaddrs string
	Serversort  map[string]*game.GameServer
}

type GameManagerOpts func(g *GameManager)

func WithGameManagerAddrs(addrs string) GameManagerOpts {
	return func(g *GameManager) {
		g.listenaddrs = addrs
	}
}

func NewGameManager(opts ...GameManagerOpts) *GameManager {
	g := &GameManager{
		listenaddrs: DefautListenGame,
		Serversort:  make(map[string]*game.GameServer),
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

func init() {
	GameManagerHander = NewGameManager()
	Register(GameManagerHander)
}
