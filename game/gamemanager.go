package game

import (
	"fmt"
	"github.com/jmesyan/nano/utils"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	CMD = NewCmd()
)
var (
	logger       = log.New(os.Stderr, "game", log.LstdFlags|log.Llongfile)
	alltablesort = make(map[int32]*GameTable)
	ticker       = 0
)

type cmd struct {
	REQ                          int32
	ACK                          int32
	OGID_MSGBASE_CONTROLBASE     int32
	OGID_CONTROL_REGIS           int32
	OGID_CONTROL_TABLES          int32
	OGID_CONTROL_HEART_BEAT      int32
	OGID_CONTROL_USER_SIGN       int32
	OGID_CONTROL_DISTRIBUTE_USER int32
	OGID_ROOMSVR_ENTERROOM int32
}

func NewCmd() *cmd {
	return &cmd{
		REQ:                          0,
		ACK:                          134217728,
		OGID_MSGBASE_CONTROLBASE:     0x2500,
		OGID_CONTROL_REGIS:           9472, //注册服务器
		OGID_CONTROL_TABLES:          9476, //注册桌子
		OGID_CONTROL_HEART_BEAT:      9485, //心跳
		OGID_CONTROL_USER_SIGN:       9622, //金币场进入游戏
		OGID_CONTROL_DISTRIBUTE_USER: 9623, //金币场玩家分桌
		OGID_ROOMSVR_ENTERROOM:12801,//进入房间
	}
}

func GetGameParamsByGsid(gsid string) (gid, rtype, ridx int) {
	gsids := strings.Split(gsid, "_")
	if len(gsids) == 3 {
		return utils.StringToInt(gsids[0]), utils.StringToInt(gsids[1]), utils.StringToInt(gsids[2])
	}
	if len(gsids) == 2 {
		return utils.StringToInt(gsids[0]), utils.StringToInt(gsids[2]), 0
	}
	return utils.StringToInt(gsid), 0, 0
}

func GetGrid(gid, rtype int) string {
	return fmt.Sprintf("%d_%d", gid, rtype)
}

func IsServerMaintence(gsid string) bool {
	return sys.MAINTEN_SERVERS[fmt.Sprintf("SYS_MAINTENANCE_%s", gsid)]
}

func RemoveServerManintence(gsid string) {
	delete(sys.MAINTEN_SERVERS, fmt.Sprintf("SYS_MAINTENANCE_%s", gsid))
}

type GameService interface {
	ProcessServer(route string, body reflect.Value)
	RegisterServer(gsid string, server *GameServer)
	GetServerByGSID(gsid string) *GameServer
	RemoveServerByGSID(gsid string)
	GetServerSort() map[string]*GameServer
}
