package game

import (
	"encoding/json"
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
	OGID_ROOMSVR_ENTERROOM       int32
	GLID_GAMEITEM_KICKTOHALL     int32
	OGID_GAME_MSG                int32
	OGID_CONTROL_CANCEL_TABLE    int32
}

func NewCmd() *cmd {
	return &cmd{
		REQ:                          0,
		ACK:                          134217728,
		OGID_MSGBASE_CONTROLBASE:     0x2500,
		OGID_CONTROL_REGIS:           9472,  //注册服务器
		OGID_CONTROL_TABLES:          9476,  //注册桌子
		OGID_CONTROL_HEART_BEAT:      9485,  //心跳
		OGID_CONTROL_CANCEL_TABLE:    9501,  //解散桌子
		OGID_CONTROL_USER_SIGN:       9622,  //金币场进入游戏
		OGID_CONTROL_DISTRIBUTE_USER: 9623,  //金币场玩家分桌
		OGID_ROOMSVR_ENTERROOM:       12801, //进入房间
		GLID_GAMEITEM_KICKTOHALL:     12372, //踢回大厅
		OGID_GAME_MSG:                12315, //玩家进出游戏
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
func GetGameParamsByGsidtid(gsidtid string) (gid, rtype, ridx, tid int) {
	gsids := strings.Split(gsidtid, "_")
	if len(gsids) == 4 {
		return utils.StringToInt(gsids[0]), utils.StringToInt(gsids[1]), utils.StringToInt(gsids[2]), utils.StringToInt(gsids[3])
	}
	if len(gsids) == 3 {
		return utils.StringToInt(gsids[0]), utils.StringToInt(gsids[2]), utils.StringToInt(gsids[3]), 0
	}
	if len(gsids) == 2 {
		return utils.StringToInt(gsids[0]), utils.StringToInt(gsids[2]), 0, 0
	}
	return utils.StringToInt(gsids[0]), 0, 0, 0
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

func MakeGameMsg(cmd int32, msg map[string]interface{}) (string, error) {
	cmds := fmt.Sprintf("000000000%d", cmd)
	cmds = cmds[len(cmds)-9 : 0]
	data, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return cmds + string(data), nil
}

type GameService interface {
	ProcessServer(route string, body reflect.Value)
	RegisterServer(gsid string, server *GameServer)
	GetServerByGSID(gsid string) *GameServer
	RemoveServerByGSID(gsid string)
	GetServerSort() map[string]*GameServer
}
