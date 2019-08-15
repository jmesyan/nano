package game

import (
	"errors"
	"github.com/jmesyan/nano/session"
)

type GamePlayer struct {
	Uid  int
	Chan *GameChannel
	Sess *session.Session
}

type GamePlayerOpt func(u *GamePlayer)

func NewGamePlayer(uid int, opts ...GamePlayerOpt) *GamePlayer {
	u := &GamePlayer{Uid: uid}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(u)
		}
	}
	return u
}

func (u *GamePlayer) SendMsg(route string, msg interface{}) error {
	if u.Sess != nil {
		err := u.Sess.Push(route, msg)
		return err
	}
	return errors.New("the user lose the connecter connection")
}
