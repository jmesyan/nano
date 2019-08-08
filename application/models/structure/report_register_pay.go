package structure

type ReportRegisterPay struct {
	Day       int `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Registers int `xorm:"not null default 0 comment('注册用户数') INT(11)" json:"registers" form:"registers" csv:"registers"`
	Pay0      int `xorm:"not null default 0 comment('当日付费人数') INT(11)" json:"pay0" form:"pay0" csv:"pay0"`
	Pay1      int `xorm:"not null default 0 comment('1日付费人数') INT(11)" json:"pay1" form:"pay1" csv:"pay1"`
	Pay2      int `xorm:"not null default 0 comment('2日付费人数') INT(11)" json:"pay2" form:"pay2" csv:"pay2"`
	Pay3      int `xorm:"not null default 0 comment('3日付费人数') INT(11)" json:"pay3" form:"pay3" csv:"pay3"`
	Pay4      int `xorm:"not null default 0 comment('4日付费人数') INT(11)" json:"pay4" form:"pay4" csv:"pay4"`
	Pay5      int `xorm:"not null default 0 comment('5日付费人数') INT(11)" json:"pay5" form:"pay5" csv:"pay5"`
	Pay6      int `xorm:"not null default 0 comment('6日付费人数') INT(11)" json:"pay6" form:"pay6" csv:"pay6"`
	Pay7      int `xorm:"not null default 0 comment('7日付费人数') INT(11)" json:"pay7" form:"pay7" csv:"pay7"`
	Pay14     int `xorm:"not null default 0 comment('14日付费人数') INT(11)" json:"pay14" form:"pay14" csv:"pay14"`
	Pay30     int `xorm:"not null default 0 comment('30日付费人数') INT(11)" json:"pay30" form:"pay30" csv:"pay30"`
}
