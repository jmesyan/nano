package structure

type UserWeixinShare struct {
	Uid        int `xorm:"not null pk default 0 comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Invites    int `xorm:"not null default 0 comment('未领取') INT(11)" json:"invites" form:"invites" csv:"invites"`
	Total      int `xorm:"not null default 0 comment('总邀请数') INT(11)" json:"total" form:"total" csv:"total"`
	ValidTotal int `xorm:"not null default 0 comment('邀请的上海用户总数') INT(11)" json:"valid_total" form:"valid_total" csv:"valid_total"`
}
