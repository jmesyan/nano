package structure

type LogUserLiushui struct {
	Ldate        int     `xorm:"not null pk default 0 comment('日期:YYYYMMDD') index INT(8)" json:"ldate" form:"ldate" csv:"ldate"`
	Uid          int     `xorm:"not null pk default 0 comment('玩家id') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	GameTime     int     `xorm:"not null default 0 comment('在线时长') index INT(10)" json:"game_time" form:"game_time" csv:"game_time"`
	Onlinetime   int64   `xorm:"not null default 0 index BIGINT(20)" json:"onlinetime" form:"onlinetime" csv:"onlinetime"`
	Nchanges     int64   `xorm:"not null default 0 comment('累计输赢总分数（房卡）') BIGINT(20)" json:"nchanges" form:"nchanges" csv:"nchanges"`
	Rounds       int     `xorm:"not null default 0 comment('累计总次数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins         int     `xorm:"not null default 0 comment('累计赢次数') INT(11)" json:"wins" form:"wins" csv:"wins"`
	Goldsliushui int64   `xorm:"not null default 0 comment('金币累计输赢') BIGINT(20)" json:"goldsliushui" form:"goldsliushui" csv:"goldsliushui"`
	Goldschanges int64   `xorm:"not null default 0 comment('金币累计赢') BIGINT(20)" json:"goldschanges" form:"goldschanges" csv:"goldschanges"`
	Goldsrounds  int     `xorm:"not null default 0 comment('金币场次（修改+s）') INT(11)" json:"goldsrounds" form:"goldsrounds" csv:"goldsrounds"`
	Goldswins    int     `xorm:"not null default 0 comment('金币累计赢') INT(11)" json:"goldswins" form:"goldswins" csv:"goldswins"`
	Cards        int     `xorm:"not null default 0 INT(11)" json:"cards" form:"cards" csv:"cards"`
	Cardsall     int     `xorm:"not null comment('每日总耗卡') INT(11)" json:"cardsAll" form:"cardsAll" csv:"cardsAll"`
	Grounds      int     `xorm:"not null default 0 comment('杠的次数') INT(11)" json:"grounds" form:"grounds" csv:"grounds"`
	Agrounds     int     `xorm:"not null default 0 comment('暗杠的次数') INT(11)" json:"agrounds" form:"agrounds" csv:"agrounds"`
	Lezi         int     `xorm:"not null default 0 comment('勒子') INT(11)" json:"lezi" form:"lezi" csv:"lezi"`
	Maxwins      int     `xorm:"not null default 0 comment('大赢家次数') INT(11)" json:"maxwins" form:"maxwins" csv:"maxwins"`
	ExGifts      int     `xorm:"not null default 0 comment('兑换礼品个数') INT(11)" json:"ex_gifts" form:"ex_gifts" csv:"ex_gifts"`
	BuyMascots   int     `xorm:"not null default 0 comment('购买吉祥物个数') INT(11)" json:"buy_mascots" form:"buy_mascots" csv:"buy_mascots"`
	Matchrounds  int     `xorm:"not null default 00000000000 comment('比赛场次') INT(11)" json:"matchrounds" form:"matchrounds" csv:"matchrounds"`
	Appid        int     `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	AlmsTimes    int     `xorm:"not null default 0 comment('救济金领取次数') TINYINT(3)" json:"alms_times" form:"alms_times" csv:"alms_times"`
	Shares       int     `xorm:"not null default 0 comment('分享次数') INT(10)" json:"shares" form:"shares" csv:"shares"`
	Points       int     `xorm:"not null default 0 comment('今天兑换的PP值') INT(10)" json:"points" form:"points" csv:"points"`
	GoldsCards   int     `xorm:"not null default 0 comment('今天有没有领取金币场房卡') INT(10)" json:"golds_cards" form:"golds_cards" csv:"golds_cards"`
	Gexcards     int     `xorm:"not null default 0 comment('兑换金币消耗房卡') INT(10)" json:"gexcards" form:"gexcards" csv:"gexcards"`
	Takeparts    int     `xorm:"not null default 0 comment('有效参与次数') INT(10)" json:"takeparts" form:"takeparts" csv:"takeparts"`
	Recharges    float32 `xorm:"not null default 0 comment('官方充值金额') FLOAT" json:"recharges" form:"recharges" csv:"recharges"`
	Roomcards    int     `xorm:"not null default 0 comment('开房房卡消耗') INT(11)" json:"roomcards" form:"roomcards" csv:"roomcards"`
	Roomrounds   int     `xorm:"not null default 0 comment('开房耗卡参与局数') INT(11)" json:"roomrounds" form:"roomrounds" csv:"roomrounds"`
	Goldchange   int64   `xorm:"not null default 0 comment('金币场外金币变化量') BIGINT(20)" json:"goldchange" form:"goldchange" csv:"goldchange"`
	Unitgolds    int64   `xorm:"not null default 0 comment('累计输赢单位分数') BIGINT(20)" json:"unitgolds" form:"unitgolds" csv:"unitgolds"`
	Unitwingolds int64   `xorm:"not null default 0 comment('累计赢的单位分数') BIGINT(20)" json:"unitwingolds" form:"unitwingolds" csv:"unitwingolds"`
	Unitrounds   int     `xorm:"not null default 0 comment('累计总局数') INT(10)" json:"unitrounds" form:"unitrounds" csv:"unitrounds"`
	Unitwins     int     `xorm:"not null default 0 comment('累计赢的局数') INT(10)" json:"unitwins" form:"unitwins" csv:"unitwins"`
	Goldrounds   int     `xorm:"not null default 0 comment('金币场次') INT(10)" json:"goldrounds" form:"goldrounds" csv:"goldrounds"`
	Ngoldchanges int64   `xorm:"not null default 0 BIGINT(20)" json:"ngoldchanges" form:"ngoldchanges" csv:"ngoldchanges"`
	Nwingolds    int64   `xorm:"not null default 0 BIGINT(20)" json:"nwingolds" form:"nwingolds" csv:"nwingolds"`
}
