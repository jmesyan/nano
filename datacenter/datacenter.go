package datacenter

type KvPair struct {
	Key   string
	Value []byte
}

type KvPairs struct {
	Pairs []*KvPair
}

type DataCenter interface {
	SetValue(k string, v []byte) error
	GetValue(k string) (*KvPair, error)
	GetPrefixValue(k string) (*KvPairs, error)
}
