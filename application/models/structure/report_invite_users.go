package structure

type ReportInviteUsers struct {
	Month       int `xorm:"not null pk default 0 comment('月份') INT(11)" json:"month" form:"month" csv:"month"`
	Uid         int `xorm:"not null pk default 0 comment('邀请人') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Invites     int `xorm:"not null default 0 comment('邀请人数') INT(11)" json:"invites" form:"invites" csv:"invites"`
	Invitecards int `xorm:"not null default 0 comment('邀请消耗房卡') INT(11)" json:"invitecards" form:"invitecards" csv:"invitecards"`
	Rewardcards int `xorm:"not null default 0 comment('奖励房卡数') INT(11)" json:"rewardcards" form:"rewardcards" csv:"rewardcards"`
	State       int `xorm:"not null default 0 comment('状态：0未发放 1已发放') TINYINT(4)" json:"state" form:"state" csv:"state"`
	Remaining   int `xorm:"not null default 0 comment('剩余房卡') INT(11)" json:"remaining" form:"remaining" csv:"remaining"`
	Invites2    int `xorm:"not null default 0 INT(11)" json:"invites2" form:"invites2" csv:"invites2"`
	Imcards     int `xorm:"not null default 0 comment('两个月内消耗房卡') INT(11)" json:"imcards" form:"imcards" csv:"imcards"`
	Omcards     int `xorm:"not null default 0 comment('两个月外消耗房卡') INT(11)" json:"omcards" form:"omcards" csv:"omcards"`
}
