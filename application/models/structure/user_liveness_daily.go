package structure

type UserLivenessDaily struct {
	Ldate    int `xorm:"not null pk default 0 comment('日期') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Uid      int `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Liveness int `xorm:"not null default 0 comment('活跃度') INT(11)" json:"liveness" form:"liveness" csv:"liveness"`
}
