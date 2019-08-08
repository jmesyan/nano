package structure

type GameRankStat struct {
	Stars          int `xorm:"not null pk default 0 comment('段位等级') TINYINT(4)" json:"stars" form:"stars" csv:"stars"`
	Rank           int `xorm:"default 0 comment('段位分') INT(10)" json:"rank" form:"rank" csv:"rank"`
	Score          int `xorm:"default 0 comment('基础积分') INT(10)" json:"score" form:"score" csv:"score"`
	Coin           int `xorm:"default 0 comment('入场需求') INT(10)" json:"coin" form:"coin" csv:"coin"`
	Tax            int `xorm:"default 0 comment('门票') INT(10)" json:"tax" form:"tax" csv:"tax"`
	Base           int `xorm:"default 0 comment('底分') INT(10)" json:"base" form:"base" csv:"base"`
	Bid            int `xorm:"default 0 comment('掉落宝箱id') INT(10)" json:"bid" form:"bid" csv:"bid"`
	Exp            int `xorm:"default 0 comment('增加经验值') INT(10)" json:"exp" form:"exp" csv:"exp"`
	Roomgoldfactor int `xorm:"not null default 0 comment('房间等级金币权重') INT(10)" json:"roomgoldfactor" form:"roomgoldfactor" csv:"roomgoldfactor"`
	Roomtimefactor int `xorm:"not null default 0 comment('房间等级时间权重(千分之)') INT(10)" json:"roomtimefactor" form:"roomtimefactor" csv:"roomtimefactor"`
}
