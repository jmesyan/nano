package structure

type LogOnlineNum struct {
	Day     int `xorm:"not null pk comment('Ymd') INT(10)" json:"day" form:"day" csv:"day"`
	H       int `xorm:"not null pk comment('小时') TINYINT(3)" json:"h" form:"h" csv:"h"`
	M       int `xorm:"not null pk comment('分钟') TINYINT(3)" json:"m" form:"m" csv:"m"`
	Gid     int `xorm:"not null pk comment('游戏') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Rtype   int `xorm:"not null pk default 0 comment('类型') INT(10)" json:"rtype" form:"rtype" csv:"rtype"`
	Onlines int `xorm:"not null default 0 comment('游戏在线人数') MEDIUMINT(8)" json:"onlines" form:"onlines" csv:"onlines"`
	Time    int `xorm:"not null default 0 comment('时间') index INT(11)" json:"time" form:"time" csv:"time"`
	Onm     int `xorm:"not null default 0 INT(11)" json:"onm" form:"onm" csv:"onm"`
	Onf     int `xorm:"not null default 0 INT(11)" json:"onf" form:"onf" csv:"onf"`
}
