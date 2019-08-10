package apiserver

import (
	"fmt"
	"github.com/jmesyan/nano/application/stores"
	"log"
	"net"
	"os"
)

var (
	sys    = stores.StoresHandler.Sys
	logger = log.New(os.Stderr, "[apiserver]", log.LstdFlags|log.Llongfile)
)

type ApiServer struct {
	conn       net.Conn
	ReadEndStr string
}
type ApiServerOpts func(as *ApiServer)

func NewApiServer(conn net.Conn, opts ...ApiServerOpts) *ApiServer {
	as := &ApiServer{
		conn:       conn,
		ReadEndStr: "\n[OVER]\n",
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(as)
		}
	}
	go as.handleConn()
	return as
}

func (as *ApiServer) handleConn() {
	defer as.conn.Close()
	buf := make([]byte, 2048)
	decoder := NewDecoder()
	for {
		n, err := as.conn.Read(buf)
		if err != nil {
			fmt.Println(n)
			return
		}
		// TODO(warning): decoder use slice for performance, packet data should be copy before next Decode
		packet, err := decoder.Decode(buf[:n])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if packet == nil {
			continue
		}
		// process all packet
		if err := as.processPacket(packet); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (as *ApiServer) processPacket(p *Packet) error {
	cmd, obj := "", p.Data
	if tmp, ok := obj["cmd"]; ok {
		cmd = tmp.(string)
	}
	switch cmd {
	case "SYS_MAINTENANCE":
		mtype, mt := obj["type"].(int), obj["t"].(int64)
		if mtype == 1 {
			sys.SYS_MAINTENANCE = true
			sys.MAINTENANCE_TIME = mt
		} else {
			sys.SYS_MAINTENANCE = false
			sys.MAINTENANCE_TIME = 0
		}
		//pomelo.app.get("cm").broadcastPlayerServerReboot({ time: sys.MAINTENANCE_TIME });
	}
	return nil
}
