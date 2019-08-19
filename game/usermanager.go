package game

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"reflect"
	"sync"
	"time"
)

var (
	UMHandler = NewUserManager()
)

type MsgReceiver struct {
	Uid int    `json:"uid"`
	Sid int64  `json:"sid"`
	Nid string `json:"nid"`
}
type MsgLoad struct {
	Receiver *MsgReceiver `json:"receiver"`
	Route    string       `json:"route"`
	Msg      interface{}  `json:"msg"`
}

type UserManager struct {
	locals  map[int]*GamePlayer
	remotes map[int]*GamePlayer
	lmu     sync.RWMutex
	rmu     sync.RWMutex
}

type UserManagerOpt func(um *UserManager)

func NewUserManager(opts ...UserManagerOpt) *UserManager {
	um := &UserManager{
		locals:  make(map[int]*GamePlayer),
		remotes: make(map[int]*GamePlayer),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(um)
		}
	}
	return um
}

func (um *UserManager) NID() string {
	return ConnectorHandler.NID()
}

func (um *UserManager) GetUser(uid int) *GamePlayer {
	if user, ok := um.locals[uid]; ok {
		return user
	}
	if user, ok := um.remotes[uid]; ok {
		return user
	}
	key := fmt.Sprintf("user_%d", uid)
	kv, err := dcm.DCManager.GetValue(key)
	if err != nil {
		logger.Println(err)
		return nil
	}
	var user *GamePlayer
	err = utils.Serializer.Unmarshal(kv.Value, user)
	if err != nil {
		logger.Println(err)
		return nil
	}
	player := NewGamePlayer(user.Uid, user.ConnectorNid, ConnectorHandler.client)
	err = um.AddRemote(player)
	if err != nil {
		logger.Println(err)
		return nil
	}
	if player != nil {
		return player
	}
	return player
}

func (um *UserManager) GetLocalUser(uid int) *GamePlayer {
	if user, ok := um.locals[uid]; ok {
		return user
	}
	return nil
}

func (um *UserManager) RemoveUser(uid int) error {
	//删除本地数据
	delete(um.locals, uid)
	delete(um.remotes, uid)
	//删除远程数据
	key := fmt.Sprintf("user_%d", uid)
	err := dcm.DCManager.DelValue(key)
	if err != nil {
		return err
	}
	return nil
}

func (um *UserManager) StoreUser(u *GamePlayer) error {
	user, err := utils.Serializer.Marshal(u)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("user_%d", u.Uid)
	err = dcm.DCManager.SetValue(key, user)
	if err != nil {
		return err

	}
	return nil
}

func (um *UserManager) KickUser(connectorNid string, receiver *MsgReceiver, state int) error {
	topic := utils.GenerateTopic(connectorNid, "kick")
	payload := &MsgLoad{Receiver: receiver, Msg: map[string]interface{}{"state": state}}
	load, err := utils.Serializer.Marshal(payload)
	if err != nil {
		return err
	}
	client := ConnectorHandler.client
	msg, err := client.Request(topic, load, 10*time.Millisecond)
	if err != nil {
		return err
	}
	resp := string(msg.Data)
	if resp == "SUCCESS" {
		return nil
	}
	return errors.New(resp)
}
func (um *UserManager) PushMsg(connectorNid string, receiver *MsgReceiver, route string, data interface{}) error {
	topic := utils.GenerateTopic(connectorNid, "push")
	payload := &MsgLoad{Receiver: receiver, Route: route, Msg: data}
	load, err := utils.Serializer.Marshal(payload)
	if err != nil {
		return err
	}
	client := ConnectorHandler.client
	msg, err := client.Request(topic, load, 10*time.Millisecond)
	if err != nil {
		return err
	}
	resp := string(msg.Data)
	if resp == "SUCCESS" {
		return nil
	}
	return errors.New(resp)
}

// Multicast  push  the message to the filtered clients
func (um *UserManager) Multicast(route string, v map[string]interface{}, filter SessionFilter) error {
	data, err := utils.SerializeOrRaw(v)
	if err != nil {
		return err
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Type=Multicast Route=%s, Data=%+v", route, v))
	}
	um.lmu.RLock()
	defer um.lmu.RUnlock()
	for _, u := range um.locals {
		if !filter(u) {
			continue
		}
		if err = u.Push(route, data); err != nil {
			logger.Println(err.Error())
		}
	}
	return nil
}

// Broadcast push  the message(s) to  all members
func (um *UserManager) Broadcast(route string, v map[string]interface{}) error {
	data, err := utils.SerializeOrRaw(v)
	if err != nil {
		return err
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Type=Broadcast Route=%s, Data=%+v", route, v))
	}

	um.lmu.RLock()
	defer um.lmu.RUnlock()

	for _, u := range um.locals {
		if err = u.Push(route, data); err != nil {
			logger.Println(fmt.Sprintf("Session push message error, ID=%d, UID=%d, Error=%s", u.Session.ID(), u.Session.UID(), err.Error()))
		}
	}

	return err
}

// Members returns all member's UID in current usermanager
func (um *UserManager) Members() []int {
	um.lmu.RLock()
	defer um.lmu.RUnlock()

	members := []int{}
	for _, u := range um.locals {
		members = append(members, u.Uid)
	}
	return members
}

// Contains check whether a UID is contained in current usermanager or not
func (um *UserManager) Contains(uid int) bool {
	_, err := um.Member(uid)
	return err == nil
}

