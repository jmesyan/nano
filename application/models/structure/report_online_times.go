package structure

type ReportOnlineTimes struct {
	Day    int `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Type   int `xorm:"not null pk default 0 comment('类型 0-5/5-10/10-20/20-30/30-40/40-50/50-60/60-120/120-') TINYINT(1)" json:"type" form:"type" csv:"type"`
	Users  int `xorm:"not null default 0 comment('在线用户数') INT(11)" json:"users" form:"users" csv:"users"`
	Tusers int `xorm:"not null default 0 comment('在线总用户数') INT(11)" json:"tusers" form:"tusers" csv:"tusers"`
}
