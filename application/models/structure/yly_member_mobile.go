package structure

type YlyMemberMobile struct {
	Uid         int    `xorm:"not null pk comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Deviceuid   string `xorm:"not null pk default '' comment('deviceuid') index(duidadid) CHAR(128)" json:"deviceuid" form:"deviceuid" csv:"deviceuid"`
	Devicetoken string `xorm:"default '' comment('devicetoken') CHAR(128)" json:"devicetoken" form:"devicetoken" csv:"devicetoken"`
	Platform    string `xorm:"comment('设备平台') VARCHAR(128)" json:"platform" form:"platform" csv:"platform"`
	Macaddress  string `xorm:"comment('mac地址') VARCHAR(128)" json:"macaddress" form:"macaddress" csv:"macaddress"`
	Lastlogin   int    `xorm:"not null comment('最后登陆') INT(10)" json:"lastlogin" form:"lastlogin" csv:"lastlogin"`
	Lastactive  int    `xorm:"not null comment('最后打开应用') INT(10)" json:"lastactive" form:"lastactive" csv:"lastactive"`
	Device      string `xorm:"comment('设备类型') VARCHAR(128)" json:"device" form:"device" csv:"device"`
	Regdate     int    `xorm:"default 0 comment('注册日期') INT(11)" json:"regdate" form:"regdate" csv:"regdate"`
	Gametype    int    `xorm:"default 8 comment('游戏类型') TINYINT(2)" json:"gametype" form:"gametype" csv:"gametype"`
	Regbymobile int    `xorm:"default 1 comment('是否移动注册') index TINYINT(2)" json:"regbymobile" form:"regbymobile" csv:"regbymobile"`
	Imei        string `xorm:"default '' comment('imei') CHAR(32)" json:"imei" form:"imei" csv:"imei"`
	Bundleid    string `xorm:"default '' comment('bundleid') CHAR(128)" json:"bundleid" form:"bundleid" csv:"bundleid"`
	Ver         string `xorm:"default ''1.0'' comment('ver') CHAR(20)" json:"ver" form:"ver" csv:"ver"`
	Staff       string `xorm:"default '' comment('staff') VARCHAR(256)" json:"staff" form:"staff" csv:"staff"`
	Adid        string `xorm:"not null default '' index(duidadid) VARCHAR(128)" json:"adid" form:"adid" csv:"adid"`
}
