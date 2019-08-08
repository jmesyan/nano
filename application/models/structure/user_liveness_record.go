package structure

type UserLivenessRecord struct {
	Livetype int    `xorm:"not null pk default 0 comment('活跃度类型 1-每日活跃度 2-每周活跃度') TINYINT(1)" json:"livetype" form:"livetype" csv:"livetype"`
	Ldate    int    `xorm:"not null pk default 0 comment('日期') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Uid      int    `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Liveness int    `xorm:"not null pk default 0 comment('活跃度') INT(10)" json:"liveness" form:"liveness" csv:"liveness"`
	Wppack   string `xorm:"not null default '' comment('达成目录获得物品，格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Ltime    int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
