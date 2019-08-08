package structure

type LogDayLogin struct {
	Day int `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Uid int `xorm:"not null pk default 0 comment('登陆用户UID') INT(11)" json:"uid" form:"uid" csv:"uid"`
}
