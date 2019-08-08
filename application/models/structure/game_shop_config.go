package structure

type GameShopConfig struct {
	Scid     int     `xorm:"not null pk autoincr comment('商店配置id') INT(10)" json:"scid" form:"scid" csv:"scid"`
	Name     string  `xorm:"not null default '' comment('配置名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Icon     int     `xorm:"not null default 0 comment('Icon-id') INT(5)" json:"icon" form:"icon" csv:"icon"`
	Channel  int     `xorm:"not null default 0 comment('发行平台 0-所有 >1其他苹果') INT(5)" json:"channel" form:"channel" csv:"channel"`
	Buytype  int     `xorm:"not null default 0 comment('购买方式  1-人民币 2-金币 3-宝石') TINYINT(3)" json:"buytype" form:"buytype" csv:"buytype"`
	Buyprice float32 `xorm:"not null default 0 comment('价格') FLOAT" json:"buyprice" form:"buyprice" csv:"buyprice"`
	Discount int     `xorm:"not null default 0 comment('折扣比%， 0没有') INT(5)" json:"discount" form:"discount" csv:"discount"`
	Paykey   string  `xorm:"not null default '' comment('支付key') VARCHAR(200)" json:"paykey" form:"paykey" csv:"paykey"`
	Wptype   int     `xorm:"not null default 0 comment('物品类型 1-金币 2-宝石 3-道具礼包') TINYINT(3)" json:"wptype" form:"wptype" csv:"wptype"`
	Wpnum    int     `xorm:"not null default 0 comment('物品数量') INT(10)" json:"wpnum" form:"wpnum" csv:"wpnum"`
	Wppack   string  `xorm:"not null default '' comment('道具礼包ids-格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Putaway  int     `xorm:"not null default 0 comment('是否上架 0-不上架 1-上架') TINYINT(1)" json:"putaway" form:"putaway" csv:"putaway"`
	Ltime    int     `xorm:"not null default 0 comment('配置时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime    int     `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
