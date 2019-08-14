package game

import "fmt"

type GameTable struct {
	gsid        string
	tableid     int32
	code        int
	gsidtid     string
	player_sort map[int]int
	gameserver  *GameServer
}

func NewGameTable() *GameTable {
	gt := &GameTable{
		gsid:        "",
		tableid:     0,
		code:        0,
		gsidtid:     "",
		gameserver:  nil,
		player_sort: make(map[int]int),
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

func (gt *GameTable) addPlayer(nuid int32) {
	uid := int(nuid)
	gt.player_sort[uid] = uid
	TableManager.AddUserToTable(gt.gsidtid, uid)
	s, err := ConnectorHandler.Member(uid)
	if err == nil {
		s.Set("gsid", gt.gsid)
		s.Set("tableid", gt.tableid)
	} else {
		fmt.Println(err)
	}
}

func (gt *GameTable) RemovePlayer(nuid int32) {
	uid := int(nuid)
	delete(gt.player_sort, uid)
	TableManager.RemoveTableUser(gt.gsidtid, uid)
	s, err := ConnectorHandler.Member(uid)
	if err == nil {
		s.Remove("gsid")
		s.Remove("tableid")
	} else {
		fmt.Println(err)
	}
}
func (gt *GameTable) GetPlayerCount() int {
	return len(gt.player_sort)
}
