package game

import (
	"fmt"
	"github.com/jmesyan/nano/utils"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

var (
	CMD = NewCmd()
)
var (
	logger       = log.New(os.Stderr, "game", log.LstdFlags|log.Llongfile)
	serversort   = make(map[string]*GameServer)
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
	}
}

func GetServerByGSID(gsid string) *GameServer {
	if server, ok := serversort[gsid]; ok {
		return server
	}
	return nil
}

func GetCenterServerByBalance(ngid int) *GameServer {
	config, ok := gds.Configs[ngid]
	if !ok {
		return nil
	}
	gcid := config.Censerver
	gsids := make(map[int]string)
	for gsid, _ := range serversort {
		gid, rtype, _ := GetGameParamsByGsid(gsid)
		grid := GetGrid(gid, rtype)
		if grid == gcid && !IsServerMaintence(gsid) {
			gsids[gds.Gcsu[gsid]] = gsid
		}
	}
	if len(gsids) > 0 {
		gsorts := make([]int, len(gsids))
		for k, _ := range gsids {
			gsorts = append(gsorts, k)
		}
		sort.Ints(gsorts)
		gsid := gsids[gsorts[0]]
		return serversort[gsid]
	}
	return nil
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

type HandlerService interface {
	ProcessServer(route string, body reflect.Value)
}
