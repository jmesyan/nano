package structure

type UserGameLiushui struct {
	Ldate            int     `xorm:"not null pk default 0 comment('日期 格式-20190101') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Uid              int     `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gid              int     `xorm:"not null pk default 0 comment('游戏id 0-全部 1000-段位  >1000 各游戏') INT(11)" json:"gid" form:"gid" csv:"gid"`
	GameWins         int     `xorm:"not null default 0 comment('累积获胜次数') INT(11)" json:"game_wins" form:"game_wins" csv:"game_wins"`
	GameRounds       int     `xorm:"not null default 0 comment('累积游戏局数') INT(11)" json:"game_rounds" form:"game_rounds" csv:"game_rounds"`
	GameWinrounds    int     `xorm:"not null default 0 comment('连赢多少局') INT(10)" json:"game_winrounds" form:"game_winrounds" csv:"game_winrounds"`
	GameSeconds      int     `xorm:"not null default 0 comment('累积游戏时长（秒）') INT(11)" json:"game_seconds" form:"game_seconds" csv:"game_seconds"`
	GameBigwincount  int     `xorm:"not null default 0 comment('单局赢X金币以上局数') INT(10)" json:"game_bigwincount" form:"game_bigwincount" csv:"game_bigwincount"`
	GameWingolds     int64   `xorm:"not null default 0 comment('累积赢多少金币') BIGINT(20)" json:"game_wingolds" form:"game_wingolds" csv:"game_wingolds"`
	GameBrokes       int     `xorm:"not null default 0 comment('累积致对手造成破产次数') INT(10)" json:"game_brokes" form:"game_brokes" csv:"game_brokes"`
	GameBombs        int     `xorm:"not null default 0 comment('累积打出炸弹次数') INT(10)" json:"game_bombs" form:"game_bombs" csv:"game_bombs"`
	MaxBombs         int     `xorm:"not null default 0 comment('单局打出炸弹次数') INT(10)" json:"max_bombs" form:"max_bombs" csv:"max_bombs"`
	GameWingscore    int     `xorm:"not null default 0 comment('累积赢得段位分') INT(11)" json:"game_wingscore" form:"game_wingscore" csv:"game_wingscore"`
	GameUpgstars     int     `xorm:"not null default 0 comment('累积升星') INT(11)" json:"game_upgstars" form:"game_upgstars" csv:"game_upgstars"`
	GameUpgrades     int     `xorm:"not null default 0 comment('累积升段位') INT(11)" json:"game_upgrades" form:"game_upgrades" csv:"game_upgrades"`
	GameWstreakboxes int     `xorm:"not null default 0 comment('获取连胜宝箱次数') INT(11)" json:"game_wstreakboxes" form:"game_wstreakboxes" csv:"game_wstreakboxes"`
	OpenWstreakboxes int     `xorm:"not null default 0 comment('开启连胜宝箱次数') INT(11)" json:"open_wstreakboxes" form:"open_wstreakboxes" csv:"open_wstreakboxes"`
	GameLawtimes     int     `xorm:"not null default 0 comment('地主身份胜利X次') INT(10)" json:"game_lawtimes" form:"game_lawtimes" csv:"game_lawtimes"`
	GameFawtimes     int     `xorm:"not null default 0 comment('农民身份胜利X次') INT(10)" json:"game_fawtimes" form:"game_fawtimes" csv:"game_fawtimes"`
	GameReliefs      int     `xorm:"not null default 0 comment('累积救济次数') INT(10)" json:"game_reliefs" form:"game_reliefs" csv:"game_reliefs"`
	GameRecharges    float32 `xorm:"not null default 0 comment('累积充值金额') FLOAT" json:"game_recharges" form:"game_recharges" csv:"game_recharges"`
	MaxRecharges     float32 `xorm:"not null default 0 comment('单笔充值最高金额') FLOAT" json:"max_recharges" form:"max_recharges" csv:"max_recharges"`
	GameMcards       int     `xorm:"not null default 0 comment('普通月卡领奖次数') INT(10)" json:"game_mcards" form:"game_mcards" csv:"game_mcards"`
	GameShares       int     `xorm:"not null default 0 comment('游戏分享次数') INT(10)" json:"game_shares" form:"game_shares" csv:"game_shares"`
	GameTelbind      int     `xorm:"not null default 0 comment('完成手机绑定') INT(10)" json:"game_telbind" form:"game_telbind" csv:"game_telbind"`
	GameRealname     int     `xorm:"not null default 0 comment('完成实名认证') INT(10)" json:"game_realname" form:"game_realname" csv:"game_realname"`
	GameRolegifts    int     `xorm:"not null default 0 comment('给人物送礼次数') INT(10)" json:"game_rolegifts" form:"game_rolegifts" csv:"game_rolegifts"`
	GameRoleups      int     `xorm:"not null default 0 comment('人物升级次数') INT(10)" json:"game_roleups" form:"game_roleups" csv:"game_roleups"`
	GameRolebreaks   int     `xorm:"not null default 0 comment('人物突破次数') INT(10)" json:"game_rolebreaks" form:"game_rolebreaks" csv:"game_rolebreaks"`
	GameRolegains    int     `xorm:"not null default 0 comment('获得人物次数') INT(10)" json:"game_rolegains" form:"game_rolegains" csv:"game_rolegains"`
	GameRolemaxs     int     `xorm:"not null default 0 comment('满级人物次数') INT(10)" json:"game_rolemaxs" form:"game_rolemaxs" csv:"game_rolemaxs"`
	GameLogins       int     `xorm:"not null default 0 comment('游戏登陆天数') INT(10)" json:"game_logins" form:"game_logins" csv:"game_logins"`
	MaxLogins        int     `xorm:"not null default 0 comment('连续登陆游戏天数') INT(10)" json:"max_logins" form:"max_logins" csv:"max_logins"`
}
