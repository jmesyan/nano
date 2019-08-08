package structure

type GameWinStreak struct {
	Wins    int `xorm:"not null default 0 comment('连胜次数') INT(8)" json:"wins" form:"wins" csv:"wins"`
	Weights int `xorm:"not null default 0 comment('加分权重 单位%') INT(10)" json:"weights" form:"weights" csv:"weights"`
}
