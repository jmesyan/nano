package structure

type ReportGoldsTotal struct {
	Ldate        int     `xorm:"not null pk comment('日期') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	TotalGolds   int64   `xorm:"not null default 0 comment('金币总数') BIGINT(20)" json:"total_golds" form:"total_golds" csv:"total_golds"`
	GoldsCharge  float32 `xorm:"not null default 0 comment('金币充值金额') FLOAT" json:"golds_charge" form:"golds_charge" csv:"golds_charge"`
	CardGolds    int64   `xorm:"not null default 0 comment('房卡兑换金币') BIGINT(20)" json:"card_golds" form:"card_golds" csv:"card_golds"`
	GiftGolds    int64   `xorm:"not null default 0 comment('低保赠送金币') BIGINT(20)" json:"gift_golds" form:"gift_golds" csv:"gift_golds"`
	RegGolds     int64   `xorm:"not null default 0 comment('注册赠送金币') BIGINT(20)" json:"reg_golds" form:"reg_golds" csv:"reg_golds"`
	RobotGolds   int64   `xorm:"not null default 0 comment('机器人产生金币') BIGINT(20)" json:"robot_golds" form:"robot_golds" csv:"robot_golds"`
	TotalRounds  int     `xorm:"not null default 0 comment('参与局数') INT(11)" json:"total_rounds" form:"total_rounds" csv:"total_rounds"`
	TotalSystax  int64   `xorm:"not null default 0 comment('总服务收取的金币数') BIGINT(20)" json:"total_systax" form:"total_systax" csv:"total_systax"`
	GiftRedpacks float32 `xorm:"not null default 0 comment('送出红包') FLOAT" json:"gift_redpacks" form:"gift_redpacks" csv:"gift_redpacks"`
	GiftTickets  int64   `xorm:"not null default 0 comment('送出礼券') BIGINT(20)" json:"gift_tickets" form:"gift_tickets" csv:"gift_tickets"`
	CardUse      int64   `xorm:"not null default 0 comment('消耗宝石') BIGINT(20)" json:"card_use" form:"card_use" csv:"card_use"`
	GoldsSign    int64   `xorm:"not null default 0 comment('签到赠送金币') BIGINT(20)" json:"golds_sign" form:"golds_sign" csv:"golds_sign"`
	GoldsRank    int64   `xorm:"not null default 0 comment('排行榜奖励金币') BIGINT(20)" json:"golds_rank" form:"golds_rank" csv:"golds_rank"`
	GoldsTask    int64   `xorm:"not null default 0 comment('每日任务奖励金币') BIGINT(20)" json:"golds_task" form:"golds_task" csv:"golds_task"`
	GoldsLucky   int64   `xorm:"not null default 0 comment('好运连连奖励金币') BIGINT(20)" json:"golds_lucky" form:"golds_lucky" csv:"golds_lucky"`
	GoldsMatch   int64   `xorm:"not null default 0 comment('比赛奖励金币') BIGINT(20)" json:"golds_match" form:"golds_match" csv:"golds_match"`
	TotalMembers int     `xorm:"not null default 0 comment('登录用户数') INT(10)" json:"total_members" form:"total_members" csv:"total_members"`
}
