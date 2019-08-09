package structure

type LogUserLiushui2051 struct {
	Ldate      int   `xorm:"not null pk default 0 comment('日期:YYYYMMDD') index(ldate_game_time) index(ldate_uid_changes) INT(8)" json:"ldate" form:"ldate" csv:"ldate"`
	Uid        int   `xorm:"not null pk default 0 comment('玩家id') index(ldate_uid_changes) INT(11)" json:"uid" form:"uid" csv:"uid"`
	GameTime   int   `xorm:"not null default 0 comment('在线时长') index(ldate_game_time) INT(10)" json:"game_time" form:"game_time" csv:"game_time"`
	Onlinetime int64 `xorm:"not null default 0 BIGINT(20)" json:"onlinetime" form:"onlinetime" csv:"onlinetime"`
	Daygolds   int64 `xorm:"not null default 0 comment('玩家当天金币数') BIGINT(20)" json:"daygolds" form:"daygolds" csv:"daygolds"`
	Wingolds   int64 `xorm:"not null default 0 comment('玩家当天累计赢分') BIGINT(20)" json:"wingolds" form:"wingolds" csv:"wingolds"`
	Nchanges   int64 `xorm:"not null default 0 comment('累计输赢总分数') BIGINT(20)" json:"nchanges" form:"nchanges" csv:"nchanges"`
	Rounds     int   `xorm:"not null default 0 comment('累计总次数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins       int   `xorm:"not null default 0 comment('累计赢次数') INT(11)" json:"wins" form:"wins" csv:"wins"`
	Cards      int   `xorm:"not null default 0 INT(11)" json:"cards" form:"cards" csv:"cards"`
	Appid      int   `xorm:"not null default 0 comment('玩家来源') index TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Roomtype   int   `xorm:"not null pk default 0 INT(11)" json:"roomtype" form:"roomtype" csv:"roomtype"`
}
