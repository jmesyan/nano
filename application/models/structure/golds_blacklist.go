package structure

type GoldsBlacklist struct {
	Uid         int `xorm:"not null pk INT(10)" json:"uid" form:"uid" csv:"uid"`
	Ltime       int `xorm:"not null comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Cheatvalues int `xorm:"not null default 0 comment('作弊值') INT(10)" json:"cheatvalues" form:"cheatvalues" csv:"cheatvalues"`
}
