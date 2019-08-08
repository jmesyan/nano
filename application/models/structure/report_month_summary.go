package structure

type ReportMonthSummary struct {
	Month       int     `xorm:"not null pk INT(11)" json:"month" form:"month" csv:"month"`
	Registers   int     `xorm:"not null default 0 comment('注册') INT(11)" json:"registers" form:"registers" csv:"registers"`
	Logins      int     `xorm:"not null default 0 comment('登陆') INT(11)" json:"logins" form:"logins" csv:"logins"`
	Usermoney   float32 `xorm:"not null default 0 comment('玩家充值') FLOAT" json:"usermoney" form:"usermoney" csv:"usermoney"`
	Weixinmoney float32 `xorm:"not null default 0 comment('weixin') FLOAT" json:"weixinmoney" form:"weixinmoney" csv:"weixinmoney"`
	Alipaymoney float32 `xorm:"not null default 0 comment('alipay') FLOAT" json:"alipaymoney" form:"alipaymoney" csv:"alipaymoney"`
}
