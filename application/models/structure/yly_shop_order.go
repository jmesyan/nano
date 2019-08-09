package structure

type YlyShopOrder struct {
	Oid         int     `xorm:"not null pk autoincr comment('本地订单id') INT(11)" json:"oid" form:"oid" csv:"oid"`
	Uid         int     `xorm:"not null comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Scid        int     `xorm:"not null comment('商店配置id') INT(10)" json:"scid" form:"scid" csv:"scid"`
	Channel     int     `xorm:"not null default 0 comment('支付渠道 0-所有 1-quick  > 1其他') INT(5)" json:"channel" form:"channel" csv:"channel"`
	Buytype     int     `xorm:"not null default 0 comment('购买方式  1-人民币 2-金币 3-宝石') TINYINT(3)" json:"buytype" form:"buytype" csv:"buytype"`
	Buyprice    float32 `xorm:"not null default 0.00 comment('价格') FLOAT(11,2)" json:"buyprice" form:"buyprice" csv:"buyprice"`
	Discount    int     `xorm:"not null default 0 comment('折扣比%， 0没有') INT(5)" json:"discount" form:"discount" csv:"discount"`
	Money       float32 `xorm:"not null default 0.00 comment('最终购买价格') FLOAT(11,2)" json:"money" form:"money" csv:"money"`
	Wptype      int     `xorm:"not null default 0 comment('物品类型 1-金币 2-宝石 3-道具礼包') TINYINT(3)" json:"wptype" form:"wptype" csv:"wptype"`
	Wpnum       int     `xorm:"not null default 0 comment('物品数量') INT(10)" json:"wpnum" form:"wpnum" csv:"wpnum"`
	Wppack      string  `xorm:"not null default '' comment('道具礼包ids') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Paytype     int     `xorm:"not null default 0 comment('人民币支付方式  1-微信 2-支付宝') TINYINT(3)" json:"paytype" form:"paytype" csv:"paytype"`
	ChannelCode string  `xorm:"not null default '' comment('渠道code') VARCHAR(50)" json:"channel_code" form:"channel_code" csv:"channel_code"`
	ChannelUid  string  `xorm:"not null default '' comment('渠道uid') VARCHAR(50)" json:"channel_uid" form:"channel_uid" csv:"channel_uid"`
	State       int     `xorm:"not null default 0 comment('状态0生成/1+成功') TINYINT(3)" json:"state" form:"state" csv:"state"`
	Ltime       int     `xorm:"not null default 0 comment('配置时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime       int     `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
	IsTest      int     `xorm:"not null default 0 comment('状态 1-测试') TINYINT(1)" json:"is_test" form:"is_test" csv:"is_test"`
}
