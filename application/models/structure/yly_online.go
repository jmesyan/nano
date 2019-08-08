package structure

type YlyOnline struct {
	Userid     int    `xorm:"not null pk default 0 comment('用户编号') INT(11)" json:"userid" form:"userid" csv:"userid"`
	Gid        int    `xorm:"not null default 0 comment('遊戲id') index(gid_rtype_ridx) INT(10)" json:"gid" form:"gid" csv:"gid"`
	Rtype      int    `xorm:"not null default 0 comment('房間類型') index(gid_rtype_ridx) TINYINT(3)" json:"rtype" form:"rtype" csv:"rtype"`
	Ridx       int    `xorm:"not null default 0 comment('房間id') index(gid_rtype_ridx) TINYINT(3)" json:"ridx" form:"ridx" csv:"ridx"`
	Tid        int    `xorm:"not null default 0 comment('桌子id') INT(11)" json:"tid" form:"tid" csv:"tid"`
	Pos        int    `xorm:"not null default 0 comment('座位號') TINYINT(3)" json:"pos" form:"pos" csv:"pos"`
	LoginTime  int    `xorm:"not null default 0 comment('玩家上线时间') index INT(10)" json:"login_time" form:"login_time" csv:"login_time"`
	UpdateTime int    `xorm:"not null default 0 comment('上次更新时间') index INT(10)" json:"update_time" form:"update_time" csv:"update_time"`
	Fromip     string `xorm:"not null default '0' comment('玩家登陆的ip') VARCHAR(15)" json:"fromip" form:"fromip" csv:"fromip"`
	Ismobile   int    `xorm:"not null default 0 TINYINT(3)" json:"ismobile" form:"ismobile" csv:"ismobile"`
	Device     string `xorm:"not null default 'pc' VARCHAR(100)" json:"device" form:"device" csv:"device"`
	Lang       string `xorm:"not null default 'zh_TW' comment('語言') VARCHAR(32)" json:"lang" form:"lang" csv:"lang"`
	Ismotor    int    `xorm:"not null default 0 comment('机器人') index TINYINT(1)" json:"ismotor" form:"ismotor" csv:"ismotor"`
}
