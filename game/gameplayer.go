package game

import (
	"errors"
	"github.com/jmesyan/nano/session"
	"github.com/nats-io/nats.go"
)

type GamePlayer struct {
	Uid        int
	Chan       *GameChannel
	Sess       *session.Session
	client     *nats.Conn
	clientAddr string
}

type GamePlayerOpt func(u *GamePlayer)

func WithPlayerSession(s *session.Session) GamePlayerOpt {
	return func(u *GamePlayer) {
		u.Sess = s
	}
}
func NewGamePlayer(uid int, clientAddr string, client *nats.Conn, opts ...GamePlayerOpt) *GamePlayer {
	u := &GamePlayer{
		Uid:        uid,
		client:     client,
		clientAddr: clientAddr,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(u)
		}
	}
	return u
}

func (u *GamePlayer) SetPlayerSession(s *session.Session) {
	u.Sess = s
}

func (u *GamePlayer) SetPlayerClient(client *nats.Conn) {
	u.client = client
}

func (u *GamePlayer) SetPlayerChannel(ch *GameChannel) {
	u.Chan = ch
}

func (u *GamePlayer) SendMsg(route string, msg interface{}) error {
	if u.Sess != nil {
		err := u.Sess.Push(route, msg)
		return err
	}
	return errors.New("the user lose the connecter connection")
}
