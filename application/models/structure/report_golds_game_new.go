package structure

type ReportGoldsGameNew struct {
	Ldate        int    `xorm:"not null pk comment('日期') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Gid          int    `xorm:"not null pk comment('游戏ID') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype        int    `xorm:"not null pk comment('房间类型') TINYINT(3)" json:"rtype" form:"rtype" csv:"rtype"`
	Rounds       int    `xorm:"not null default 0 comment('参与局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Members      int    `xorm:"not null comment('参与人数') INT(11)" json:"members" form:"members" csv:"members"`
	RobotGolds   int64  `xorm:"not null default 0 comment('机器人产出金币') BIGINT(20)" json:"robot_golds" form:"robot_golds" csv:"robot_golds"`
	Systax       int64  `xorm:"not null comment('金币服务费') BIGINT(20)" json:"systax" form:"systax" csv:"systax"`
	Stonetax     int64  `xorm:"not null default 0 comment('宝石服务费') BIGINT(20)" json:"stonetax" form:"stonetax" csv:"stonetax"`
	Stoneuse     int64  `xorm:"not null default 0 comment('购买消耗宝石') BIGINT(20)" json:"stoneuse" form:"stoneuse" csv:"stoneuse"`
	GiftRedpack  string `xorm:"not null default 0.00 comment('送出红包') DECIMAL(18,2)" json:"gift_redpack" form:"gift_redpack" csv:"gift_redpack"`
	GiftTicket   int64  `xorm:"not null default 0 comment('送出礼券') BIGINT(20)" json:"gift_ticket" form:"gift_ticket" csv:"gift_ticket"`
	Totalmembers int    `xorm:"not null default 0 comment('总参与人数') INT(11)" json:"totalmembers" form:"totalmembers" csv:"totalmembers"`
	Stoneuse2    int64  `xorm:"not null default 0 comment('分游戏统计兑换消耗宝石') BIGINT(20)" json:"stoneuse2" form:"stoneuse2" csv:"stoneuse2"`
}
