package structure

type LogServerFeedback struct {
	Day    int    `xorm:"not null pk comment('天') INT(11)" json:"day" form:"day" csv:"day"`
	Server string `xorm:"not null pk default '' comment('server') VARCHAR(50)" json:"server" form:"server" csv:"server"`
	Uid    int    `xorm:"not null pk comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Ip     string `xorm:"not null default '' comment('IP') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Count  int    `xorm:"not null comment('次数') INT(11)" json:"count" form:"count" csv:"count"`
}
