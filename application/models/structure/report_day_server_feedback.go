package structure

type ReportDayServerFeedback struct {
	Ldate  int   `xorm:"not null pk default 0 comment('日期') INT(10)" json:"ldate" form:"ldate" csv:"ldate"`
	Users  int   `xorm:"not null default 0 comment('用户数') INT(10)" json:"users" form:"users" csv:"users"`
	Counts int64 `xorm:"not null default 0 comment('次数') BIGINT(20)" json:"counts" form:"counts" csv:"counts"`
}
