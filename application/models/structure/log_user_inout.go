package structure

import (
	"time"
)

type LogUserInout struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)" json:"id" form:"id" csv:"id"`
	Uid         int       `xorm:"not null default 0 INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gameid      int       `xorm:"not null default 0 INT(11)" json:"gameid" form:"gameid" csv:"gameid"`
	Roomtype    int       `xorm:"not null default 0 INT(10)" json:"roomtype" form:"roomtype" csv:"roomtype"`
	Roomidx     int       `xorm:"not null default 0 INT(10)" json:"roomidx" form:"roomidx" csv:"roomidx"`
	Enterscore  int64     `xorm:"not null default 0 comment('进入时金币') BIGINT(20)" json:"enterscore" form:"enterscore" csv:"enterscore"`
	Entertime   time.Time `xorm:"not null comment('进入时间') DATETIME" json:"entertime" form:"entertime" csv:"entertime"`
	Enteraddr   string    `xorm:"not null default ''000.000.000.000'' comment('进入地址') VARCHAR(32)" json:"enteraddr" form:"enteraddr" csv:"enteraddr"`
	Leavetime   time.Time `xorm:"not null comment('离开时间') DATETIME" json:"leavetime" form:"leavetime" csv:"leavetime"`
	Leavereason int       `xorm:"not null default 0 comment('离开原因') TINYINT(4)" json:"leavereason" form:"leavereason" csv:"leavereason"`
	Leaveaddr   string    `xorm:"not null default ''000.000.000.000'' comment('离开地址') VARCHAR(32)" json:"leaveaddr" form:"leaveaddr" csv:"leaveaddr"`
	Score       int64     `xorm:"not null default 0 comment('金币变化') BIGINT(20)" json:"score" form:"score" csv:"score"`
	Betscore    int64     `xorm:"not null default 0 comment('下注') BIGINT(20)" json:"betscore" form:"betscore" csv:"betscore"`
	Winscore    int64     `xorm:"not null default 0 comment('赢回') BIGINT(20)" json:"winscore" form:"winscore" csv:"winscore"`
	Onlinetime  int       `xorm:"not null default 0 comment('在线时间(秒)') INT(11)" json:"onlinetime" form:"onlinetime" csv:"onlinetime"`
	Experience  int       `xorm:"not null default 0 comment('经验变化') INT(11)" json:"experience" form:"experience" csv:"experience"`
	Maxscore    int64     `xorm:"not null default 0 comment('本次游戏中最高金币') BIGINT(20)" json:"maxscore" form:"maxscore" csv:"maxscore"`
	Minscore    int64     `xorm:"not null default 0 comment('本次游戏中最低金币') BIGINT(20)" json:"minscore" form:"minscore" csv:"minscore"`
	Coupons     int       `xorm:"not null default 0 comment('获得点券') INT(11)" json:"coupons" form:"coupons" csv:"coupons"`
	Appid       int       `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
}
