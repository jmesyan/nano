package structure

type LogCreateUserGoldLiushui struct {
	Lid         int   `xorm:"not null pk autoincr index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid         int   `xorm:"not null default 0 index INT(10)" json:"uid" form:"uid" csv:"uid"`
	Gid         int   `xorm:"not null default 0 INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rid         int   `xorm:"not null default 0 INT(11)" json:"rid" form:"rid" csv:"rid"`
	Goldschange int64 `xorm:"not null default 0 comment('金币变化') BIGINT(20)" json:"goldschange" form:"goldschange" csv:"goldschange"`
	Ltime       int   `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Testday     int   `xorm:"not null default 0 comment('测试天数') INT(11)" json:"testday" form:"testday" csv:"testday"`
}
