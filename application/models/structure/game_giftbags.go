package structure

type GameGiftbags struct {
	Code        string `xorm:"not null pk default '0' comment('礼包码') VARCHAR(45)" json:"code" form:"code" csv:"code"`
	Name        string `xorm:"not null default '' comment('礼包名称') VARCHAR(50)" json:"name" form:"name" csv:"name"`
	Wppack      string `xorm:"not null default '' comment('礼包内容') VARCHAR(200)" json:"wppack" form:"wppack" csv:"wppack"`
	Extime      int    `xorm:"not null default 0 comment('过期时间段-分钟') INT(10)" json:"extime" form:"extime" csv:"extime"`
	Expire      int    `xorm:"not null default 0 comment('过期时间') INT(11)" json:"expire" form:"expire" csv:"expire"`
	Channel     int    `xorm:"not null default 0 comment('使用渠道 0-全渠道') INT(5)" json:"channel" form:"channel" csv:"channel"`
	LimitGlobal int    `xorm:"not null default 1 comment('使用次数限制 0-不限制 ') INT(10)" json:"limit_global" form:"limit_global" csv:"limit_global"`
	LimitSingle int    `xorm:"not null default 1 comment('单人使用次数限制0-不限制') INT(10)" json:"limit_single" form:"limit_single" csv:"limit_single"`
	Used        int    `xorm:"not null default 0 comment('已使用次数') INT(10)" json:"used" form:"used" csv:"used"`
	Ltime       int    `xorm:"not null default 0 INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
