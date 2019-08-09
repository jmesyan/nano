package structure

type GameGuidance struct {
	Step      int    `xorm:"not null pk default 0 comment('引导序号') INT(10)" json:"step" form:"step" csv:"step"`
	Stage     int    `xorm:"default 0 comment('引导阶段  1-登陆欢迎 2-抽卡引导 3-段位赛引导 4-宝箱引导 5-每日任务引导 6-任务角色引导') INT(10)" json:"stage" form:"stage" csv:"stage"`
	Title     string `xorm:"default '' comment('操作步骤') VARCHAR(50)" json:"title" form:"title" csv:"title"`
	Interrupt string `xorm:"default '' comment('中断操作 内容自己约定') VARCHAR(200)" json:"interrupt" form:"interrupt" csv:"interrupt"`
	Tips      string `xorm:"default '' comment('操作提示') VARCHAR(200)" json:"tips" form:"tips" csv:"tips"`
	Guide     string `xorm:"not null default '' comment('引导位置坐标') VARCHAR(100)" json:"guide" form:"guide" csv:"guide"`
	Box       string `xorm:"not null default '' comment('对话框坐标') VARCHAR(100)" json:"box" form:"box" csv:"box"`
}
