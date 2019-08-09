package structure

type RobotGradeScore struct {
	Grade    int `xorm:"not null pk comment('段位') INT(10)" json:"grade" form:"grade" csv:"grade"`
	Minscore int `xorm:"not null default 0 comment('最低积分') INT(11)" json:"minscore" form:"minscore" csv:"minscore"`
	Maxscore int `xorm:"not null comment('最高积分') INT(11)" json:"maxscore" form:"maxscore" csv:"maxscore"`
}
