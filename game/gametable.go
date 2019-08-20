package game

import "fmt"

type GameTable struct {
	Gsid        string      `json:"gsid"`
	Tableid     int32       `json:"tableid"`
	Code        int         `json:"code"`
	Gsidtid     string      `json:"gsidtid"`
	Player_sort map[int]int `json:"player_sort"`
}

func NewGameTable() *GameTable {
	gt := &GameTable{
		Gsid:        "",
		Tableid:     0,
		Code:        0,
		Gsidtid:     "",
		Player_sort: make(map[int]int),
	}
	return gt
}

func (gt *GameTable) Init(Gsid string, table *ControlRoomUsersTableInfo) {
	if table != nil {
		gt.Gsid = Gsid
		gt.Tableid = table.GetTid()
		gt.Gsidtid = fmt.Sprintf("%s_%d", Gsid, gt.Tableid)
		uids := table.GetUid()
		if len(uids) > 0 {
			for _, uid := range uids {
				gt.AddPlayer(uid)
			}
		}
	}
}

func (gt *GameTable) AddPlayer(nuid int32) {
	uid := int(nuid)
	gt.Player_sort[uid] = uid
	TableManager.AddUserToTable(gt.Gsidtid, uid)
	u, err := UMHandler.Member(uid)
	if err == nil {
		u.Set("Gsid", gt.Gsid)
		u.Set("Tableid", gt.Tableid)
	} else {
		fmt.Println(err)
	}
}

func (gt *GameTable) RemovePlayer(nuid int32) {
	uid := int(nuid)
	delete(gt.Player_sort, uid)
	TableManager.RemoveTableUser(gt.Gsidtid, uid)
	u, err := UMHandler.Member(uid)
	if err == nil {
		u.Remove("Gsid")
		u.Remove("Tableid")
	} else {
		fmt.Println(err)
	}
}
func (gt *GameTable) GetPlayerCount() int {
	return len(gt.Player_sort)
}

func (gt *GameTable) Dispose() {
	for uid, _ := range gt.Player_sort {
		gt.RemovePlayer(int32(uid))
	}
	gt.Code = 0
}
