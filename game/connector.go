// Copyright (c) nano Author. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package game

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/users"
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"sync"
	"sync/atomic"
	"time"
)

var (
	envdebug         = true
	ConnectorHandler *Connector
)

const (
	connectorStatusWorking = 0
	connectorStatusClosed  = 1
)

var (
	ErrMemberNotFound       = errors.New("member not found in the connector")
	ErrClosedConnector      = errors.New("connector closed")
	ErrSessionDuplication   = errors.New("session has existed in the current connector")
	ErrCloseClosedConnector = errors.New("close closed connector")
	ResponseSuccess         = []byte("SUCCESS")
	ResponseFail            = []byte("FAIL")
)

// SessionFilter represents a filter which was used to filter session when Multicast,
// the session will receive the message while filter returns true.
type SessionFilter func(*GamePlayer) bool

// Connector represents a session connector which used to manage a number of
// users, data send to the connector will send to all session in it.
type Connector struct {
	node       *nodes.Node
	mu         sync.RWMutex
	status     int32               // channel current status
	users      map[int]*GamePlayer // session id map to session instance
	natsaddrs  string
	client     *nats.Conn
	msgch      chan *nats.Msg
	shut       chan struct{}
	kickTopic  string
	pushTopic  string
	s2cTopic   string
	s2cDestory string
}

type ConnectorOpts func(g *Connector)

func WithConnectorNatsaddrs(address string) ConnectorOpts {
	return func(c *Connector) {
		c.natsaddrs = address
	}
}

// StartConnector start a new connector instance
func NewConnector(opts ...ConnectorOpts) *Connector {
	c := &Connector{
		status:    connectorStatusWorking,
		users:     make(map[int]*GamePlayer),
		natsaddrs: nats.DefaultURL,
		msgch:     make(chan *nats.Msg, 64),
		shut:      make(chan struct{}, 1),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(c)
		}
	}
	return c
}

func (c *Connector) NID() string {
	return c.node.Nid
}

func (c *Connector) Status() int32 {
	return c.status
}
func (c *Connector) GetClient() *nats.Conn {
	return c.client
}

