package structure

type UserMatchScore struct {
	Season   int    `xorm:"not null pk default 0 comment('赛季') INT(11)" json:"season" form:"season" csv:"season"`
	Uid      int    `xorm:"not null pk default 0 comment('用户') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Stars    int    `xorm:"not null default 0 comment('星级') INT(11)" json:"stars" form:"stars" csv:"stars"`
	Score    int    `xorm:"not null default 0 comment('积分') INT(11)" json:"score" form:"score" csv:"score"`
	Lasttime int    `xorm:"not null default 0 comment('最后入库时间') INT(11)" json:"lasttime" form:"lasttime" csv:"lasttime"`
	Rounds   int    `xorm:"not null default 0 comment('总局数') INT(10)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins     int    `xorm:"not null default 0 comment('胜利局数') INT(10)" json:"wins" form:"wins" csv:"wins"`
	Rank     int    `xorm:"not null default 0 comment('排名') INT(8)" json:"rank" form:"rank" csv:"rank"`
	Pids     string `xorm:"not null default '' comment('获得道具ids, 格式 pid1-num1, pid2-num2') VARCHAR(300)" json:"pids" form:"pids" csv:"pids"`
	Ppid     int    `xorm:"not null default 0 comment('奖励道具邮件id') INT(11)" json:"ppid" form:"ppid" csv:"ppid"`
	Read     int    `xorm:"not null default 0 comment('前端是否读取') TINYINT(1)" json:"read" form:"read" csv:"read"`
	Astate   int    `xorm:"not null default 0 comment('奖励道具领取状态 0-未领取 1-已领取') TINYINT(1)" json:"astate" form:"astate" csv:"astate"`
	Atime    int    `xorm:"not null default 0 comment('道具奖励领取时间') INT(11)" json:"atime" form:"atime" csv:"atime"`
}
