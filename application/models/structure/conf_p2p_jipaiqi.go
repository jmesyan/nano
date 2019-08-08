package structure

type ConfP2pJipaiqi struct {
	Gameid int   `xorm:"not null pk default 0 comment('游戏id') INT(10)" json:"gameid" form:"gameid" csv:"gameid"`
	Days   int   `xorm:"not null pk comment('有效天数') INT(8)" json:"days" form:"days" csv:"days"`
	Golds  int64 `xorm:"not null default 0 comment('购买需要金币数量') BIGINT(20)" json:"golds" form:"golds" csv:"golds"`
	Ltime  int   `xorm:"not null default 0 comment('时间') INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
}
