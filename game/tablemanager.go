package game

import (
	"fmt"
	"github.com/jmesyan/nano/application/cache"
	"github.com/jmesyan/nano/application/stores"
	"github.com/jmesyan/nano/utils"
	"strings"
)

var (
	TableManager *RandomAssignGameTable
	p2p          = stores.StoresHandler.P2p
)

type RandomAssignGameTable struct {
}

func NewTableManager() *RandomAssignGameTable {
	return &RandomAssignGameTable{}
}

func (tm *RandomAssignGameTable) registerTables(gsid string, tables []*ControlRoomUsersTableInfo) {
	gsids := strings.Split(gsid, "_")
	gid := utils.StringToInt(gsids[0])
	if len(tables) > 0 {
		delete(sys.MAINTEN_SERVERS, fmt.Sprintf("SYS_MAINTENANCE_%s", gsid))
		cache.CacheManager.RemoveGameManintence(gsid)
		codeCmp := func(a interface{}, b interface{}) bool {
			c1 := a.(int32)
			c2 := b.(int32)
			return c1 == c2
		}
		allcodes := make([]interface{}, len(p2p.Mj.AllCode))
		for k, v := range p2p.Mj.AllCode {
			allcodes[k] = v
		}
		use, tids := 0, []int32{}
		for _, table := range tables {
			tid, code, ownid, uids := table.GetTid(), table.GetCode(), table.GetOwnerid(), table.GetUid()
			gsidtid := fmt.Sprintf("%s_%d", gsid, tid)
			if code > 0 && ownid > 0 {
				use++
				if !utils.InArray(gsidtid, p2p.Mj.Use) {
					p2p.Mj.Use = append(p2p.Mj.Use, gsidtid)
				}
				if !utils.InArray(code, p2p.Mj.Code) {
					p2p.Mj.Code[code] = gsidtid
					p2p.Mj.CodeSort[gsidtid] = code
				}
				if len(allcodes) > 0 {
					for index := utils.IndexOf(allcodes, code, codeCmp); index != -1; {
						p2p.Mj.AllCode = append(p2p.Mj.AllCode[:index], p2p.Mj.AllCode[index+1:]...)
						allcodes = append(allcodes[:index], allcodes[index+1:]...)
					}
				}
				tm.addUserToTable(gsidtid, int(ownid))
				if len(uids) > 0 {
					for _, uid := range uids {
						tm.addUserToTable(gsidtid, int(uid))
						//在线进行重连
					}
				}
			} else {
				if !utils.InArray(gsidtid, p2p.Mj.Nouse[gid]) {
					p2p.Mj.Nouse[gid] = append(p2p.Mj.Nouse[gid], gsidtid)
				}
			}
			tids = append(tids, tid)
		}
		logger.Printf("registerTables %s 注册:%d,使用:%d,总使用桌子:%d,总可用桌子:%d,总可用房号:%d\n", gsid, len(tables), use, len(p2p.Mj.Use), len(p2p.Mj.Nouse[gid]), len(p2p.Mj.AllCode))
	}

}

func (tm *RandomAssignGameTable) addUserToTable(gsidtid string, uid int) {
	if len(gsidtid) == 0 || uid < 1 {
		return
	}
	if !utils.InArray(uid, p2p.Mj.Gsidtid[gsidtid]) {
		p2p.Mj.Gsidtid[gsidtid] = append(p2p.Mj.Gsidtid[gsidtid], uid)
	}
}

func init() {
	TableManager = NewTableManager()
}
