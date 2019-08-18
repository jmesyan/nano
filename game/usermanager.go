package game

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/users"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"strings"
	"sync"
	"time"
)

var (
	UserManagerHandler = NewUserManager()
)

type MsgReceiver struct {
	Uid int    `json:"uid"`
	Sid int64  `json:"sid"`
	Nid string `json:"nid"`
}
type MsgLoad struct {
	Receiver *MsgReceiver           `json:"receiver"`
	Route    string                 `json:"route"`
	Msg      map[string]interface{} `json:"msg"`
}

type UserManager struct {
	locals  map[int]*GamePlayer
	remotes map[int]*GamePlayer
	lmu     sync.RWMutex
	rmu     sync.RWMutex
}

type UserManagerOpt func(um *UserManager)

func NewUserManager(opts ...UserManagerOpt) *UserManager {
	um := &UserManager{}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(um)
		}
	}
	return um
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
	user := &users.User{}
	err = utils.Serializer.Unmarshal(kv.Value, user)
	if err != nil {
		logger.Println(err)
		return nil
	}
	clientAddr := strings.TrimLeft(user.ConnectorNid, "connector_")
	player := NewGamePlayer(user.Uid, clientAddr, ConnectorHandler.client)
	um.AddRemote(player)
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

func (um *UserManager) StoreUser(uid int, data *users.User) error {
	user, err := utils.Serializer.Marshal(data)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("user_%d", uid)
	err = dcm.DCManager.SetValue(key, user)
	if err != nil {
		return err

	}
	return nil
}

func (um *UserManager) KickUser(usermanager string, receiver *MsgReceiver, state int) error {
	topic := utils.GenerateTopic(usermanager, "kick")
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
func (um *UserManager) PushMsg(usermanager string, receiver *MsgReceiver, route string, data map[string]interface{}) error {
	topic := utils.GenerateTopic(usermanager, "push")
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
func (um *UserManager) Multicast(route string, v interface{}, filter SessionFilter) error {
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
		if err = u.Sess.Push(route, data); err != nil {
			logger.Println(err.Error())
		}
	}
	return nil
}

// Broadcast push  the message(s) to  all members
func (um *UserManager) Broadcast(route string, v interface{}) error {
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
		if err = u.Sess.Push(route, data); err != nil {
			logger.Println(fmt.Sprintf("Session push message error, ID=%d, UID=%d, Error=%s", u.Sess.ID(), u.Sess.UID(), err.Error()))
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
		logger.Println(fmt.Sprintf("Add session to usermanager, ID=%d, UID=%d", u.Sess.ID(), u.Sess.UID()))
	}

	um.lmu.Lock()
	defer um.lmu.Unlock()
	delete(um.remotes, u.Uid)

	id := u.Uid
	_, ok := um.locals[id]
	if ok {
		return ErrSessionDuplication
	}
	um.locals[id] = u
	return nil
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
		data := payload.Msg
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
				if user != nil && user.Chan != nil && user.Chan.Status < ChannelWorking {
					um.RemoveUser(uid)
					msg.Respond(ResponseSuccess)
					return
				} else {
					msg.Respond(ResponseFail)
					return
				}
			}
			if sid == u.Sess.ID() {
				u.Sess.Push("quit", map[string]interface{}{"state": state, "id": sid})
				u.Sess.Clear()
				if state != 1 {
					u.Sess.Close()
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
			if sid == u.Sess.ID() {
				err = u.Sess.Push(payload.Route, payload.Msg)
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
		u.Sess.Push("game", payload)
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
				fmt.Println(err)
			}
		}
	}

}
