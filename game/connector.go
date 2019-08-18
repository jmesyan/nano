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
	"github.com/jmesyan/nano/utils"
	"github.com/nats-io/nats.go"
	"sync"
	"sync/atomic"
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
	status     int32 // channel current status
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

func (c *Connector) watcher() {
	for {
		select {
		case msg := <-c.msgch:
			UMHandler.DoConnectorMsg(c, msg)
		case <-c.shut:
			logger.Println("receive stop msg")
			close(c.msgch)
			c.node.Status = nodes.NodeStoping
			return
		}
	}
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
	UMHandler.LeaveAll()
	return nil
}
