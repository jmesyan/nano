package structure

type LogVipLevel struct {
	Uid    int    `xorm:"not null pk default 0 comment('用户id') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Nlevel int    `xorm:"not null pk default 0 comment('当前vip等级') INT(5)" json:"nlevel" form:"nlevel" csv:"nlevel"`
	Olevel int    `xorm:"not null default 0 comment('触发升级时vip等级') INT(5)" json:"olevel" form:"olevel" csv:"olevel"`
	Boxcap int    `xorm:"not null default 0 comment('增加宝箱位数量') INT(5)" json:"boxcap" form:"boxcap" csv:"boxcap"`
	State  int    `xorm:"not null default 0 comment('礼包领取状态 0-未领取 1-已领取') TINYINT(1)" json:"state" form:"state" csv:"state"`
	Wppack string `xorm:"not null default '' comment('道具礼包ids-格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Ltime  int    `xorm:"not null default 0 comment('升级时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime  int    `xorm:"not null default 0 comment('领取礼包时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
