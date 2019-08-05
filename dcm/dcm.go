package dcm

import (
	"github.com/jmesyan/nano/nodes"
	"log"
	"os"
)

type KvPair struct {
	Key   string
	Value []byte
}

type KvPairs struct {
	Pairs []*KvPair
}

type DCM interface {
	SetValue(k string, v []byte) error
	GetValue(k string) (*KvPair, error)
	DelValue(k string) error
	GetPrefixValue(k string) (*KvPairs, error)
}

var (
	DCManager DCM
	logger    = log.New(os.Stderr, "DCManger", log.LstdFlags|log.Llongfile)
)

func init() {
	DCManager = NewRedisDCM(WithAddrs([]string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005", "127.0.0.1:7006"}))
}

func RegisterNode(nid string, node *nodes.Node) error {

	DCManager.SetValue(nid)
}

func DeRegisterNode(nid string) error {
	return DCManager.DelValue(nid)
}
