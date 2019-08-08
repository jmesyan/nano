package structure

type MatchSeason struct {
	Sid    int    `xorm:"not null pk comment('赛季id') INT(11)" json:"sid" form:"sid" csv:"sid"`
	Season int    `xorm:"not null pk default 0 comment('赛季') INT(10)" json:"season" form:"season" csv:"season"`
	Name   string `xorm:"not null default '' comment('赛季名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Isrank int    `xorm:"not null default 0 comment('是否排名发奖励邮件') TINYINT(1)" json:"isrank" form:"isrank" csv:"isrank"`
}
