package structure

import (
	"time"
)

type LogUserGolds struct {
	Lid       int       `xorm:"not null pk autoincr comment('编号') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid       int       `xorm:"not null comment('玩家id') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gameid    int       `xorm:"not null default 0 comment('游戏id') INT(10)" json:"gameid" form:"gameid" csv:"gameid"`
	Roomtype  int       `xorm:"not null default 0 comment('房间类型') TINYINT(3)" json:"roomtype" form:"roomtype" csv:"roomtype"`
	Roomidx   int       `xorm:"not null default 0 comment('房间id') TINYINT(3)" json:"roomidx" form:"roomidx" csv:"roomidx"`
	Tid       int       `xorm:"not null default 0 comment('桌子id') INT(11)" json:"tid" form:"tid" csv:"tid"`
	Ltime     time.Time `xorm:"not null comment('时间') index DATETIME" json:"ltime" form:"ltime" csv:"ltime"`
	Changes   int64     `xorm:"not null default 0 comment('变化量') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
	Remaining int64     `xorm:"not null default 0 comment('结余') BIGINT(20)" json:"remaining" form:"remaining" csv:"remaining"`
	Popcoins  int64     `xorm:"not null default 0 comment('所下筹码数') BIGINT(20)" json:"popcoins" form:"popcoins" csv:"popcoins"`
	State     int       `xorm:"not null default 0 comment('类型') INT(10)" json:"state" form:"state" csv:"state"`
	Remarks   string    `xorm:"comment('备注') TEXT" json:"remarks" form:"remarks" csv:"remarks"`
	Lid2      int       `xorm:"not null default 0 index INT(11)" json:"lid2" form:"lid2" csv:"lid2"`
	Appid     int       `xorm:"not null comment('玩家来源') TINYINT(11)" json:"appid" form:"appid" csv:"appid"`
}
