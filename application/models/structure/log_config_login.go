package structure

type LogConfigLogin struct {
	Lid   int    `xorm:"not null pk autoincr comment('日志') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid   int    `xorm:"not null default 0 comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ip    string `xorm:"not null default '' comment('ip') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Ltime int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
