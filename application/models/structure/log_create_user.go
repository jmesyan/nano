package structure

type LogCreateUser struct {
	Uid          int   `xorm:"not null pk INT(10)" json:"uid" form:"uid" csv:"uid"`
	Lid          int   `xorm:"not null pk default 0 index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Gid          int   `xorm:"not null pk default 0 INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype        int   `xorm:"not null default 0 INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	Ridx         int   `xorm:"not null default 0 INT(11)" json:"ridx" form:"ridx" csv:"ridx"`
	Tid          int   `xorm:"not null default 0 INT(11)" json:"tid" form:"tid" csv:"tid"`
	Tablecode    int   `xorm:"not null default 0 comment('房间ID') index INT(11)" json:"tablecode" form:"tablecode" csv:"tablecode"`
	Round        int   `xorm:"not null default 0 comment('总局数') TINYINT(3)" json:"round" form:"round" csv:"round"`
	Cur          int   `xorm:"not null default 0 comment('当前局数') TINYINT(3)" json:"cur" form:"cur" csv:"cur"`
	IsOver       int   `xorm:"not null default 0 comment('是否结束 0未开始 1游戏中 2 结束') TINYINT(1)" json:"is_over" form:"is_over" csv:"is_over"`
	Location     int   `xorm:"not null default 0 comment('桌子位置') TINYINT(1)" json:"location" form:"location" csv:"location"`
	Change       int64 `xorm:"not null default 0 comment('积分变化') BIGINT(20)" json:"change" form:"change" csv:"change"`
	Remain       int64 `xorm:"not null default 0 comment('结余') BIGINT(20)" json:"remain" form:"remain" csv:"remain"`
	Origolds     int64 `xorm:"not null default 0 comment('原始筹码') BIGINT(20)" json:"origolds" form:"origolds" csv:"origolds"`
	Agordz       int   `xorm:"not null default 0 comment('暗杠/地主次数') INT(11)" json:"agordz" form:"agordz" csv:"agordz"`
	Mgorchuned   int   `xorm:"not null default 0 comment('明杠/被春天次数') INT(11)" json:"mgorchuned" form:"mgorchuned" csv:"mgorchuned"`
	Pgoraddtimes int   `xorm:"not null default 0 comment('碰杠/加倍次数') INT(11)" json:"pgoraddtimes" form:"pgoraddtimes" csv:"pgoraddtimes"`
	Zmorgrab     int   `xorm:"not null default 0 comment('自摸/抢地主次数') INT(11)" json:"zmorgrab" form:"zmorgrab" csv:"zmorgrab"`
	Dporbombs    int   `xorm:"not null default 0 comment('点炮/炸弹次数') INT(11)" json:"dporbombs" form:"dporbombs" csv:"dporbombs"`
	Jiepao       int   `xorm:"not null default 0 comment('接炮次数') INT(11)" json:"jiepao" form:"jiepao" csv:"jiepao"`
	Ltime        int   `xorm:"not null default 0 comment('时间') index INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
	Piaohua      int   `xorm:"not null default -1 TINYINT(4)" json:"piaohua" form:"piaohua" csv:"piaohua"`
	Huaer        int   `xorm:"not null default -1 TINYINT(4)" json:"huaer" form:"huaer" csv:"huaer"`
	Jz           int   `xorm:"not null default -1 TINYINT(4)" json:"jz" form:"jz" csv:"jz"`
	Feng         int   `xorm:"not null default -1 TINYINT(4)" json:"feng" form:"feng" csv:"feng"`
	Appid        int   `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
}
