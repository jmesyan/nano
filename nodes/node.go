package nodes

import (
	"fmt"
	"time"
)

var (
	NodeTypesToKind = map[NodeType]string{
		NodeMaster:     "master",
		NodeConnector:  "connector",
		NodeGameServer: "gameserver",
	}
	NodeKindToTypes = map[string]NodeType{
		"master":     NodeMaster,
		"connector":  NodeConnector,
		"gameserver": NodeGameServer,
	}

	NodeTypesToStatus = map[NodeStatus]string{
		NodeStoped:   "stoped",
		NodeStoping:  "stoping",
		NodeStarting: "starting",
		NodeStarted:  "started",
	}
	NodeStatusToTypes = map[string]NodeStatus{
		"stoped":   NodeStoped,
		"stoping":  NodeStoping,
		"starting": NodeStarting,
		"started":  NodeStarted,
	}
)

type NodeType int

const (
	_ NodeType = iota
	NodeMaster
	NodeConnector
	NodeGameServer
)

type NodeStatus int

const (
	_ NodeStatus = iota
	NodeStoped
	NodeStoping
	NodeStarting
	NodeStarted
)

type Node struct {
	Nid        string
	Name       string
	Address    string
	Host       string
	Port       int
	Type       NodeType
	Status     NodeStatus
	UpInterval time.Duration
}
type NodeOpts func(n *Node)

func WithNodeStatus(status NodeStatus) NodeOpts {
	return func(n *Node) {
		n.Status = status
	}
}

func WithNodeUpInterval(d time.Duration) NodeOpts {
	return func(n *Node) {
		n.UpInterval = d
	}
}
func (n *Node) GenerateNodeId(kind, address string) string {
	return fmt.Sprintf("%s_%s", kind, address)
}
func (n *Node) GenerateAddress() {
	n.Address = fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func NewNode(name, host string, port int, ntype NodeType, opts ...NodeOpts) *Node {
	node := &Node{
		Name:       name,
		Host:       host,
		Port:       port,
		Type:       ntype,
		Status:     NodeStarting,
		UpInterval: time.Second,
	}
	node.GenerateAddress()
	node.Nid = node.GenerateNodeId(NodeTypesToKind[ntype], node.Address)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(node)
		}
	}
	return node
}

func (n *Node) SetNodeStatus(nstatus NodeStatus, opt NodeOpts) {
	n.Status = nstatus
	opt(n)
}
