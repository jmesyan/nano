package structure

type LogForbid struct {
	Id           int    `xorm:"not null pk autoincr INT(11)" json:"id" form:"id" csv:"id"`
	Uid          int    `xorm:"not null default 0 comment('用户ID') INT(10)" json:"uid" form:"uid" csv:"uid"`
	Forbid       int    `xorm:"not null default 0 comment(' 封号管理员') INT(10)" json:"forbid" form:"forbid" csv:"forbid"`
	ForbidTime   int    `xorm:"not null default 0 comment('封号时间') INT(10)" json:"forbid_time" form:"forbid_time" csv:"forbid_time"`
	Unforbid     int    `xorm:"not null default 0 comment('解封管理员') INT(10)" json:"unforbid" form:"unforbid" csv:"unforbid"`
	UnforbidTime int    `xorm:"not null default 0 comment('解封时间') INT(10)" json:"unforbid_time" form:"unforbid_time" csv:"unforbid_time"`
	Remark       string `xorm:"not null default '' comment('封号备注') VARCHAR(256)" json:"remark" form:"remark" csv:"remark"`
}
