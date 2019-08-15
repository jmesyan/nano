package game

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmesyan/nano/application/cache"
	"github.com/jmesyan/nano/application/models"
	"github.com/nats-io/nats.go"
	"math"
	"strings"
	"sync"
)

var (
	channels = make(map[int]*GameChannel)
	chanelmu sync.Mutex
)

type ChannelStatus int

const (
	_ ChannelStatus = iota
	ChannelClosed
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
	SrvData    *ServerData
	FromGame   bool
	conn       *nats.Conn
	service    GameService
	c2sTopic   string
	s2cTopic   string
	s2cDestory string
	c2sDestory string
}

type GameChannelOpt func(gc *GameChannel)

func WithChanConn(conn *nats.Conn) GameChannelOpt {
	return func(gc *GameChannel) {
		gc.conn = conn
	}
}
func WithFromSource(game bool) GameChannelOpt {
	return func(gc *GameChannel) {
		gc.FromGame = game
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
	gc.c2sDestory = fmt.Sprintf("%s.channel.destory", gameNid)
	gc.s2cTopic = fmt.Sprintf("%s.s2c", gc.ClientNid)
	gc.s2cDestory = fmt.Sprintf("%s.channel.destory", gc.ClientNid)
	gc.Status = ChannelCreated
	UpdateChannel(gc)
}

//IsPeer 客户端和服务端是否在一起
func (gc *GameChannel) IsPeer() bool {
	return gc.ClientAddr == gc.GameAddr
}

func (gc *GameChannel) C2S(msg string, args ...string) error {
	cmd := "00"
	if len(args) > 0 {
		cmd = args[0]
	}
	if gc.Status == ChannelCreated {
		data := fmt.Sprintf("04AAAA%s%s%s", gc.Id, cmd, msg)
		if gc.IsPeer() {
			server := gc.service.GetServerByGSID(gc.Gsid)
			if server != nil {
				server.SendString(data)
			} else {
				return errors.New(fmt.Sprintf("no find the peer server,cmd:%s,msg:%s", cmd, msg))
			}
		} else {
			if gc.conn != nil {
				err := gc.conn.Publish(gc.c2sTopic, []byte(data))
				if err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf("the c2s conn  is close,cmd:%s,msg:%s", cmd, msg))
			}
		}
	} else {
		return errors.New(fmt.Sprintf("the c2s channel is not really,cmd:%s,msg:%s", cmd, msg))
	}
	return nil
}

func (gc *GameChannel) S2C(heart, cmd int32, msg []byte) error {
	if gc.Status == ChannelCreated {
		data := map[string]interface{}{
			"cmd":  cmd,
			"n":    heart,
			"body": base64.StdEncoding.EncodeToString(msg),
		}
		if gc.IsPeer() {
			s, err := ConnectorHandler.Member(gc.Uid)
			if err != nil {
				return err
			} else {
				err := s.Push("game", data)
				if err != nil {
					return err
				}
			}
		} else {
			data["uid"] = gc.Uid
			raw, err := json.Marshal(data)
			if err != nil {
				return err
			}
			if gc.conn != nil {
				err := gc.conn.Publish(gc.s2cTopic, raw)
				if err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf("the s2c conn  is close, cmd:%d,msg:%s", cmd, msg))
			}
		}
	} else {
		return errors.New(fmt.Sprintf("the s2c conn  is not really, cmd:%d,msg:%s", cmd, msg))
	}
	return nil
}

func (gc *GameChannel) LoginGame(serverdata *ServerData, tick int32, isretry bool) error {
	if gc.Status < ChannelCreated {
		return errors.New(fmt.Sprintf("channel not really:%d", gc.Uid))
	}
	gc.SrvData = serverdata
	player := cache.CacheManager.GetUser(gc.Uid)
	achieve := models.GetUserAchievement(gc.Uid)
	tid, quick, quicksit := serverdata.Tid, serverdata.Quick, serverdata.QuickSit
	usertype := 0
	if player.Ismotor == 1 {
		usertype = 4
	}
	bfrom, firstin, outgolds, intime := 0, 0, 0, 0
	name, username, useprop := player.Nickname, player.Username, player.UseProp
	gobj := serverdata.Gobj
	if gobj == nil {
		gobj = make(map[string]interface{})
	}
	gobj["use_avatar"] = player.UseAvatar
	gobj["isandroid"] = player.Ismotor
	gobj["jipaiqi"] = 0
	gobj["lang"] = "0.0"
	gobj["lat"] = "0.0"
	gobj["nchanges"] = achieve.Nchanges
	gobj["totalrounds"] = achieve.Rounds
	gobj["totalloses"] = achieve.Totalloses
	gobj["chancerounds"] = achieve.Chancerounds

	//开始登陆
	instate := 4
	if serverdata.State > 0 {
		instate = serverdata.State
	}
	if serverdata.Scorescale > 0 && instate != 1 {
		outgolds = int(math.Ceil(float64(outgolds) * serverdata.Scorescale))
	}
	if serverdata.Code > 0 {
		gobj["code"] = serverdata.Code
	}
	if serverdata.Lid > 0 {
		gobj["lid"] = serverdata.Lid
	}
	gobj["mlid"] = 0

	gobjbyte, err := json.Marshal(gobj)
	if err != nil {
		return err
	}
	loginip := player.LoginIp
	if len(loginip) == 0 {
		loginip = player.RegIp
	}
	if len(loginip) == 0 {
		loginip = "0.0.0.0"
	}
	gobjstr := string(gobjbyte)
	msg := fmt.Sprintf("%d|%d|%s|%s|%s|%d|%d|%d|%s|%d|%d|%d|%d|%d|%d|%s", tick, gc.Uid, username, name, player.Gender, usertype,
		tid, bfrom, loginip, quick, quicksit, firstin, outgolds, intime, useprop, gobjstr)
	err = gc.C2S(msg, "01")
	if err != nil {
		return err
	}
	if isretry {
		logger.Println(msg)
	}
	return nil
}

func (gc *GameChannel) LogoutGame(destory bool) error {
	err := gc.C2S("", "02")
	if err != nil {
		return err
	}
	if destory {
		err = gc.Destory(gc.IsPeer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (gc *GameChannel) Destory(isPeer bool) error {
	gc.Status = ChannelClosed
	delete(channels, gc.Uid)
	if !isPeer {
		data := map[string]interface{}{
			"uid": gc.Uid,
		}
		raw, err := json.Marshal(data)
		if err != nil {
			return err
		}
		if gc.conn != nil {
			return errors.New(fmt.Sprintf("the destory conn  is close,uid:%d", gc.Uid))
		}
		if gc.FromGame {
			//给前端发送消息
			err := gc.conn.Publish(gc.s2cDestory, []byte(raw))
			if err != nil {
				return err
			}
		} else {
			//给后端发送消息
			err := gc.conn.Publish(gc.c2sDestory, []byte(raw))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
