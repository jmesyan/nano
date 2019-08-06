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

package nano

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/dcm"
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/session"
	"github.com/jmesyan/nano/users"
	"github.com/nats-io/nats.go"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ConnectorHandler *Connector
	ResponseSuccess  = []byte("SUCCESS")
	ResponseFail     = []byte("FAIL")
)

const (
	connectorStatusWorking = 0
	connectorStatusClosed  = 1
)

// SessionFilter represents a filter which was used to filter session when Multicast,
// the session will receive the message while filter returns true.
type SessionFilter func(*session.Session) bool

// Connector represents a session connector which used to manage a number of
// sessions, data send to the connector will send to all session in it.
type Connector struct {
	node      *nodes.Node
	mu        sync.RWMutex
	status    int32                      // channel current status
	sessions  map[int64]*session.Session // session id map to session instance
	listen    string
	client    *nats.Conn
	msgch     chan *nats.Msg
	shut      chan struct{}
	kickTopic string
	pushTopic string
}

type ConnectorOpts func(g *Connector)

func WithListen(address string) ConnectorOpts {
	return func(c *Connector) {
		c.listen = address
	}
}

// StartConnector start a new connector instance
func NewConnector(opts ...ConnectorOpts) *Connector {
	c := &Connector{
		status:   connectorStatusWorking,
		sessions: make(map[int64]*session.Session),
		listen:   nats.DefaultURL,
		msgch:    make(chan *nats.Msg, 64),
		shut:     make(chan struct{}, 1),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(c)
		}
	}
	ConnectorHandler = c
	return c
}

func (c *Connector) NID() string {
	return c.node.Nid
}

func (c *Connector) Status() nodes.NodeStatus {
	return c.node.Status
}

func (c *Connector) Init() {
	var err error
	nid := generateNodeId(nodes.NodeConnector, "")
	n := nodes.NewNode("connector", nid, nodes.NodeConnector, nodes.WithNodeAddress(generateLocalAddr()))
	dcm.RegisterNode(nid, n)
	c.node = n
	c.client, err = nats.Connect(c.listen)
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
	c.kickTopic = generateTopic(c.node.Nid, "kick")
	c.pushTopic = generateTopic(c.node.Nid, "push")
}

func (c *Connector) AfterInit() {
	go c.watcher()
}

// Member returns specified UID's session
func (c *Connector) Member(uid int) (*session.Session, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	nuid := int64(uid)
	for _, s := range c.sessions {
		if s.UID() == nuid {
			return s, nil
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
	payload := &MsgLoad{}
	err := serializer.Unmarshal(msg.Data, payload)
	if err != nil {
		logger.Println(err)
		msg.Respond(ResponseFail)
		return
	}
	uid := payload.Receiver.Uid
	sid := payload.Receiver.Sid
	nid := payload.Receiver.Nid
	data := payload.Msg
	switch msg.Subject {
	case c.kickTopic:
		//收到踢人消息
		state := 0
		if tmp, ok := data["state"]; ok {
			state = int(tmp.(float64))
		}
		if nid == c.node.Nid {
			sess, err := c.Member(uid)
			if err != nil {
				logger.Println(err)
				msg.Respond(ResponseFail)
				return
			}
			if sid == sess.ID() {
				sess.Push("quit", map[string]interface{}{"state": state, "id": sid})
				sess.Clear()
				if state != 1 {
					sess.Close()
				}
				c.RemoveUser(uid)
			}
		}
	case c.pushTopic:
		if nid == c.node.Nid {
			sess, err := c.Member(uid)
			if err != nil {
				logger.Println(err)
				msg.Respond(ResponseFail)
				return
			}
			if sid == sess.ID() {
				err = sess.Push(payload.Route, payload.Msg)
				if err != nil {
					logger.Println(err)
				}
			}
		}
	}
	msg.Respond(ResponseSuccess)
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
	topic := generateTopic(connector, "push")
	payload := &MsgLoad{Receiver: receiver, Route: route, Msg: data}
	load, err := serializer.Marshal(payload)
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
	topic := generateTopic(connector, "kick")
	payload := &MsgLoad{Receiver: receiver, Msg: map[string]interface{}{"state": state}}
	load, err := serializer.Marshal(payload)
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
	user, err := serializer.Marshal(data)
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
	err = serializer.Unmarshal(kv.Value, user)
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
	for _, s := range c.sessions {
		members = append(members, s.UID())
	}

	return members
}

// Multicast  push  the message to the filtered clients
func (c *Connector) Multicast(route string, v interface{}, filter SessionFilter) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	data, err := serializeOrRaw(v)
	if err != nil {
		return err
	}

	if env.debug {
		logger.Println(fmt.Sprintf("Type=Multicast Route=%s, Data=%+v", route, v))
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, s := range c.sessions {
		if !filter(s) {
			continue
		}
		if err = s.Push(route, data); err != nil {
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

	data, err := serializeOrRaw(v)
	if err != nil {
		return err
	}

	if env.debug {
		logger.Println(fmt.Sprintf("Type=Broadcast Route=%s, Data=%+v", route, v))
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, s := range c.sessions {
		if err = s.Push(route, data); err != nil {
			logger.Println(fmt.Sprintf("Session push message error, ID=%d, UID=%d, Error=%s", s.ID(), s.UID(), err.Error()))
		}
	}

	return err
}

// Contains check whether a UID is contained in current connector or not
func (c *Connector) Contains(uid int) bool {
	_, err := c.Member(uid)
	return err == nil
}

// Add add session to connector
func (c *Connector) Add(s *session.Session) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	if env.debug {
		logger.Println(fmt.Sprintf("Add session to connector, ID=%d, UID=%d", s.ID(), s.UID()))
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	id := s.ID()
	_, ok := c.sessions[s.ID()]
	if ok {
		return ErrSessionDuplication
	}

	c.sessions[id] = s
	return nil
}

// Leave remove specified UID related session from connector
func (c *Connector) Leave(s *session.Session) error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	if env.debug {
		logger.Println(fmt.Sprintf("Remove session from connector, UID=%d", s.UID()))
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.sessions, s.ID())
	return nil
}

// LeaveAll clear all sessions in the connector
func (c *Connector) LeaveAll() error {
	if c.isClosed() {
		return ErrClosedConnector
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.sessions = make(map[int64]*session.Session)
	return nil
}

// Count get current member amount in the connector
func (c *Connector) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.sessions)
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
	c.sessions = make(map[int64]*session.Session)
	return nil
}

func init() {
	ConnectorHandler = NewConnector()
	Register(ConnectorHandler)
}
