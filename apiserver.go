package nano

import (
	"encoding/json"
	"fmt"
	"github.com/jmesyan/nano/apiserver"
	"github.com/jmesyan/nano/game"
	"github.com/jmesyan/nano/session"
	"net"
)

type ApiServer struct {
	conn       net.Conn
	ReadEndStr []byte
}
type ApiServerOpts func(as *ApiServer)

func NewApiServer(conn net.Conn, opts ...ApiServerOpts) *ApiServer {
	as := &ApiServer{
		conn:       conn,
		ReadEndStr: []byte("\n[OVER]\n"),
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
	decoder := apiserver.NewDecoder()
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

func (as *ApiServer) processPacket(p *apiserver.Packet) error {
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
		filter := func(s *session.Session) bool {
			tableid := s.Int("tableid")
			return tableid == 0
		}
		err := game.ConnectorHandler.Multicast("serverReboot", map[string]interface{}{"time": sys.MAINTENANCE_TIME}, filter)
		if err != nil {
			logger.Println(err)
		}
	case "getGameListState":
		list := GameManagerHander.GetGameListState()
		as.SendObj(list)
	}
	return nil
}

func (as *ApiServer) SendObj(data map[string]interface{}) {
	if as.conn != nil {
		msg, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		} else {
			_, err = as.conn.Write(msg)
			fmt.Println(err)
		}
		_, err = as.conn.Write(as.ReadEndStr)
		fmt.Println(err)
	}
}
