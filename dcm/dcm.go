package dcm

import (
	"github.com/jmesyan/nano/nodes"
	"github.com/jmesyan/nano/serialize/json"
	"log"
	"os"
)

var (
	DefaultDCMAddrs = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005", "127.0.0.1:7006"}
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
	StoreUserData(uid int, k string, v interface{}) error
	GetUserData(uid int, k string, v interface{}) error
	DelUserData(uid int, k string) error
	StoreGameData(k string, v interface{}) error
	GetGameData(k string, v interface{}) error
	DelGameData(k string) error
}

var (
	DCManager  DCM
	logger     = log.New(os.Stderr, "DCManger", log.LstdFlags|log.Llongfile)
	serializer = json.NewSerializer()
)

func init() {
	DCManager = NewRedisDCM(WithAddrs(DefaultDCMAddrs))
}

func RegisterNode(nid string, node *nodes.Node) error {
	data, err := serializer.Marshal(node)
	if err != nil {
		return err
	}
	return DCManager.SetValue(nid, data)
}

func DeRegisterNode(nid string) error {
	return DCManager.DelValue(nid)
}
