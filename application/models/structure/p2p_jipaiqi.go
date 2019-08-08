package structure

type P2pJipaiqi struct {
	Uid      int   `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gameid   int   `xorm:"not null pk default 0 comment('游戏id') INT(10)" json:"gameid" form:"gameid" csv:"gameid"`
	Golds    int64 `xorm:"not null default 0 comment('累积购买消费金币') BIGINT(20)" json:"golds" form:"golds" csv:"golds"`
	Buytimes int   `xorm:"not null default 0 comment('累积购买次数') INT(8)" json:"buytimes" form:"buytimes" csv:"buytimes"`
	Deadline int   `xorm:"not null default 0 comment('有效期截止日期') index INT(10)" json:"deadline" form:"deadline" csv:"deadline"`
	Ltime    int   `xorm:"not null default 0 comment('开始购买时间') INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
	Uptime   int   `xorm:"not null default 0 comment('最后更新时间') INT(10)" json:"uptime" form:"uptime" csv:"uptime"`
}
