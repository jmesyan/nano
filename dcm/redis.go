package dcm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

var (
	DefaultPrefix      = "DCM:"
	StoreUserKeyprefix = "store_user_"
	StoreGameKeyprefix = "store_game_"
)

type RedisDCMOpts func(r *RedisDCM)
type RedisDCM struct {
	prefix  string
	client  *redis.Client
	options *redis.Options
}

func WithAddrs(addrs []string) RedisDCMOpts {
	return func(r *RedisDCM) {
		if r.options == nil {
			r.options = &redis.Options{}
		}
		r.options.Addr = addrs[0]
	}
}
func WithPrefix(prefix string) RedisDCMOpts {
	return func(r *RedisDCM) {
		r.prefix = prefix
	}
}
func NewRedisDCM(opts ...RedisDCMOpts) *RedisDCM {
	rs := &RedisDCM{prefix: DefaultPrefix}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(rs)
		}
	}
	//client := redis.NewClusterClient(rs.options)
	client := redis.NewClient(rs.options)
	_, err := client.Ping().Result()
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	rs.client = client
	return rs
}
func (r *RedisDCM) WrapKey(k string) string {
	return fmt.Sprintf("%s%s", r.prefix, k)
}
func (r *RedisDCM) SetValue(k string, v []byte) error {
	key := r.WrapKey(k)
	_, err := r.client.Set(key, v, 0).Result()
	return err
}

func (r *RedisDCM) GetValue(k string) (*KvPair, error) {
	key := r.WrapKey(k)
	b, err := r.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return &KvPair{Key: k, Value: b}, nil
}

func (r *RedisDCM) GetPatternKeys(prefix string) []string {
	key := fmt.Sprintf("%s_*", prefix)
	key = r.WrapKey(key)
	keys := r.client.Keys(key)
	return keys.Val()
}

func (r *RedisDCM) GetPrefixValue(k string) (*KvPairs, error) {
	keys := r.GetPatternKeys(k)
	kvs := &KvPairs{}
	for _, key := range keys {
		key = strings.TrimLeft(key, r.prefix)
		kv, err := r.GetValue(key)
		if err == nil {
			kvs.Pairs = append(kvs.Pairs, kv)
		}
	}
	if len(kvs.Pairs) == 0 {
		return nil, errors.New("empty data")
	}
	return kvs, nil
}

func (r *RedisDCM) DelValue(k string) error {
	key := r.WrapKey(k)
	_, err := r.client.Del(key).Result()
	return err
}

func (r *RedisDCM) StoreUserData(uid int, k string, v interface{}) error {
	key := fmt.Sprintf("%s_%d_%s", StoreUserKeyprefix, uid, k)
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = r.SetValue(key, value)
	return err
}

func (r *RedisDCM) GetUserData(uid int, k string, v interface{}) error {
	key := fmt.Sprintf("%s_%d_%s", StoreUserKeyprefix, uid, k)
	values, err := r.GetValue(key)
	if err != nil {
		return err
	}
	value := values.Value
	err = json.Unmarshal(value, v)
	return err
}

func (r *RedisDCM) DelUserData(uid int, k string) error {
	key := fmt.Sprintf("%s_%d_%s", StoreUserKeyprefix, uid, k)
	return r.DelValue(key)
}

func (r *RedisDCM) StoreGameData(k string, v interface{}) error {
	key := StoreGameKeyprefix + k
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = r.SetValue(key, value)
	return err
}

func (r *RedisDCM) GetGameData(k string, v interface{}) error {
	key := StoreGameKeyprefix + k
	values, err := r.GetValue(key)
	if err != nil {
		return err
	}
	value := values.Value
	err = json.Unmarshal(value, v)
	return err
}

func (r *RedisDCM) DelGameData(k string) error {
	key := StoreGameKeyprefix + k
	return r.DelValue(key)
}
