package structure

type LogReconnect struct {
	Lid     int `xorm:"not null pk autoincr INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid     int `xorm:"not null INT(11)" json:"uid" form:"uid" csv:"uid"`
	Type    int `xorm:"not null INT(5)" json:"type" form:"type" csv:"type"`
	Subtype int `xorm:"not null default 0 INT(5)" json:"subtype" form:"subtype" csv:"subtype"`
	Ltime   int `xorm:"not null index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
