package nano

import (
	"fmt"
	"github.com/jmesyan/nano/game"
	"net"
)

var (
	GameManagerHander *GameManager
	DefautListenGame  = "0.0.0.0:20572"
)

type GameManager struct {
	listenaddrs string
	gameservers map[string]*game.GameServer
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
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(g)
		}
	}
	GameManagerHander = g
	return g
}
func (g *GameManager) Init() {
	listen, err := net.Listen("tcp", "0.0.0.0:20572")
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
		game.NewGameServer(conn)
	}
}
func (g *GameManager) AfterInit() {

}
func (g *GameManager) BeforeShutdown() {

}
func (g *GameManager) Shutdown() {

}
