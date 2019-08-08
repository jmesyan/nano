package structure

type ReportRetention struct {
	Day         int `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Registers   int `xorm:"not null default 0 comment('注册用户数') INT(11)" json:"registers" form:"registers" csv:"registers"`
	Retention1  int `xorm:"not null default 0 comment('1日留存人数') INT(11)" json:"retention1" form:"retention1" csv:"retention1"`
	Retention2  int `xorm:"not null default 0 comment('2日留存人数') INT(11)" json:"retention2" form:"retention2" csv:"retention2"`
	Retention3  int `xorm:"not null default 0 comment('3日留存人数') INT(11)" json:"retention3" form:"retention3" csv:"retention3"`
	Retention4  int `xorm:"not null default 0 comment('4日留存人数') INT(11)" json:"retention4" form:"retention4" csv:"retention4"`
	Retention5  int `xorm:"not null default 0 comment('5日留存人数') INT(11)" json:"retention5" form:"retention5" csv:"retention5"`
	Retention6  int `xorm:"not null default 0 comment('6日留存人数') INT(11)" json:"retention6" form:"retention6" csv:"retention6"`
	Retention7  int `xorm:"not null default 0 comment('7日留存人数') INT(11)" json:"retention7" form:"retention7" csv:"retention7"`
	Retention14 int `xorm:"not null default 0 comment('14日留存人数') INT(11)" json:"retention14" form:"retention14" csv:"retention14"`
	Retention30 int `xorm:"not null default 0 comment('30日留存人数') INT(11)" json:"retention30" form:"retention30" csv:"retention30"`
}
