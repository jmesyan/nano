package structure

type ReportGoldsSummary struct {
	Day   int    `xorm:"not null pk default 0 comment('日期ymd') INT(11)" json:"day" form:"day" csv:"day"`
	Gid   int    `xorm:"not null pk default 0 comment('游戏ID') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype int    `xorm:"not null pk default 0 comment('房间类型') INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	State int    `xorm:"not null pk default 0 comment('状态') INT(11)" json:"state" form:"state" csv:"state"`
	Golds string `xorm:"not null default 0.00 comment('金币') DECIMAL(18,2)" json:"golds" form:"golds" csv:"golds"`
}
