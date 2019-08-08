package structure

type GameRecordInfo2051 struct {
	Index       int64  `xorm:"not null pk autoincr comment('索引') BIGINT(30)" json:"index" form:"index" csv:"index"`
	Cards       string `xorm:"comment('手牌') TEXT" json:"cards" form:"cards" csv:"cards"`
	Backcard    string `xorm:"comment('底牌') TEXT" json:"backcard" form:"backcard" csv:"backcard"`
	Points      string `xorm:"comment('分数') TEXT" json:"points" form:"points" csv:"points"`
	Turns       string `xorm:"comment('手数') TEXT" json:"turns" form:"turns" csv:"turns"`
	Bombs       string `xorm:"comment('炸弹') TEXT" json:"bombs" form:"bombs" csv:"bombs"`
	Banker      int    `xorm:"comment('地主玩家') TINYINT(1)" json:"banker" form:"banker" csv:"banker"`
	Doubletimes int    `xorm:"comment('倍数') index INT(10)" json:"doubletimes" form:"doubletimes" csv:"doubletimes"`
	Basescore   int    `xorm:"comment('底分') TINYINT(3)" json:"basescore" form:"basescore" csv:"basescore"`
	Spring      int    `xorm:"comment('1反春 2春天') TINYINT(1)" json:"spring" form:"spring" csv:"spring"`
	Winpos      int    `xorm:"comment('赢家位置') TINYINT(1)" json:"winpos" form:"winpos" csv:"winpos"`
	Gameturns   int    `xorm:"comment('游戏手数') INT(10)" json:"gameturns" form:"gameturns" csv:"gameturns"`
	OrderBy     int    `xorm:"not null default 0 comment('按照倍数排序') index INT(10)" json:"order_by" form:"order_by" csv:"order_by"`
	Cha         int    `xorm:"not null default 0 comment('差值') index INT(10)" json:"cha" form:"cha" csv:"cha"`
}
