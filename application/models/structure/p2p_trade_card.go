package structure

type P2pTradeCard struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" json:"id" form:"id" csv:"id"`
	Type       int    `xorm:"not null default 1 comment('1 经销商向系统买房卡  2  经销商向代理卖卡、经销商或代理向玩家卖房卡  3 赠送  4 二代推荐返利 5一代推荐返利 6二代官方充值返利一代现金 7微信公众号买卡8支付宝购卡 20渠道管理员加卡') index INT(10)" json:"type" form:"type" csv:"type"`
	Uid        int    `xorm:"not null default 0 comment('卖家') index INT(10)" json:"uid" form:"uid" csv:"uid"`
	Touid      int    `xorm:"not null default 0 comment('买家') index INT(10)" json:"touid" form:"touid" csv:"touid"`
	Money      string `xorm:"not null default 0 comment('金额') DECIMAL(10)" json:"money" form:"money" csv:"money"`
	CardAmount int    `xorm:"not null default 0 comment('房卡数量') INT(10)" json:"card_amount" form:"card_amount" csv:"card_amount"`
	GiveAmount int    `xorm:"not null default 0 comment('赠送数量') INT(10)" json:"give_amount" form:"give_amount" csv:"give_amount"`
	Timeline   int    `xorm:"not null default 0 comment('时间') index INT(10)" json:"timeline" form:"timeline" csv:"timeline"`
	Admin      int    `xorm:"not null default 0 comment('管理员') INT(11)" json:"admin" form:"admin" csv:"admin"`
	Original   int    `xorm:"not null comment('原房卡') INT(10)" json:"original" form:"original" csv:"original"`
	Remaining  int    `xorm:"not null comment('交易后房卡') INT(10)" json:"remaining" form:"remaining" csv:"remaining"`
	Reason     string `xorm:"not null default '' comment('售卡原因') VARCHAR(255)" json:"reason" form:"reason" csv:"reason"`
	Sellip     string `xorm:"not null default '' comment('ip') VARCHAR(20)" json:"sellip" form:"sellip" csv:"sellip"`
}
