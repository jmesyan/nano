package structure

type LogUserInvite struct {
	Invitor int `xorm:"not null pk default 0 comment('邀请人') INT(11)" json:"invitor" form:"invitor" csv:"invitor"`
	Day     int `xorm:"not null pk default 0 comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Uid     int `xorm:"not null pk default 0 comment('被邀请人') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Cards   int `xorm:"not null default 0 comment('消耗房卡') INT(11)" json:"cards" form:"cards" csv:"cards"`
	Cards2  int `xorm:"not null default 0 comment('消耗房卡') INT(11)" json:"cards2" form:"cards2" csv:"cards2"`
	Imcards int `xorm:"not null default 0 comment('两个月内消耗房卡') INT(11)" json:"imcards" form:"imcards" csv:"imcards"`
	Omcards int `xorm:"not null default 0 comment('两个月外消耗房卡') INT(11)" json:"omcards" form:"omcards" csv:"omcards"`
	Today   int `xorm:"TINYINT(4)" json:"today" form:"today" csv:"today"`
}
