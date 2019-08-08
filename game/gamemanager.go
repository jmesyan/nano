package game

import (
	"log"
	"os"
)

var (
	BASE uint32 = 0x2500
	ACK  int32  = 134217728
	REQ  uint32 = 0
)
var (
	logger     = log.New(os.Stderr, "game", log.LstdFlags|log.Llongfile)
	serversort = make(map[string]*GameServer)
)
