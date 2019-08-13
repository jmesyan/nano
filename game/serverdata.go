package game

import "fmt"

type ServerData struct {
	Gid        int                    `json:"gid"`
	Rtype      int                    `json:"rtype"`
	Ridx       int                    `json:"ridx"`
	Tid        int                    `json:"tid"`
	Gsid       string                 `json:"gsid"`
	Gsidtid    string                 `json:"gsidtid"`
	State      int                    `json:"state"`
	Initscore  int                    `json:"initscore"`
	Motor      int                    `json:"motor"`
	Lid        int                    `json:"lid"`
	Code       int                    `json:"code"`
	Quick      int                    `json:"quick"`
	QuickSit   int                    `json:"quicksit"`
	Scorescale float64                `json:"scorescale"`
	Gobj       map[string]interface{} `json:"gobj"`
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
