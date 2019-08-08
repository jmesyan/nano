package structure

type LogCreateUserLiushui101 struct {
	Llid     int   `xorm:"not null pk autoincr INT(11)" json:"llid" form:"llid" csv:"llid"`
	Lid      int   `xorm:"not null default 0 index INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid      int   `xorm:"not null default 0 index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Cur      int   `xorm:"not null default 0 TINYINT(3)" json:"cur" form:"cur" csv:"cur"`
	Ncode    int   `xorm:"not null default 0 comment('房号') INT(11)" json:"ncode" form:"ncode" csv:"ncode"`
	Gid      int   `xorm:"not null default 0 INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype    int   `xorm:"not null default 0 INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	Ridx     int   `xorm:"not null default 0 INT(11)" json:"ridx" form:"ridx" csv:"ridx"`
	Tid      int   `xorm:"not null default 0 INT(11)" json:"tid" form:"tid" csv:"tid"`
	Nchange  int64 `xorm:"not null default 0 comment('变化') BIGINT(20)" json:"nchange" form:"nchange" csv:"nchange"`
	Origolds int64 `xorm:"not null default 0 comment('原始筹码') BIGINT(20)" json:"origolds" form:"origolds" csv:"origolds"`
	Angang   int   `xorm:"default 0 comment('暗杠') INT(11)" json:"angang" form:"angang" csv:"angang"`
	Minggang int   `xorm:"not null default 0 comment('明杠') INT(11)" json:"minggang" form:"minggang" csv:"minggang"`
	Penggang int   `xorm:"not null default 0 comment('碰杠') INT(11)" json:"penggang" form:"penggang" csv:"penggang"`
	Type     int   `xorm:"not null default 0 comment('0表示与自己无关，1表示输家，2接赢家') INT(11)" json:"type" form:"type" csv:"type"`
	Ntime    int   `xorm:"not null default 0 index INT(11)" json:"ntime" form:"ntime" csv:"ntime"`
}
