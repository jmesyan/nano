package structure

type LogUserGiftbags struct {
	Lid    int    `xorm:"not null pk autoincr INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid    int    `xorm:"not null default 0 comment('用户ID') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Code   string `xorm:"not null default '' comment('礼包码') VARCHAR(45)" json:"code" form:"code" csv:"code"`
	Wppack string `xorm:"not null default '' comment('礼包内容') VARCHAR(200)" json:"wppack" form:"wppack" csv:"wppack"`
	Ltime  int    `xorm:"not null default 0 comment('领取时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
