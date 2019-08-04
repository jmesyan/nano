package discovery

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"hewolf/cluster/datacenter"
	"hewolf/cluster/nodes"
	"hewolf/config"
	"time"
)

type ConsulDiscovery struct {
	client *consulapi.Client
}

func NewConsulDiscovery() *ConsulDiscovery {
	conf := consulapi.DefaultConfig()
	conf.Address = config.CONSUL_ADDRESS
	client, err := consulapi.NewClient(conf)
	if err != nil {
		panic(err)
	}
	return &ConsulDiscovery{
		client: client,
	}
}

func (d *ConsulDiscovery) UpdateNode(node *nodes.Node) error {
	reg := &consulapi.AgentServiceRegistration{
		ID:      node.Address,
		Name:    node.Name,
		Address: node.Host,
		Port:    node.Port,
		Meta:    map[string]string{"type": nodes.NodeTypesToKind[node.Type], "status": nodes.NodeTypesToStatus[node.Status]},
	}
	if err := d.client.Agent().ServiceRegister(reg); err != nil {
		return err
	}
	return nil
}

func (d *ConsulDiscovery) RegisterNode(node *nodes.Node) error {
	reg := &consulapi.AgentServiceRegistration{
		ID:      node.Address,
		Name:    node.Name,
		Address: node.Host,
		Port:    node.Port,
		Meta:    map[string]string{"type": nodes.NodeTypesToKind[node.Type], "status": nodes.NodeTypesToStatus[node.Status]},
	}
	if err := d.client.Agent().ServiceRegister(reg); err != nil {
		panic(err)
	}
	// initial register service check
	check := consulapi.AgentServiceCheck{TTL: fmt.Sprintf("%ds", config.CONSUL_TTL), Status: consulapi.HealthPassing}
	err := d.client.Agent().CheckRegister(&consulapi.AgentCheckRegistration{
		ID:                node.Address,
		Name:              node.Name,
		ServiceID:         node.Address,
		AgentServiceCheck: check,
	})
	if err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(node.UpInterval)
		for {
			<-ticker.C
			err = d.client.Agent().UpdateTTL(node.Address, "", check.Status)
			if err != nil {
				logger.Println("update ttl of service error: ", err.Error())
			}
		}
	}()

	return nil
}

func (d *ConsulDiscovery) DeRegister(id string) error {
	err := d.client.Agent().ServiceDeregister(id)
	if err != nil {
		logger.Println("deregister service error: ", err.Error())
	} else {
		logger.Println("deregistered service from consul server.")
	}

	err = d.client.Agent().CheckDeregister(id)
	if err != nil {
		logger.Println("deregister check error: ", err.Error())
	}

	return nil
}

func (d *ConsulDiscovery) TraverseServices(f TraverseHandler) error {
	services, err := d.client.Agent().Services()
	if err != nil {
		return err
	}
	for _, v := range services {
		s := &TraverseService{
			Service: v.Service,
			Tags:    v.Tags,
			Meta:    v.Meta,
			Port:    v.Port,
			Address: v.Address,
		}
		f(d, s)
	}
	return nil
}

func (d *ConsulDiscovery) SetValue(k string, v []byte) error {
	kv := d.client.KV()
	kvpair := &consulapi.KVPair{Key: k, Value: v}
	if _, err := kv.Put(kvpair, nil); err != nil {
		return err
	}
	return nil
}

func (d *ConsulDiscovery) GetValue(k string) (*datacenter.KvPair, error) {
	kv := d.client.KV()
	pair, _, err := kv.Get(k, nil)
	if err != nil {
		return nil, err
	}
	if pair == nil {
		return nil, nil
	}
	return &datacenter.KvPair{Key: pair.Key, Value: pair.Value}, nil
}

func (d *ConsulDiscovery) GetPrefixValue(k string) (*datacenter.KvPairs, error) {
	kv := d.client.KV()
	pairs, _, err := kv.List(k, nil)
	if err != nil {
		return nil, err
	}
	if pairs == nil {
		return nil, nil
	}
	list := new(datacenter.KvPairs)
	for _, pair := range pairs {
		list.Pairs = append(list.Pairs, &datacenter.KvPair{Key: pair.Key, Value: pair.Value})
	}
	return list, nil
}
