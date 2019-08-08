package structure

type LogMobileLogin struct {
	Lid       int    `xorm:"not null pk autoincr comment('编号') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid       int    `xorm:"not null comment('用户') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ltime     int    `xorm:"not null comment('登陆时间') index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Platform  string `xorm:"default '' comment('平台') VARCHAR(32)" json:"platform" form:"platform" csv:"platform"`
	Device    string `xorm:"default '' comment('设备名称') index(device_ver) VARCHAR(10)" json:"device" form:"device" csv:"device"`
	Deviceuid string `xorm:"comment('设备编号') VARCHAR(128)" json:"deviceuid" form:"deviceuid" csv:"deviceuid"`
	Adid      string `xorm:"VARCHAR(128)" json:"adid" form:"adid" csv:"adid"`
	Ip        string `xorm:"VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Ver       string `xorm:"index(device_ver) VARCHAR(10)" json:"ver" form:"ver" csv:"ver"`
}
