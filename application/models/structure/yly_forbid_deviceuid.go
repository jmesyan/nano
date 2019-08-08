package structure

type YlyForbidDeviceuid struct {
	Id        int    `xorm:"not null pk autoincr INT(11)" json:"id" form:"id" csv:"id"`
	Deviceuid string `xorm:"not null default '' comment(' 设备ID') VARCHAR(50)" json:"deviceuid" form:"deviceuid" csv:"deviceuid"`
	Admin     int    `xorm:"not null default 0 comment('管理员') INT(10)" json:"admin" form:"admin" csv:"admin"`
	Remark    string `xorm:"not null default '''' comment('封禁原因') VARCHAR(200)" json:"remark" form:"remark" csv:"remark"`
	Ltime     int    `xorm:"not null default 0 comment('封禁时间') INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
}
