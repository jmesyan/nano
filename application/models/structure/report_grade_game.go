package structure

type ReportGradeGame struct {
	Ldate       int   `xorm:"not null pk comment('日期') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Gid         int   `xorm:"not null pk comment('游戏id') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype       int   `xorm:"not null pk comment('房间类型') INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	Rounds      int   `xorm:"not null default 0 comment('参与局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Members     int   `xorm:"not null default 0 comment('参与人数') INT(11)" json:"members" form:"members" csv:"members"`
	Systax      int64 `xorm:"not null default 0 comment('金币服务费') BIGINT(20)" json:"systax" form:"systax" csv:"systax"`
	Changes     int64 `xorm:"not null default 0 comment('机器人输赢') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
	Gametime    int64 `xorm:"not null default 0 comment('游戏时间') BIGINT(20)" json:"gametime" form:"gametime" csv:"gametime"`
	AvgGametime int64 `xorm:"not null default 0 comment('平均游戏时间') BIGINT(20)" json:"avg_gametime" form:"avg_gametime" csv:"avg_gametime"`
}
