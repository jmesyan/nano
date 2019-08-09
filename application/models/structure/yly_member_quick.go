package structure

type YlyMemberQuick struct {
	Uid         int    `xorm:"not null pk default 0 comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	ChannelCode string `xorm:"not null default '' comment('渠道code') index(channel) VARCHAR(50)" json:"channel_code" form:"channel_code" csv:"channel_code"`
	ChannelUid  string `xorm:"not null default '' comment('渠道uid') index(channel) VARCHAR(50)" json:"channel_uid" form:"channel_uid" csv:"channel_uid"`
	Invitor     int    `xorm:"not null default 0 comment('邀请者') index INT(11)" json:"invitor" form:"invitor" csv:"invitor"`
	Ltime       int    `xorm:"not null default 0 comment('时间') index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
