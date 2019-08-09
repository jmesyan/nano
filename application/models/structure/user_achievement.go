package structure

import (
	"time"
)

type UserAchievement struct {
	Uid                int       `xorm:"not null pk INT(11)" json:"uid" form:"uid" csv:"uid"`
	GameTime           int64     `xorm:"not null default 0 comment('游戏总时长') BIGINT(20)" json:"game_time" form:"game_time" csv:"game_time"`
	Onlinetime         int64     `xorm:"not null default 0 BIGINT(20)" json:"onlinetime" form:"onlinetime" csv:"onlinetime"`
	Nchanges           int64     `xorm:"not null default 0 comment('累计输赢总分数(房卡)') BIGINT(20)" json:"nchanges" form:"nchanges" csv:"nchanges"`
	Rounds             int       `xorm:"not null default 0 comment('累计总次数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Wins               int       `xorm:"not null comment('累计赢次数') INT(11)" json:"wins" form:"wins" csv:"wins"`
	Redpackmoney       string    `xorm:"not null default 0.00 comment('领取红包') DECIMAL(18,2)" json:"redpackmoney" form:"redpackmoney" csv:"redpackmoney"`
	Redpacks           int       `xorm:"not null default 0 comment('领取红包次数') INT(11)" json:"redpacks" form:"redpacks" csv:"redpacks"`
	Freecards          int       `xorm:"not null default 0 comment('免费房卡') INT(11)" json:"freecards" form:"freecards" csv:"freecards"`
	Grounds            int       `xorm:"not null default 0 comment('累计杠的次数') INT(11)" json:"grounds" form:"grounds" csv:"grounds"`
	Agrounds           int       `xorm:"not null default 0 comment('累计暗杠次数') INT(11)" json:"agrounds" form:"agrounds" csv:"agrounds"`
	Lezi               int       `xorm:"not null default 0 comment('累计勒子次数') INT(11)" json:"lezi" form:"lezi" csv:"lezi"`
	Maxwins            int       `xorm:"not null default 0 comment('大赢家次数') INT(11)" json:"maxwins" form:"maxwins" csv:"maxwins"`
	BuyCards           int       `xorm:"not null default 0 comment('累计购买房卡张数') INT(11)" json:"buy_cards" form:"buy_cards" csv:"buy_cards"`
	InuseCards         int       `xorm:"not null default 0 comment('游戏内消耗房卡张数') INT(11)" json:"inuse_cards" form:"inuse_cards" csv:"inuse_cards"`
	ExGifts            int       `xorm:"not null default 0 comment('兑换礼品个数') INT(11)" json:"ex_gifts" form:"ex_gifts" csv:"ex_gifts"`
	RedCards           int       `xorm:"not null default 0 comment('红包生成期间累计游戏内消耗房卡数') INT(11)" json:"red_cards" form:"red_cards" csv:"red_cards"`
	RedRounds          int       `xorm:"not null default 0 comment('红包生成期间累计参与局数') INT(11)" json:"red_rounds" form:"red_rounds" csv:"red_rounds"`
	Tally              int       `xorm:"not null default 0 comment('成就点数') INT(10)" json:"tally" form:"tally" csv:"tally"`
	Atitle             string    `xorm:"not null default '' comment('最新成就头衔') VARCHAR(50)" json:"atitle" form:"atitle" csv:"atitle"`
	RegGolds           int       `xorm:"not null TINYINT(1)" json:"reg_golds" form:"reg_golds" csv:"reg_golds"`
	Goldsliushui       int64     `xorm:"not null default 0 comment('金币累计输赢') BIGINT(20)" json:"goldsliushui" form:"goldsliushui" csv:"goldsliushui"`
	Goldschanges       int64     `xorm:"not null default 0 comment('金币累计赢') BIGINT(20)" json:"goldschanges" form:"goldschanges" csv:"goldschanges"`
	Goldsrounds        int       `xorm:"not null default 0 comment('金币场次') INT(11)" json:"goldsrounds" form:"goldsrounds" csv:"goldsrounds"`
	Goldswins          int       `xorm:"not null default 0 comment('金币赢次数') INT(10)" json:"goldswins" form:"goldswins" csv:"goldswins"`
	Matchrounds        int       `xorm:"not null default 0 comment('比赛场次') INT(11)" json:"matchrounds" form:"matchrounds" csv:"matchrounds"`
	Appid              int       `xorm:"not null default 0 comment('玩家来源') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Shares             int       `xorm:"not null default 0 comment('分享次数') INT(10)" json:"shares" form:"shares" csv:"shares"`
	Unitgolds          int64     `xorm:"not null default 0 comment('累计单位总分数') BIGINT(20)" json:"unitgolds" form:"unitgolds" csv:"unitgolds"`
	Unitwingolds       int64     `xorm:"not null default 0 comment('累计赢单位积分') BIGINT(20)" json:"unitwingolds" form:"unitwingolds" csv:"unitwingolds"`
	Unitrounds         int       `xorm:"not null default 0 comment('累计总局数') INT(10)" json:"unitrounds" form:"unitrounds" csv:"unitrounds"`
	Unitwins           int       `xorm:"not null default 0 comment('累计赢的局数') INT(10)" json:"unitwins" form:"unitwins" csv:"unitwins"`
	Unitloses          int       `xorm:"not null default 0 comment('连输次数') INT(10)" json:"unitloses" form:"unitloses" csv:"unitloses"`
	Totalloses         int       `xorm:"not null default 0 comment('大局连输次数') INT(10)" json:"totalloses" form:"totalloses" csv:"totalloses"`
	Chancerounds       int       `xorm:"not null default 0 comment('奖励做好牌次数') INT(10)" json:"chancerounds" form:"chancerounds" csv:"chancerounds"`
	Goldrounds         int       `xorm:"not null default 0 comment('金币场次') INT(10)" json:"goldrounds" form:"goldrounds" csv:"goldrounds"`
	Nnchanges          int64     `xorm:"not null default 0 BIGINT(20)" json:"nnchanges" form:"nnchanges" csv:"nnchanges"`
	Nnrounds           int64     `xorm:"not null default 0 BIGINT(20)" json:"nnrounds" form:"nnrounds" csv:"nnrounds"`
	Liushui            int64     `xorm:"not null BIGINT(20)" json:"liushui" form:"liushui" csv:"liushui"`
	Gametime           int       `xorm:"default 0 comment('今日游戏时间') INT(11)" json:"gametime" form:"gametime" csv:"gametime"`
	Lastlogintime      time.Time `xorm:"comment('上次登录时间') DATETIME" json:"lastlogintime" form:"lastlogintime" csv:"lastlogintime"`
	Deviateround       int       `xorm:"default 0 comment('实际输赢与计算偏差的局数') INT(11)" json:"deviateround" form:"deviateround" csv:"deviateround"`
	Level              int       `xorm:"not null default 1 comment('等级') INT(11)" json:"level" form:"level" csv:"level"`
	Exp                int       `xorm:"not null default 0 comment('经验') INT(11)" json:"exp" form:"exp" csv:"exp"`
	Preranklianwins    int       `xorm:"not null default 0 comment('中断前连胜次数') INT(11)" json:"preranklianwins" form:"preranklianwins" csv:"preranklianwins"`
	Curranklianwinsbox int       `xorm:"not null default 0 comment('箱子连胜次数') INT(11)" json:"curranklianwinsbox" form:"curranklianwinsbox" csv:"curranklianwinsbox"`
	Curranklianwins    int       `xorm:"not null default 0 comment('当前段位场连赢局数') INT(11)" json:"curranklianwins" form:"curranklianwins" csv:"curranklianwins"`
	Dailyrankwins      int       `xorm:"not null default 0 comment('每日排位胜场') INT(11)" json:"dailyrankwins" form:"dailyrankwins" csv:"dailyrankwins"`
	Paytime            int       `xorm:"not null default 0 comment('充值时间') INT(11)" json:"paytime" form:"paytime" csv:"paytime"`
	Paygolds           int64     `xorm:"not null default 0 comment('充值金额') BIGINT(20)" json:"paygolds" form:"paygolds" csv:"paygolds"`
	Paycount           int       `xorm:"not null default 0 comment('充值次数') INT(11)" json:"paycount" form:"paycount" csv:"paycount"`
	Newpaycount        int       `xorm:"not null default 0 comment('新充值次数') INT(11)" json:"newpaycount" form:"newpaycount" csv:"newpaycount"`
}
