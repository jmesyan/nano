package structure

import (
	"time"
)

type LogRound2052 struct {
	Lxid     int       `xorm:"not null pk autoincr INT(11)" json:"lxid" form:"lxid" csv:"lxid"`
	Lid      int       `xorm:"not null comment('开桌流水号') index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid      int       `xorm:"not null comment('第一个非机器人的Uid') INT(10)" json:"uid" form:"uid" csv:"uid"`
	Body     string    `xorm:"comment('录像记录') TEXT" json:"body" form:"body" csv:"body"`
	Notetime time.Time `xorm:"comment('记录时间') DATETIME" json:"notetime" form:"notetime" csv:"notetime"`
}