func (c *Connector) Init() {
	var err error
	nid := utils.GenerateNodeId(nodes.NodeConnector, "")
	n := nodes.NewNode("connector", nid, nodes.NodeConnector, nodes.WithNodeAddress(utils.GenerateLocalAddr()))
	dcm.RegisterNode(nid, n)
	c.node = n
	c.client, err = nats.Connect(c.natsaddrs)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = c.client.ChanQueueSubscribe("queue_connector.>", "queue_connector", c.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = c.client.ChanSubscribe("connector.>", c.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = c.client.ChanSubscribe(fmt.Sprintf("%s.>", n.Nid), c.msgch)
	if err != nil {
		logger.Fatal(err)
		return
	}
	//设置topic
	c.kickTopic = utils.GenerateTopic(c.node.Nid, "kick")
	c.pushTopic = utils.GenerateTopic(c.node.Nid, "push")
	c.s2cTopic = utils.GenerateTopic(c.node.Nid, "s2c")
	c.s2cDestory = utils.GenerateTopic(c.node.Nid, "channel.destory")
}

func (c *Connector) AfterInit() {
	go c.watcher()
}

// Member returns specified UID's session
func (c *Connector) Member(uid int) (*GamePlayer, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, u := range c.users {
		if u.Uid == uid {
			return u, nil
		}
	}

	return nil, ErrMemberNotFound
}

func (c *Connector) watcher() {
	for {
		select {
		case msg := <-c.msgch:
			c.HandleMsg(msg)
		case <-c.shut:
			logger.Println("receive stop msg")
			close(c.msgch)
			c.node.Status = nodes.NodeStoping
			return
		}
	}
}

func (c *Connector) HandleMsg(msg *nats.Msg) {
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
			u, err := c.Member(uid)
			if err != nil {
				logger.Println(err)
				user := c.GetUser(uid)
				if user != nil && len(user.GameserverNid) == 0 {
					c.RemoveUser(uid)
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
				c.RemoveUser(uid)
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
			u, err := c.Member(uid)
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
		u, err := c.Member(uid)
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

func (c *Connector) PushMsg(connector string, receiver *MsgReceiver, route string, data map[string]interface{}) error {
	topic := utils.GenerateTopic(connector, "push")
	payload := &MsgLoad{Receiver: receiver, Route: route, Msg: data}
	load, err := utils.Serializer.Marshal(payload)
	if err != nil {
		return err
	}
	msg, err := c.client.Request(topic, load, 10*time.Millisecond)
	if err != nil {
		return err
	}
	resp := string(msg.Data)
	if resp == "SUCCESS" {
		return nil
	}
	return errors.New(resp)
}

func (c *Connector) KickUser(connector string, receiver *MsgReceiver, state int) error {
	topic := utils.GenerateTopic(connector, "kick")
	payload := &MsgLoad{Receiver: receiver, Msg: map[string]interface{}{"state": state}}
	load, err := utils.Serializer.Marshal(payload)
	if err != nil {
		return err
	}
	msg, err := c.client.Request(topic, load, 10*time.Millisecond)
	if err != nil {
		return err
	}
	resp := string(msg.Data)
	if resp == "SUCCESS" {
		return nil
	}
	return errors.New(resp)
}

func (c *Connector) StoreUser(uid int, data *users.User) error {
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
func (c *Connector) RemoveUser(uid int) error {
	key := fmt.Sprintf("user_%d", uid)
	err := dcm.DCManager.DelValue(key)
	if err != nil {
		return err
	}
	return nil
}

func (c *Connector) GetUser(uid int) *users.User {
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
	return user
}

func (c *Connector) BeforeShutdown() {
	c.shut <- struct{}{}
}

func (c *Connector) Shutdown() {
	err := dcm.DeRegisterNode(c.node.Nid)
	if err != nil {
		fmt.Println(err)
	}
	close(c.shut)
	c.node.Status = nodes.NodeStoped
}

// Members returns all member's UID in current connector
func (c *Connector) Members() []int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	members := []int64{}
	for _, s := range c.users {
		members = append(members, s.Sess.UID())
	}

	return members
}

// Multicast  push  the message to the filtered clients
func (c *Connector) Multicast(route string, v interface{}, filter SessionFilter) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	data, err := utils.SerializeOrRaw(v)
	if err != nil {
		return err
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Type=Multicast Route=%s, Data=%+v", route, v))
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, u := range c.users {
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
func (c *Connector) Broadcast(route string, v interface{}) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	data, err := utils.SerializeOrRaw(v)
	if err != nil {
		return err
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Type=Broadcast Route=%s, Data=%+v", route, v))
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, u := range c.users {
		if err = u.Sess.Push(route, data); err != nil {
			logger.Println(fmt.Sprintf("Session push message error, ID=%d, UID=%d, Error=%s", u.Sess.ID(), u.Sess.UID(), err.Error()))
		}
	}

	return err
}

// Contains check whether a UID is contained in current connector or not
func (c *Connector) Contains(uid int) bool {
	_, err := c.Member(uid)
	return err == nil
}

// Add add user to connector
func (c *Connector) Add(s *GamePlayer) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Add session to connector, ID=%d, UID=%d", s.Sess.ID(), s.Sess.UID()))
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	id := s.Uid
	_, ok := c.users[id]
	if ok {
		return ErrSessionDuplication
	}
	c.users[id] = s
	return nil
}

// Leave remove specified UID related session from connector
func (c *Connector) Leave(s *GamePlayer) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	if envdebug {
		logger.Println(fmt.Sprintf("Remove session from connector, UID=%d", s.Uid))
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.users, s.Uid)
	return nil
}

// LeaveAll clear all users in the connector
func (c *Connector) LeaveAll() error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.users = make(map[int]*GamePlayer)
	return nil
}

// Count get current member amount in the connector
func (c *Connector) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.users)
}

func (c *Connector) isClosed() bool {
	if atomic.LoadInt32(&c.status) == connectorStatusClosed {
		return true
	}
	return false
}

// Close destroy connector, which will release all resource in the connector
func (c *Connector) Close() error {
	if c.isClosed() {
		return ErrCloseClosedConnector
	}

	atomic.StoreInt32(&c.status, connectorStatusClosed)

	// release all reference
	c.users = make(map[int]*GamePlayer)
	return nil
}
