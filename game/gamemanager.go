package game

import (
	"log"
	"os"
)

var (
	CMD = NewCmd()
)
var (
	logger       = log.New(os.Stderr, "game", log.LstdFlags|log.Llongfile)
	serversort   = make(map[string]*GameServer)
	alltablesort = make(map[int32]*GameTable)
)

type cmd struct {
	REQ                      int32
	ACK                      int32
	OGID_MSGBASE_CONTROLBASE int32
	OGID_CONTROL_REGIS       int32
	OGID_CONTROL_TABLES      int32
	OGID_CONTROL_HEART_BEAT  int32
}

func NewCmd() *cmd {
	return &cmd{
		REQ:                      0,
		ACK:                      134217728,
		OGID_MSGBASE_CONTROLBASE: 0x2500,
		OGID_CONTROL_REGIS:       9472, //注册服务器
		OGID_CONTROL_TABLES:      9476, //注册桌子
		OGID_CONTROL_HEART_BEAT:  9485, //心跳
	}
}
