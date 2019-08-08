package structure

type CreateRoomConfig struct {
	Gid       int    `xorm:"not null pk comment('游戏') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Gameset   string `xorm:"comment('03CRET指令JSON体数据 ') VARCHAR(500)" json:"gameset" form:"gameset" csv:"gameset"`
	Roundset  string `xorm:"comment('03CRET指令|分隔数据  0普通 1AA 2大赢家 3代开 4 分时段免费') VARCHAR(255)" json:"roundset" form:"roundset" csv:"roundset"`
	Starttime int    `xorm:"not null default 0 INT(11)" json:"startTime" form:"startTime" csv:"startTime"`
	Endtime   int    `xorm:"not null default 0 INT(11)" json:"endTime" form:"endTime" csv:"endTime"`
}