// Member returns specified UID's session
func (um *UserManager) Member(uid int) (*GamePlayer, error) {
	um.lmu.RLock()
	defer um.lmu.RUnlock()
	for _, u := range um.locals {
		if u.Uid == uid {
			return u, nil
		}
	}

	return nil, ErrMemberNotFound
}

// Add add user to usermanager
func (um *UserManager) Add(u *GamePlayer) error {
	if envdebug {
		logger.Println(fmt.Sprintf("Add session to usermanager, ID=%d, UID=%d", u.Session.ID(), u.Session.UID()))
	}

	um.lmu.Lock()
	defer um.lmu.Unlock()
	delete(um.remotes, u.Uid)

	id := u.Uid
	_, ok := um.locals[id]
	if ok {
		err := um.StoreUser(u)
		return err
	}
	um.locals[id] = u
	err := um.StoreUser(u)
	return err
}
func (um *UserManager) AddRemote(u *GamePlayer) error {
	if envdebug {
		logger.Println(fmt.Sprintf("Add remote  to usermanager, UID=%d", u.Uid))
	}

	um.rmu.Lock()
	defer um.rmu.Unlock()
	delete(um.locals, u.Uid)
	id := u.Uid
	_, ok := um.remotes[id]
	if ok {
		return ErrSessionDuplication
	}
	um.remotes[id] = u
	return nil
}

// Leave remove specified UID related session from usermanager
func (um *UserManager) Leave(u *GamePlayer) error {
	if envdebug {
		logger.Println(fmt.Sprintf("Remove session from usermanager, UID=%d", u.Uid))
	}

	um.lmu.Lock()
	defer um.lmu.Unlock()

	delete(um.locals, u.Uid)
	delete(um.remotes, u.Uid)
	return nil
}

// LeaveAll clear all locals in the usermanager
func (um *UserManager) LeaveAll() error {
	um.lmu.Lock()
	defer um.lmu.Unlock()
	um.locals = make(map[int]*GamePlayer)
	um.remotes = make(map[int]*GamePlayer)
	return nil
}

// Count get current member amount in the usermanager
func (um *UserManager) Count() int {
	um.lmu.RLock()
	defer um.lmu.RUnlock()
	return len(um.locals)
}

func (um *UserManager) DoConnectorMsg(c *Connector, msg *nats.Msg) {
	logger.Printf("handle connector nats msg:%#v\n", msg)
	switch msg.Subject {
	case c.kickTopic:
		payload := &MsgLoad{}
		err := utils.Serializer.Unmarshal(msg.Data, payload)
		if err != nil {
			logger.Println(err)
			msg.Respond(ResponseFail)
			return
		}
		uid := payload.Receiver.Uid
		sid := payload.Receiver.Sid
		nid := payload.Receiver.Nid
		data := payload.Msg.(map[string]interface{})
		//收到踢人消息
		state := 0
		if tmp, ok := data["state"]; ok {
			state = int(tmp.(float64))
		}
		if nid == c.node.Nid {
			u, err := um.Member(uid)
			if err != nil {
				logger.Println(err)
				user := um.GetLocalUser(uid)
				if user != nil && user.Channel != nil {
					um.RemoveUser(uid)
					msg.Respond(ResponseSuccess)
					return
				} else {
					msg.Respond(ResponseFail)
					return
				}
			}
			if sid == u.Session.ID() {
				u.Push("quit", map[string]interface{}{"state": state, "id": sid})
				u.Session.Clear()
				if state != 1 {
					u.Session.Close()
				}
				um.RemoveUser(uid)
			}
		}
		msg.Respond(ResponseSuccess)
	case c.pushTopic:
		payload := &MsgLoad{}
		err := utils.Serializer.Unmarshal(msg.Data, payload)
		if err != nil {
			logger.Println(err)
			msg.Respond(ResponseFail)
			return
		}
		uid := payload.Receiver.Uid
		sid := payload.Receiver.Sid
		nid := payload.Receiver.Nid
		if nid == c.node.Nid {
			u, err := um.Member(uid)
			if err != nil {
				logger.Println(err)
				msg.Respond(ResponseFail)
				return
			}
			if sid == u.Session.ID() {
				err = u.Session.Push(payload.Route, payload.Msg)
				if err != nil {
					logger.Println(err)
				}
			}
		}
		msg.Respond(ResponseSuccess)
	case c.s2cTopic:
		payload := make(map[string]interface{})
		err := utils.Serializer.Unmarshal(msg.Data, payload)
		if err != nil {
			logger.Println(err)
			return
		}
		uid := int(payload["uid"].(float64))
		u, err := um.Member(uid)
		if err != nil {
			logger.Println(err)
			return
		}
		delete(payload, "uid")
		u.Push("game", payload)
	case c.s2cDestory:
		payload := make(map[string]interface{})
		err := utils.Serializer.Unmarshal(msg.Data, payload)
		if err != nil {
			logger.Println(err)
			return
		}
		uid := int(payload["uid"].(float64))
		cn := GetChannel(uid)
		if cn != nil {
			err = cn.Destory(true)
			if err != nil {
				logger.Println(err)
			}
		}
	case c.enterTopic:
		body := &ControlUserEnterroom{}
		err := proto.Unmarshal(msg.Data, body)
		if err != nil {
			logger.Println(err)
			return
		} else {
			uid := int(body.GetUid())
			user := UMHandler.GetUser(uid)
			if user == nil {
				logger.Println("error enter user:", uid)
				return
			}
			if user.IsPeer() {
				TickHandler.ExecTick(body.GetTick(), reflect.ValueOf(body))
			}
		}
	}

}
