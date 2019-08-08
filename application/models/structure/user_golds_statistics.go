package structure

type UserGoldsStatistics struct {
	Date      int   `xorm:"not null pk INT(11)" json:"date" form:"date" csv:"date"`
	Uid       int   `xorm:"not null pk INT(11)" json:"uid" form:"uid" csv:"uid"`
	Wingolds  int64 `xorm:"not null default 0 comment('赢得金币总数') BIGINT(20)" json:"wingolds" form:"wingolds" csv:"wingolds"`
	Losegolds int64 `xorm:"not null default 0 comment('输的金币总数') BIGINT(20)" json:"losegolds" form:"losegolds" csv:"losegolds"`
	Liushui   int64 `xorm:"not null default 0 comment('游戏流水') BIGINT(20)" json:"liushui" form:"liushui" csv:"liushui"`
	Changes   int64 `xorm:"not null default 0 comment('总流水') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
	GameTime  int   `xorm:"not null default 0 comment('在线时长(s)') INT(10)" json:"game_time" form:"game_time" csv:"game_time"`
	Appid     int   `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Rounds    int   `xorm:"not null default 0 comment('参加局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins      int   `xorm:"not null default 0 comment('累计赢次数') INT(11)" json:"wins" form:"wins" csv:"wins"`
	Grounds   int   `xorm:"not null default 0 comment('杠的次数') INT(11)" json:"grounds" form:"grounds" csv:"grounds"`
	Agrounds  int   `xorm:"not null default 0 comment('暗杠的次数') INT(11)" json:"agrounds" form:"agrounds" csv:"agrounds"`
	Lezi      int   `xorm:"not null default 0 comment('勒子') INT(11)" json:"lezi" form:"lezi" csv:"lezi"`
}
