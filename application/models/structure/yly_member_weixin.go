package structure

type YlyMemberWeixin struct {
	Uid     int    `xorm:"not null pk default 0 comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Unionid string `xorm:"not null default '' comment('微信unionid') index VARCHAR(50)" json:"unionid" form:"unionid" csv:"unionid"`
	Openid  string `xorm:"not null default '' comment('微信openid') VARCHAR(50)" json:"openid" form:"openid" csv:"openid"`
	Invitor int    `xorm:"not null default 0 comment('邀请者') index INT(11)" json:"invitor" form:"invitor" csv:"invitor"`
	Ltime   int    `xorm:"not null default 0 comment('时间') index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Invalid int    `xorm:"not null default 0 comment('0 伯乐有效邀请 1 伯乐无效邀请') index TINYINT(5)" json:"invalid" form:"invalid" csv:"invalid"`
}
