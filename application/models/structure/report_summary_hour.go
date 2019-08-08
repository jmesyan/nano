package structure

type ReportSummaryHour struct {
	Day       int     `xorm:"not null pk default 0 comment('ymd') INT(11)" json:"day" form:"day" csv:"day"`
	H         int     `xorm:"not null pk default 0 comment('hour') TINYINT(3)" json:"h" form:"h" csv:"h"`
	Registers int     `xorm:"not null default 0 comment('注册用户数') INT(11)" json:"registers" form:"registers" csv:"registers"`
	FbRegs    int     `xorm:"not null default 0 comment('fb注册用户数') INT(11)" json:"fb_regs" form:"fb_regs" csv:"fb_regs"`
	MobRegs   int     `xorm:"not null default 0 comment('手机注册用户数') INT(11)" json:"mob_regs" form:"mob_regs" csv:"mob_regs"`
	Logins    int     `xorm:"not null default 0 comment('登陆用户数') INT(11)" json:"logins" form:"logins" csv:"logins"`
	FbLogins  int     `xorm:"not null default 0 comment('fb登录用户数') INT(11)" json:"fb_logins" form:"fb_logins" csv:"fb_logins"`
	MobLogins int     `xorm:"not null default 0 comment('手机登录用户数') INT(11)" json:"mob_logins" form:"mob_logins" csv:"mob_logins"`
	PayMoney  float32 `xorm:"not null default 0 comment('充值金额') FLOAT" json:"pay_money" form:"pay_money" csv:"pay_money"`
	PayUsers  int     `xorm:"not null default 0 comment('充值人数') INT(11)" json:"pay_users" form:"pay_users" csv:"pay_users"`
	Pays      int     `xorm:"not null default 0 comment('充值次数') INT(11)" json:"pays" form:"pays" csv:"pays"`
	Onlines   int     `xorm:"not null default 0 comment('最高在线人数') INT(11)" json:"onlines" form:"onlines" csv:"onlines"`
}
