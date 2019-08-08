package structure

type GameLiveness struct {
	Livetype int    `xorm:"not null pk default 0 comment('活跃度类型 1-每天 2-每周') INT(10)" json:"livetype" form:"livetype" csv:"livetype"`
	Liveness int    `xorm:"not null pk comment('累积活跃度') INT(10)" json:"liveness" form:"liveness" csv:"liveness"`
	Wppack   string `xorm:"not null default '' comment('奖励礼包 格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
}
