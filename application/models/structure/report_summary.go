package structure

type ReportSummary struct {
	Day           int    `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	TotalUsers    int    `xorm:"not null default 0 comment('总用户数') INT(11)" json:"total_users" form:"total_users" csv:"total_users"`
	Registers     int    `xorm:"not null default 0 comment('注册用户数') INT(11)" json:"registers" form:"registers" csv:"registers"`
	FbRegs        int    `xorm:"not null default 0 comment('fb注册用户数') INT(11)" json:"fb_regs" form:"fb_regs" csv:"fb_regs"`
	MobRegs       int    `xorm:"not null default 0 comment('手机注册用户数') INT(11)" json:"mob_regs" form:"mob_regs" csv:"mob_regs"`
	Logins        int    `xorm:"not null default 0 comment('登陆用户数') INT(11)" json:"logins" form:"logins" csv:"logins"`
	FbLogins      int    `xorm:"not null default 0 comment('fb登录用户数') INT(11)" json:"fb_logins" form:"fb_logins" csv:"fb_logins"`
	MobLogins     int    `xorm:"not null default 0 comment('手机登录用户数') INT(11)" json:"mob_logins" form:"mob_logins" csv:"mob_logins"`
	PayMoney      string `xorm:"not null default 0.00 comment('充值金额') DECIMAL(18,2)" json:"pay_money" form:"pay_money" csv:"pay_money"`
	PayUsers      int    `xorm:"not null default 0 comment('充值人数') INT(11)" json:"pay_users" form:"pay_users" csv:"pay_users"`
	Pays          int    `xorm:"not null default 0 comment('充值次数') INT(11)" json:"pays" form:"pays" csv:"pays"`
	Onlines       int    `xorm:"not null default 0 comment('最高在线人数') INT(11)" json:"onlines" form:"onlines" csv:"onlines"`
	FirstPayMoney string `xorm:"not null default 0.00 comment('首充金额') DECIMAL(18,2)" json:"first_pay_money" form:"first_pay_money" csv:"first_pay_money"`
	FirstPayUsers int    `xorm:"not null default 0 comment('首充用户数') INT(11)" json:"first_pay_users" form:"first_pay_users" csv:"first_pay_users"`
	FirstPays     int    `xorm:"not null default 0 comment('首充次数') INT(11)" json:"first_pays" form:"first_pays" csv:"first_pays"`
	TotalStones   int64  `xorm:"not null comment('宝石总数') BIGINT(20)" json:"total_stones" form:"total_stones" csv:"total_stones"`
	InviteRegs    int    `xorm:"not null default 0 INT(10)" json:"invite_regs" form:"invite_regs" csv:"invite_regs"`
	TotalGolds    int64  `xorm:"not null default 0 comment('金币总数') BIGINT(20)" json:"total_golds" form:"total_golds" csv:"total_golds"`
	Shares        int    `xorm:"not null default 0 comment('分享次数') INT(11)" json:"shares" form:"shares" csv:"shares"`
	Loginm        int    `xorm:"not null default 0 INT(11)" json:"loginm" form:"loginm" csv:"loginm"`
	Loginf        int    `xorm:"not null default 0 INT(11)" json:"loginf" form:"loginf" csv:"loginf"`
	Recu          int    `xorm:"not null INT(11)" json:"recu" form:"recu" csv:"recu"`
}
