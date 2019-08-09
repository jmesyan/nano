package apiserver

import (
	"fmt"
	"github.com/jmesyan/nano"
	"github.com/jmesyan/nano/application/stores"
	"log"
	"net"
	"os"
)

var (
	ApiManagerHandler *ApiManager
	DefautApiAddrs    = "127.0.0.1:20066"
	logger            = log.New(os.Stderr, "apiserver", log.LstdFlags|log.Llongfile)
)
var (
	sys = stores.StoresHandler.Sys
	gds = stores.StoresHandler.Gds
)

type ApiManager struct {
	listenaddrs string
}
type ApiManagerOpts func(am *ApiManager)

func WithApiAddrs(addrs string) ApiManagerOpts {
	return func(am *ApiManager) {
		am.listenaddrs = addrs
	}
}

func NewApiManager(opts ...ApiManagerOpts) *ApiManager {
	am := &ApiManager{
		listenaddrs: DefautApiAddrs,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(am)
		}
	}
	return am
}

func (am *ApiManager) Init() {
	listen, err := net.Listen("tcp", am.listenaddrs)
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
		NewApiServer(conn)
	}
}

func (am *ApiManager) AfterInit() {

}
func (am *ApiManager) BeforeShutdown() {

}

func (am *ApiManager) Shutdown() {

}

func init() {
	ApiManagerHandler = NewApiManager()
	nano.Register(ApiManagerHandler)
}
