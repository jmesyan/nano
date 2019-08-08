package structure

type GameUserServerHs struct {
	Sid   int    `xorm:"not null pk autoincr comment('自动ID') INT(11)" json:"sid" form:"sid" csv:"sid"`
	Uid   int    `xorm:"not null default 0 comment('用户ID') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ip    string `xorm:"not null default '' comment('IP') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Isdie int    `xorm:"not null default 0 comment('死了') TINYINT(4)" json:"isdie" form:"isdie" csv:"isdie"`
	Ltime int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
