package structure

type UserBoxesRelated struct {
	//个人属性
	Ubid    int    `xorm:"not null pk autoincr comment('用户宝箱id') INT(11)" json:"ubid" form:"ubid" csv:"ubid"`
	Uid     int    `xorm:"not null default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Bid     int    `xorm:"not null default 0 comment('宝箱id') INT(11)" json:"bid" form:"bid" csv:"bid"`
	Optime  int    `xorm:"not null default 0 comment('开启时间') INT(11)" json:"optime" form:"optime" csv:"optime"`
	Bstart  int    `xorm:"not null default 0 comment('计时宝箱是否开始计时 0-未开始 1-已开始') TINYINT(1)" json:"bstart" form:"bstart" csv:"bstart"`
	State   int    `xorm:"not null default 0 comment('状态 0-未开 1-开启 2-临时宝箱 3-即开宝箱') TINYINT(1)" json:"state" form:"state" csv:"state"`
	Optype  int    `xorm:"not null default 0 comment('开启方式 0-到时开启 1-直接宝石开启') TINYINT(1)" json:"optype" form:"optype" csv:"optype"`
	Oprtime int    `xorm:"not null default 0 comment('宝箱真实开启时间') INT(11)" json:"oprtime" form:"oprtime" csv:"oprtime"`
	Opprops string `xorm:"not null default '' comment('获得道具ids，格式 pid1-num1, pid2-num2') VARCHAR(500)" json:"opprops" form:"opprops" csv:"opprops"`
	Ltime   int    `xorm:"not null default 0 comment('宝箱获得时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	//系统属性
	Name    string `xorm:"not null default '0' comment('宝箱名称') VARCHAR(150)" json:"name" form:"name" csv:"name"`
	Btype   int    `xorm:"not null default 0 comment('宝箱类型 1-游戏宝箱 2-抽卡宝箱 3-活动宝箱') TINYINT(3)" json:"btype" form:"btype" csv:"btype"`
	Btime   int    `xorm:"not null default 0 comment('打开时间，单位分钟') INT(8)" json:"btime" form:"btime" csv:"btime"`
	Bnum    int    `xorm:"not null default 0 comment('宝箱开出物品数量') INT(8)" json:"bnum" form:"bnum" csv:"bnum"`
	Quality int    `xorm:"not null default 0 comment('宝箱品质') TINYINT(3)" json:"quality" form:"quality" csv:"quality"`
	Tab     int    `xorm:"not null default 0 comment('标签页') TINYINT(3)" json:"tab" form:"tab" csv:"tab"`
}
