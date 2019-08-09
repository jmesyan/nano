package apiserver

import (
	"fmt"
	"github.com/jmesyan/nano"
	"github.com/jmesyan/nano/application/cache"
	"github.com/jmesyan/nano/application/stores"
	"github.com/jmesyan/nano/utils"
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
	sys.SYS_MAINTENANCE = true
	//维护开关
	maintence := cache.CacheManager.GetMaintence()
	if maintence != nil {
		if maintence.Type == 1 {
			sys.SYS_MAINTENANCE = true
		}
		time := utils.Time()
		if maintence.Type == 1 && time < maintence.Time {
			sys.MAINTENANCE_TIME2 = maintence.Time
		}
	}
	//加载金币房间配置
	loadGoldsTypeConfig()

}

func loadGoldsTypeConfig() {
	gds.Configs = cache.CacheManager.GetGameGoldsType()
}

func (am *ApiManager) AfterInit() {
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
func (am *ApiManager) BeforeShutdown() {

}

func (am *ApiManager) Shutdown() {

}

func init() {
	ApiManagerHandler = NewApiManager()
	nano.Register(ApiManagerHandler)
}
