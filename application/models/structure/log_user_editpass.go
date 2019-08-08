package structure

type LogUserEditpass struct {
	Lid    int    `xorm:"not null pk autoincr INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid    int    `xorm:"not null default 0 comment('用户') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ip     string `xorm:"not null default '' comment('ip') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Device string `xorm:"not null default '' comment('设备') VARCHAR(50)" json:"device" form:"device" csv:"device"`
	Admin  int    `xorm:"not null default 0 comment('操作人') INT(11)" json:"admin" form:"admin" csv:"admin"`
	Ltime  int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Ps     string `xorm:"not null default '' VARCHAR(100)" json:"ps" form:"ps" csv:"ps"`
}
