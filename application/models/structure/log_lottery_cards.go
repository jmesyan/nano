package structure

type LogLotteryCards struct {
	Lid     int    `xorm:"not null pk autoincr comment('流水记录ID') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid     int    `xorm:"not null default 0 comment('用户ID') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Bid     int    `xorm:"not null default 0 comment('宝箱ID') INT(11)" json:"bid" form:"bid" csv:"bid"`
	Bnum    int    `xorm:"not null default 0 comment('宝箱开出物品数量') INT(8)" json:"bnum" form:"bnum" csv:"bnum"`
	Optype  int    `xorm:"not null default 0 comment('抽卡方式 0-赠送 1-耗钻') TINYINT(1)" json:"optype" form:"optype" csv:"optype"`
	Oprtime int    `xorm:"not null default 0 comment('抽卡处理时间') INT(11)" json:"oprtime" form:"oprtime" csv:"oprtime"`
	Opprops string `xorm:"not null default '' comment('获得道具ids，格式 pid1-num1, pid2-num2') VARCHAR(500)" json:"opprops" form:"opprops" csv:"opprops"`
	Price   int    `xorm:"not null default 0 comment('抽卡系统抽卡耗钻数') INT(10)" json:"price" form:"price" csv:"price"`
	State   int    `xorm:"not null default 0 comment('是否处理 0-未处理  1-已处理') TINYINT(1)" json:"state" form:"state" csv:"state"`
	Ltime   int    `xorm:"not null default 0 comment('抽卡时间') index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
