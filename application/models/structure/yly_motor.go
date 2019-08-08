package structure

type YlyMotor struct {
	Uid       int   `xorm:"not null pk comment('机器人编号') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Type      int   `xorm:"not null default 0 comment('1 陪玩机器人 2指定机器人 3可编程机器人') TINYINT(3)" json:"type" form:"type" csv:"type"`
	Mingolds  int64 `xorm:"not null default 0 comment('最小携带') BIGINT(11)" json:"mingolds" form:"mingolds" csv:"mingolds"`
	Maxgolds  int64 `xorm:"comment('最大携带') BIGINT(20)" json:"maxgolds" form:"maxgolds" csv:"maxgolds"`
	State     int   `xorm:"not null default 0 comment('状态 1正在使用 0未使用') index TINYINT(3)" json:"state" form:"state" csv:"state"`
	Gameid    int   `xorm:"not null default 0 comment('游戏编号') index(index2) INT(10)" json:"gameid" form:"gameid" csv:"gameid"`
	Roomtype  int   `xorm:"not null default 0 comment('房间类型') index(index2) TINYINT(3)" json:"roomtype" form:"roomtype" csv:"roomtype"`
	Roomidx   int   `xorm:"not null default 0 comment('房间编号') index(index2) TINYINT(3)" json:"roomidx" form:"roomidx" csv:"roomidx"`
	Tid       int   `xorm:"not null default 0 comment('桌子编号') index(index2) INT(11)" json:"tid" form:"tid" csv:"tid"`
	Lastlogin int   `xorm:"not null default 0 comment('最后一次登陆时间') index INT(10)" json:"lastlogin" form:"lastlogin" csv:"lastlogin"`
	Times     int   `xorm:"not null default 0 comment('服务时长') INT(11)" json:"times" form:"times" csv:"times"`
	Rounds    int   `xorm:"not null default 0 comment('每桌服务局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
}
