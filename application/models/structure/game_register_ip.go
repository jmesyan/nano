package structure

type GameRegisterIp struct {
	Day   int    `xorm:"not null pk comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Ip    string `xorm:"not null pk default '' comment('ip') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Count int    `xorm:"not null default 0 comment('注册用户数') INT(11)" json:"count" form:"count" csv:"count"`
}
