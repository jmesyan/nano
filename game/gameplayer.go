package game

import (
	"errors"
	"fmt"
	"github.com/jmesyan/nano/session"
	"github.com/nats-io/nats.go"
	"sync"
)

type GamePlayer struct {
	Sid           int64                  `json:"sid"`            // session_id
	Uid           int                    `json:"uid"`            //用户ID
	ConnectorNid  string                 `json:"connector_nid"`  //客户端ID
	GameserverNid string                 `json:"gameserver_nid"` //服务端ID
	Channel       *GameChannel           `json:"channel"`        //游戏通道
	Data          map[string]interface{} `json:"data"`           //游戏数据
	Session       *session.Session
	client        *nats.Conn
	sync.RWMutex  // protect data
}

type GamePlayerOpt func(u *GamePlayer)

func WithPlayerSession(s *session.Session) GamePlayerOpt {
	return func(u *GamePlayer) {
		u.Sid = s.ID()
		u.Session = s
	}
}
func NewGamePlayer(uid int, connectorNid string, client *nats.Conn, opts ...GamePlayerOpt) *GamePlayer {
	u := &GamePlayer{
		Uid:          uid,
		client:       client,
		ConnectorNid: connectorNid,
		Data:         make(map[string]interface{}),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(u)
		}
	}
	return u
}

func (u *GamePlayer) NotifyConnector(topic string, data []byte) {
	ntopic := fmt.Sprintf("%s.%s", u.ConnectorNid, topic)
	err := u.client.Publish(ntopic, data)
	if err != nil {
		logger.Println(err)
	}
}

func (u *GamePlayer) SetPlayerSession(s *session.Session) {
	u.Sid = s.ID()
	u.Session = s
}

func (u *GamePlayer) SetPlayerClient(client *nats.Conn) {
	u.client = client
}

func (u *GamePlayer) SetPlayerChannel(ch *GameChannel) {
	u.Channel = ch
}

func (u *GamePlayer) SendMsg(route string, msg interface{}) error {
	if u.Session != nil {
		err := u.Session.Push(route, msg)
		return err
	}
	return errors.New("the user lose the connecter connection")
}

func (u *GamePlayer) IsPeer() bool {
	return u.ConnectorNid == ConnectorHandler.NID()
}

// Remove delete data associated with the key from session storage
func (u *GamePlayer) Remove(key string) {
	u.Lock()
	defer u.Unlock()

	delete(u.Data, key)
}

// Set associates value with the key in session storage
func (u *GamePlayer) Set(key string, value interface{}) {
	u.Lock()
	defer u.Unlock()

	u.Data[key] = value
}

// HasKey decides whether a key has associated value
func (u *GamePlayer) HasKey(key string) bool {
	u.RLock()
	defer u.RUnlock()

	_, has := u.Data[key]
	return has
}

// Int returns the value associated with the key as a int.
func (u *GamePlayer) Int(key string) int {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int)
	if !ok {
		return 0
	}
	return value
}

// Int8 returns the value associated with the key as a int8.
func (u *GamePlayer) Int8(key string) int8 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int8)
	if !ok {
		return 0
	}
	return value
}

// Int16 returns the value associated with the key as a int16.
func (u *GamePlayer) Int16(key string) int16 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int16)
	if !ok {
		return 0
	}
	return value
}

// Int32 returns the value associated with the key as a int32.
func (u *GamePlayer) Int32(key string) int32 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int32)
	if !ok {
		return 0
	}
	return value
}

// Int64 returns the value associated with the key as a int64.
func (u *GamePlayer) Int64(key string) int64 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int64)
	if !ok {
		return 0
	}
	return value
}

// Uint returns the value associated with the key as a uint.
func (u *GamePlayer) Uint(key string) uint {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(uint)
	if !ok {
		return 0
	}
	return value
}

// Uint8 returns the value associated with the key as a uint8.
func (u *GamePlayer) Uint8(key string) uint8 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(uint8)
	if !ok {
		return 0
	}
	return value
}

// Uint16 returns the value associated with the key as a uint16.
func (u *GamePlayer) Uint16(key string) uint16 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(uint16)
	if !ok {
		return 0
	}
	return value
}

// Uint32 returns the value associated with the key as a uint32.
func (u *GamePlayer) Uint32(key string) uint32 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(uint32)
	if !ok {
		return 0
	}
	return value
}

// Uint64 returns the value associated with the key as a uint64.
func (u *GamePlayer) Uint64(key string) uint64 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(uint64)
	if !ok {
		return 0
	}
	return value
}

// Float32 returns the value associated with the key as a float32.
func (u *GamePlayer) Float32(key string) float32 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(float32)
	if !ok {
		return 0
	}
	return value
}

// Float64 returns the value associated with the key as a float64.
func (u *GamePlayer) Float64(key string) float64 {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return 0
	}

	value, ok := v.(float64)
	if !ok {
		return 0
	}
	return value
}

// String returns the value associated with the key as a string.
func (u *GamePlayer) String(key string) string {
	u.RLock()
	defer u.RUnlock()

	v, ok := u.Data[key]
	if !ok {
		return ""
	}

	value, ok := v.(string)
	if !ok {
		return ""
	}
	return value
}

// Value returns the value associated with the key as a interface{}.
func (u *GamePlayer) Value(key string) interface{} {
	u.RLock()
	defer u.RUnlock()

	return u.Data[key]
}

// State returns all session state
func (u *GamePlayer) State() map[string]interface{} {
	u.RLock()
	defer u.RUnlock()

	return u.Data
}

// Restore session state after reconnect
func (u *GamePlayer) Restore(data map[string]interface{}) {
	u.Data = data
}

// Clear releases all data related to current session
func (u *GamePlayer) Clear() {
	u.Lock()
	defer u.Unlock()
	u.Data = map[string]interface{}{}
}
