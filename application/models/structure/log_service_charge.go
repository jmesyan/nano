package structure

type LogServiceCharge struct {
	Appid int   `xorm:"not null pk default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Day   int   `xorm:"not null pk default 0 comment('天') index INT(10)" json:"day" form:"day" csv:"day"`
	Gid   int   `xorm:"not null pk default 0 comment('游戏id') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Uid   int   `xorm:"not null pk index INT(10)" json:"uid" form:"uid" csv:"uid"`
	Golds int64 `xorm:"not null default 0 comment('服务费') BIGINT(20)" json:"golds" form:"golds" csv:"golds"`
	Num   int   `xorm:"not null default 0 comment('每日局数') INT(10)" json:"num" form:"num" csv:"num"`
	State int   `xorm:"not null default 1 comment('是否服务费 1是 0否') TINYINT(1)" json:"state" form:"state" csv:"state"`
}
