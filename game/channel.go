package game

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jmesyan/nano/connectors"
	"github.com/nats-io/nats.go"
	"strings"
	"sync"
)

var (
	channels = make(map[int]*GameChannel)
	chanelmu sync.Locker
)

type ChannelStatus int

const (
	_ ChannelStatus = iota
	ChannelCreating
	ChannelCreated
)

func UpdateChannel(gc *GameChannel) {
	chanelmu.Lock()
	channels[gc.Uid] = gc
	chanelmu.Unlock()
}

func GetChannel(uid int) *GameChannel {
	if channel, ok := channels[uid]; ok {
		return channel
	}
	return nil
}

type GameChannel struct {
	Id         string
	Uid        int
	ClientNid  string
	ClientAddr string
	GameNid    string
	GameAddr   string
	Gsid       string
	Status     ChannelStatus
	conn       *nats.Conn
	service    GameService
	c2sTopic   string
	s2cTopic   string
}

type GameChannelOpt func(gc *GameChannel)

func WitchChanConn(conn *nats.Conn) GameChannelOpt {
	return func(gc *GameChannel) {
		gc.conn = conn
	}
}

func generateChannelId(uid int) string {
	str := fmt.Sprintf("00000000000%d", uid)
	return str[len(str)-11:]
}

func NewGameChannel(uid int, clientNid string, conn *nats.Conn, service GameService, opts ...GameChannelOpt) *GameChannel {
	gc := &GameChannel{
		Id:         generateChannelId(uid),
		Uid:        uid,
		ClientNid:  clientNid,
		ClientAddr: strings.TrimLeft(clientNid, "connector_"),
		conn:       conn,
		service:    service,
		Status:     ChannelCreating,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(gc)
		}
	}
	UpdateChannel(gc)
	return gc
}

func (gc *GameChannel) SetGameNid(gameNid, gameAddr string) {
	gc.GameNid = gameNid
	gc.GameAddr = gameAddr
	gc.Gsid = strings.TrimLeft(gameNid, "gameserver_")
	gc.c2sTopic = fmt.Sprintf("%s.c2s", gameNid)
	gc.s2cTopic = fmt.Sprintf("%s.s2c", gc.ClientNid)
	gc.Status = ChannelCreated
	UpdateChannel(gc)
}

//IsPeer 客户端和服务端是否在一起
func (gc *GameChannel) IsPeer() bool {
	return gc.ClientAddr == gc.GameAddr
}

func (gc *GameChannel) C2S(cmd, msg string) {
	if gc.Status == ChannelCreated {
		if len(cmd) == 0 {
			cmd = "00"
		}
		data := fmt.Sprintf("04AAAA%s%s%s", gc.Id, cmd, msg)
		if gc.IsPeer() {
			server := gc.service.GetServerByGSID(gc.Gsid)
			if server != nil {
				server.SendString(data)
			} else {
				fmt.Println("no find the peer server", cmd, msg)
			}
		} else {
			if gc.conn != nil {
				err := gc.conn.Publish(gc.c2sTopic, []byte(data))
				if err != nil {
					logger.Println(err)
				}
			} else {
				fmt.Println("the conn  is close:", cmd, msg)
			}
		}
	} else {
		logger.Println("the c2s channel is not really:", cmd, msg)
	}
}

func (gc *GameChannel) S2C(heart, cmd int32, msg []byte) {
	if gc.Status == ChannelCreated {
		data := map[string]interface{}{
			"cmd":  cmd,
			"n":    heart,
			"body": base64.StdEncoding.EncodeToString(msg),
		}
		if gc.IsPeer() {
			s, err := connectors.ConnectorHandler.Member(gc.Uid)
			if err != nil {
				fmt.Println(err)
			} else {
				s.Push("game", data)
			}
		} else {
			data["uid"] = gc.Uid
			raw, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			if gc.conn != nil {
				err := gc.conn.Publish(gc.s2cTopic, raw)
				if err != nil {
					logger.Println(err)
				}
			} else {
				fmt.Println("the conn  is close:", cmd, msg)
			}
		}

	} else {
		logger.Println("the s2c channel is not really:", heart, cmd, msg)
	}
}
