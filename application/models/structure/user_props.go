package structure

type UserProps struct {
	Upid      int `xorm:"not null pk autoincr comment('用户道具ID') INT(11)" json:"upid" form:"upid" csv:"upid"`
	Uid       int `xorm:"not null comment('用户id') index(uid,pid) INT(11)" json:"uid" form:"uid" csv:"uid"`
	Pid       int `xorm:"not null comment('道具id') index(uid,pid) INT(11)" json:"pid" form:"pid" csv:"pid"`
	Syxhtype  int `xorm:"not null default 0 comment('使用消耗货币类型 1-金币 2-宝石') TINYINT(3)" json:"syxhtype" form:"syxhtype" csv:"syxhtype"`
	Syxhnum   int `xorm:"not null default 0 comment('使用消耗货币数量') INT(5)" json:"syxhnum" form:"syxhnum" csv:"syxhnum"`
	Dqtype    int `xorm:"not null default 0 comment('丢弃类型 0-不可丢弃 1-可丢弃') TINYINT(2)" json:"dqtype" form:"dqtype" csv:"dqtype"`
	Sytype    int `xorm:"not null default 0 comment('使用类型 0-不可使用 1-可使用') TINYINT(2)" json:"sytype" form:"sytype" csv:"sytype"`
	Extime    int `xorm:"not null default 0 comment('过期时间 0永不过期') INT(11)" json:"extime" form:"extime" csv:"extime"`
	Ulroleid  int `xorm:"not null default 0 comment('体验人物ID') INT(11)" json:"ulroleid" form:"ulroleid" csv:"ulroleid"`
	Ulrolelv  int `xorm:"not null default 0 comment('体验人物等级') INT(8)" json:"ulrolelv" form:"ulrolelv" csv:"ulrolelv"`
	Available int `xorm:"not null default 0 comment('可用的数量') index INT(8)" json:"available" form:"available" csv:"available"`
	Used      int `xorm:"not null default 0 comment('已经使用数量') INT(8)" json:"used" form:"used" csv:"used"`
	Abandond  int `xorm:"not null default 0 comment('丢弃的数量') INT(8)" json:"abandond" form:"abandond" csv:"abandond"`
	Expired   int `xorm:"not null default 0 comment('过期的数量') INT(8)" json:"expired" form:"expired" csv:"expired"`
	Ltime     int `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime     int `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
