package game

import "fmt"

type GameTable struct {
	gsid       string
	tableid    int32
	code       int
	gsidtid    string
	gameserver *GameServer
}

func NewGameTable() *GameTable {
	gt := &GameTable{
		gsid:       "",
		tableid:    0,
		code:       0,
		gsidtid:    "",
		gameserver: nil,
	}
	return gt
}

func (gt *GameTable) Init(gsid string, table *ControlRoomUsersTableInfo) {
	if table != nil {
		gt.gsid = gsid
		gt.tableid = table.GetTid()
		gt.gsidtid = fmt.Sprintf("%s_%d", gsid, gt.tableid)
		uids := table.GetUid()
		if len(uids) > 0 {
			for _, uid := range uids {
				gt.addPlayer(uid)
			}
		}
	}
}

func (gt *GameTable) addPlayer(uid int32) {

}
