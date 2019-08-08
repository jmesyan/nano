package apiserver

var (
	ApiManagerHandler *ApiManager
	DefautApiAddrs    = "127.0.0.1:20066"
)

type ApiManager struct {
	listenaddrs string
}
type ApiManagerOpts func(am *ApiManager)

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

}
