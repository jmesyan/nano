package structure

import "time"

type LogCreateUserLiushuiRelated struct {
	Lid      int       `xorm:"not null pk autoincr index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid1     int       `xorm:"not null default 0 index INT(10)" json:"uid1" form:"uid1" csv:"uid1"`
	Pregolds int64     `xorm:"not null default 0 comment('原始筹码') BIGINT(20)" json:"pregolds" form:"pregolds" csv:"pregolds"`
	Uid2     int       `xorm:"not null default 0 INT(10)" json:"uid2" form:"uid2" csv:"uid2"`
	Uid3     int       `xorm:"not null default 0 INT(10)" json:"uid3" form:"uid3" csv:"uid3"`
	Uid4     int       `xorm:"not null default 0 INT(10)" json:"uid4" form:"uid4" csv:"uid4"`
	Roomtype int       `xorm:"not null default 0 comment('房间类型') INT(11)" json:"roomtype" form:"roomtype" csv:"roomtype"`
	Notetime time.Time `xorm:"comment('记录时间') DATETIME" json:"notetime" form:"notetime" csv:"notetime"`
	Systax   int64     `xorm:"default 0 comment('税收') BIGINT(20)" json:"systax" form:"systax" csv:"systax"`
	Change   int64     `xorm:"not null default 0 comment('输赢') BIGINT(20)" json:"change" form:"change" csv:"change"`
	Wins     int64     `xorm:"not null default 0 comment('总输赢') BIGINT(20)" json:"wins" form:"wins" csv:"wins"`
	Fan      int       `xorm:"not null default 0 comment('番数') INT(11)" json:"fan" form:"fan" csv:"fan"`
	Task     int       `xorm:"not null default 0 comment('任务倍数') INT(11)" json:"task" form:"task" csv:"task"`
	Lxid     int       `xorm:"not null default 0 comment('录像lid') INT(11)" json:"lxid" form:"lxid" csv:"lxid"`
	Appid    int       `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
}
