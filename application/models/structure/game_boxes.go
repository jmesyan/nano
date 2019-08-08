package structure

type GameBoxes struct {
	Bid     int    `xorm:"not null pk autoincr comment('宝箱ID') INT(11)" json:"bid" form:"bid" csv:"bid"`
	Name    string `xorm:"not null default '0' comment('宝箱名称') VARCHAR(150)" json:"name" form:"name" csv:"name"`
	Btype   int    `xorm:"not null default 0 comment('宝箱类型 1-游戏宝箱 2-抽卡宝箱 3-活动宝箱') TINYINT(3)" json:"btype" form:"btype" csv:"btype"`
	Btime   int    `xorm:"not null default 0 comment('打开时间，单位分钟') INT(8)" json:"btime" form:"btime" csv:"btime"`
	Bnum    int    `xorm:"not null default 0 comment('宝箱开出物品数量') INT(8)" json:"bnum" form:"bnum" csv:"bnum"`
	Quality int    `xorm:"not null default 0 comment('宝箱品质') TINYINT(3)" json:"quality" form:"quality" csv:"quality"`
	Weights string `xorm:"not null comment('权重配置-json字符串') TEXT" json:"weights" form:"weights" csv:"weights"`
	Frees   int    `xorm:"not null default 0 comment('抽卡系统每天免费次数') TINYINT(3)" json:"frees" form:"frees" csv:"frees"`
	Price   int    `xorm:"not null default 0 comment('抽卡系统抽卡耗钻数') INT(10)" json:"price" form:"price" csv:"price"`
	Tab     int    `xorm:"not null default 0 comment('标签页') TINYINT(3)" json:"tab" form:"tab" csv:"tab"`
	Objects string `xorm:"not null comment('宝箱物品内容') TEXT" json:"objects" form:"objects" csv:"objects"`
}
