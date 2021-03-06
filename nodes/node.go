package nodes

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
	Nid     string     `json:"nid"`
	Name    string     `json:"name"`
	Gsid    string     `json:"gsid"`
	Address string     `json:"address"`
	Type    NodeType   `json:"type"`
	Status  NodeStatus `json:"status"`
}
type NodeOpts func(n *Node)

func WithNodeID(nid string) NodeOpts {
	return func(n *Node) {
		n.Nid = nid
	}
}
func WithNodeAddress(addr string) NodeOpts {
	return func(n *Node) {
		n.Address = addr
	}
}
func WithNodeGsid(gsid string) NodeOpts {
	return func(n *Node) {
		n.Gsid = gsid
	}
}
func WithNodeStatus(status NodeStatus) NodeOpts {
	return func(n *Node) {
		n.Status = status
	}
}

func NewNode(name, nid string, ntype NodeType, opts ...NodeOpts) *Node {
	node := &Node{
		Name:   name,
		Nid:    nid,
		Type:   ntype,
		Status: NodeStarted,
	}
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
