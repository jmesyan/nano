package structure

type GameServers struct {
	Hid        int    `xorm:"not null pk autoincr INT(11)" json:"hid" form:"hid" csv:"hid"`
	Title      string `xorm:"not null default '' comment('备注') VARCHAR(45)" json:"title" form:"title" csv:"title"`
	Server     string `xorm:"not null default '' comment('服务器') unique VARCHAR(30)" json:"server" form:"server" csv:"server"`
	Status     int    `xorm:"not null default 0 comment('0 不可用 1 可用') TINYINT(3)" json:"status" form:"status" csv:"status"`
	Admin      int    `xorm:"not null default 0 comment('操作员') INT(10)" json:"admin" form:"admin" csv:"admin"`
	Lasttime   int    `xorm:"not null default 0 comment('修改时间') INT(10)" json:"lasttime" form:"lasttime" csv:"lasttime"`
	Stype      int    `xorm:"not null default 1 comment('服务器类型 1大厅服务器 2PHP接口服务器 3分享服务器') TINYINT(1)" json:"stype" form:"stype" csv:"stype"`
	Uses       int    `xorm:"not null default 0 comment('使用次数') INT(11)" json:"uses" form:"uses" csv:"uses"`
	Usetime    int    `xorm:"not null default 0 comment('使用时间') index INT(11)" json:"usetime" form:"usetime" csv:"usetime"`
	Hallstatus int    `xorm:"not null default 0 comment('大厅状态') TINYINT(3)" json:"hallstatus" form:"hallstatus" csv:"hallstatus"`
}
