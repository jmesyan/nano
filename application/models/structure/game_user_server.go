package structure

type GameUserServer struct {
	Uid       int    `xorm:"not null pk default 0 comment('用户') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ip        string `xorm:"not null default '' comment('ip') index VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Enters    int    `xorm:"not null default 0 comment('进入次数') INT(11)" json:"enters" form:"enters" csv:"enters"`
	Isdie     int    `xorm:"not null default 0 comment('死了') TINYINT(4)" json:"isdie" form:"isdie" csv:"isdie"`
	Dietime   int    `xorm:"not null default 0 comment('时间') INT(11)" json:"dietime" form:"dietime" csv:"dietime"`
	Entertime int    `xorm:"not null default 0 comment('进入时间') INT(11)" json:"entertime" form:"entertime" csv:"entertime"`
}
