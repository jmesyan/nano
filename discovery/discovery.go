package discovery

import (
	"github.com/sirupsen/logrus"
	"hewolf/app"
	"hewolf/cluster/nodes"
)

var (
	Discover Discovery
	logger   *logrus.Entry
)

type TraverseService struct {
	Service string
	Tags    []string
	Meta    map[string]string
	Port    int
	Address string
}

type TraverseHandler func(Discovery, *TraverseService)

type Discovery interface {
	RegisterNode(node *nodes.Node) error
	UpdateNode(node *nodes.Node) error
	DeRegister(id string) error
	TraverseServices(f TraverseHandler) error
}

func init() {
	Discover = NewConsulDiscovery()
	logger = app.FilesystemLogger("discovery", "discovery")
}
