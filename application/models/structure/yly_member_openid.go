package structure

type YlyMemberOpenid struct {
	Appid  int    `xorm:"not null pk comment('appid') INT(11)" json:"appid" form:"appid" csv:"appid"`
	Uid    int    `xorm:"not null pk comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Openid string `xorm:"not null default '' comment('openid') VARCHAR(64)" json:"openid" form:"openid" csv:"openid"`
}
