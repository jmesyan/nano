package structure

type UserAchievement2052 struct {
	Uid         int    `xorm:"not null pk INT(11)" json:"uid" form:"uid" csv:"uid"`
	GameTime    int64  `xorm:"not null default 0 comment('游戏总时长') BIGINT(20)" json:"game_time" form:"game_time" csv:"game_time"`
	Onlinetime  int64  `xorm:"not null default 0 BIGINT(20)" json:"onlinetime" form:"onlinetime" csv:"onlinetime"`
	Liushui     int64  `xorm:"not null default 0 comment('游戏流水') BIGINT(20)" json:"liushui" form:"liushui" csv:"liushui"`
	Wingolds    int64  `xorm:"not null default 0 comment('赢的总数') BIGINT(20)" json:"wingolds" form:"wingolds" csv:"wingolds"`
	Losegolds   int64  `xorm:"not null default 0 comment('输的总数') BIGINT(20)" json:"losegolds" form:"losegolds" csv:"losegolds"`
	Rounds      int    `xorm:"not null default 0 comment('累计总次数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins        int    `xorm:"not null default 0 comment('累计赢次数') INT(11)" json:"wins" form:"wins" csv:"wins"`
	Frounds     int    `xorm:"not null default 0 comment('进入flop的局数') INT(11)" json:"frounds" form:"frounds" csv:"frounds"`
	Rrounds     int    `xorm:"not null default 0 comment('完整完成游戏的局数') INT(11)" json:"rrounds" form:"rrounds" csv:"rrounds"`
	Maxtype     int    `xorm:"not null default 0 comment('最大牌型') TINYINT(3)" json:"maxtype" form:"maxtype" csv:"maxtype"`
	Maxcs       string `xorm:"not null default ' ' comment('最大牌型的牌') VARCHAR(16)" json:"maxcs" form:"maxcs" csv:"maxcs"`
	Maxwingolds int64  `xorm:"not null default 0 comment('最大赢取金币') BIGINT(20)" json:"maxwingolds" form:"maxwingolds" csv:"maxwingolds"`
	Maxtimes    int    `xorm:"not null default 0 comment('最大倍数(ddz用)') INT(10)" json:"maxtimes" form:"maxtimes" csv:"maxtimes"`
	Maxlianwins int    `xorm:"not null default 0 comment('最大连赢次数') INT(11)" json:"maxlianwins" form:"maxlianwins" csv:"maxlianwins"`
	Curlianwins int    `xorm:"not null default 0 comment('当前连赢次数') INT(11)" json:"curlianwins" form:"curlianwins" csv:"curlianwins"`
}
