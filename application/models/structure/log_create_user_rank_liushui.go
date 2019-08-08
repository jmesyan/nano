package structure

type LogCreateUserRankLiushui struct {
	Lid         int   `xorm:"not null pk autoincr index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid         int   `xorm:"not null default 0 index INT(10)" json:"uid" form:"uid" csv:"uid"`
	Gid         int   `xorm:"not null default 0 INT(11)" json:"gid" form:"gid" csv:"gid"`
	Goldschange int64 `xorm:"not null default 0 comment('金币变化') BIGINT(20)" json:"goldschange" form:"goldschange" csv:"goldschange"`
	Change      int64 `xorm:"not null default 0 comment('输赢') BIGINT(20)" json:"change" form:"change" csv:"change"`
	Season      int   `xorm:"not null default 0 comment('赛季') INT(11)" json:"season" form:"season" csv:"season"`
	Stars       int   `xorm:"not null default 0 comment('星级变化') INT(11)" json:"stars" form:"stars" csv:"stars"`
	Score       int   `xorm:"not null default 0 comment('积分') INT(11)" json:"score" form:"score" csv:"score"`
	Ltime       int   `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
