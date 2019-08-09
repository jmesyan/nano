package structure

type GameGradeAddition struct {
	Season   int    `xorm:"not null pk default 0 comment('赛季') INT(11)" json:"season" form:"season" csv:"season"`
	Gid      int    `xorm:"not null pk default 0 comment('游戏id') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Addition string `xorm:"not null default '' comment('加成比例') VARCHAR(300)" json:"addition" form:"addition" csv:"addition"`
}
