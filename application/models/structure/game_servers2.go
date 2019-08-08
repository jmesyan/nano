package structure

type GameServers2 struct {
	Ip      string `xorm:"not null pk default '' comment('ip') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Depth   int    `xorm:"not null default 0 comment('深度') index INT(11)" json:"depth" form:"depth" csv:"depth"`
	Isdie   int    `xorm:"not null default 0 comment('是否死了') TINYINT(1)" json:"isdie" form:"isdie" csv:"isdie"`
	Dietime int    `xorm:"not null default 0 comment('死的时间') INT(11)" json:"dietime" form:"dietime" csv:"dietime"`
	Usetime int    `xorm:"not null default 0 comment('使用时间') INT(11)" json:"usetime" form:"usetime" csv:"usetime"`
}
