package structure

type LogCreateUserMatch struct {
	Uid      int   `xorm:"not null pk INT(10)" json:"uid" form:"uid" csv:"uid"`
	Lid      int   `xorm:"not null pk default 0 index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Gid      int   `xorm:"not null default 0 INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype    int   `xorm:"not null default 0 INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	Ridx     int   `xorm:"not null default 0 INT(11)" json:"ridx" form:"ridx" csv:"ridx"`
	Tid      int   `xorm:"not null default 0 INT(11)" json:"tid" form:"tid" csv:"tid"`
	Nroomid  int   `xorm:"not null default 0 comment('房间ID') index INT(11)" json:"nroomid" form:"nroomid" csv:"nroomid"`
	Round    int   `xorm:"not null default 0 comment('总局数') TINYINT(3)" json:"round" form:"round" csv:"round"`
	Cur      int   `xorm:"not null default 0 comment('当前局数') TINYINT(3)" json:"cur" form:"cur" csv:"cur"`
	IsOver   int   `xorm:"not null default 0 comment('是否结束 0未开始 1游戏中 2 结束') TINYINT(1)" json:"is_over" form:"is_over" csv:"is_over"`
	Location int   `xorm:"not null default 0 comment('桌子位置') TINYINT(1)" json:"location" form:"location" csv:"location"`
	Nchange  int64 `xorm:"not null default 0 comment('变化') BIGINT(20)" json:"nchange" form:"nchange" csv:"nchange"`
	Remain   int64 `xorm:"not null default 0 comment('结余') BIGINT(20)" json:"remain" form:"remain" csv:"remain"`
	Origolds int64 `xorm:"not null default 0 comment('原始筹码') BIGINT(20)" json:"origolds" form:"origolds" csv:"origolds"`
	Angang   int   `xorm:"not null default 0 comment('暗杠次数') INT(11)" json:"angang" form:"angang" csv:"angang"`
	Minggang int   `xorm:"not null default 0 comment('明杠次数') INT(11)" json:"minggang" form:"minggang" csv:"minggang"`
	Penggang int   `xorm:"not null default 0 comment('彭刚') INT(11)" json:"penggang" form:"penggang" csv:"penggang"`
	Zimo     int   `xorm:"not null default 0 comment('自摸次数') INT(11)" json:"zimo" form:"zimo" csv:"zimo"`
	Dianpao  int   `xorm:"not null default 0 comment('点炮次数') INT(11)" json:"dianpao" form:"dianpao" csv:"dianpao"`
	Jiepao   int   `xorm:"not null default 0 comment('接炮次数') INT(11)" json:"jiepao" form:"jiepao" csv:"jiepao"`
	Ltime    int   `xorm:"not null default 0 comment('时间') index INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
	Mid      int   `xorm:"not null default 0 INT(11)" json:"mid" form:"mid" csv:"mid"`
	Mlid     int   `xorm:"not null default 0 INT(11)" json:"mlid" form:"mlid" csv:"mlid"`
	Intime   int   `xorm:"not null default 0 comment('进入时间') INT(11)" json:"intime" form:"intime" csv:"intime"`
}
