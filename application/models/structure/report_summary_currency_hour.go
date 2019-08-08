package structure

type ReportSummaryCurrencyHour struct {
	Day      int    `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	H        int    `xorm:"not null pk default 0 TINYINT(3)" json:"h" form:"h" csv:"h"`
	Paytype  int    `xorm:"not null pk default 0 comment('人民币支付方式  1-微信 2-支付宝') TINYINT(3)" json:"paytype" form:"paytype" csv:"paytype"`
	PayGolds int64  `xorm:"not null default 0 comment('充值金币') BIGINT(20)" json:"pay_golds" form:"pay_golds" csv:"pay_golds"`
	PayMoney string `xorm:"not null default 0.00 comment('充值金额') DECIMAL(18,2)" json:"pay_money" form:"pay_money" csv:"pay_money"`
	PayUsers int    `xorm:"not null default 0 comment('充值人数') INT(11)" json:"pay_users" form:"pay_users" csv:"pay_users"`
	Pays     int    `xorm:"not null default 0 comment('充值次数') INT(11)" json:"pays" form:"pays" csv:"pays"`
}
