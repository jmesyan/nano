package game

import "fmt"

type ServerData struct {
	Gid        int
	Rtype      int
	Ridx       int
	Tid        int
	Gsid       string
	Gsidtid    string
	State      int
	Initscore  int
	Motor      int
	Lid        int
	Code       int
	Quick      int
	QuickSit   int
	Scorescale float64
	Gobj       map[string]interface{}
}

type ServerDataOpt func(sd *ServerData)

func NewServerData(gid, rtype, ridx, tid int, gobj map[string]interface{}, opts ...ServerDataOpt) *ServerData {
	sd := &ServerData{
		Gid:     gid,
		Rtype:   rtype,
		Ridx:    ridx,
		Tid:     tid,
		Gobj:    gobj,
		Gsid:    fmt.Sprintf("%d_%d_%d", gid, rtype, ridx),
		Gsidtid: fmt.Sprintf("%d_%d_%d_%d", gid, rtype, ridx, tid),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(sd)
		}
	}
	return sd
}
