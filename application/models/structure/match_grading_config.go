package structure

type MatchGradingConfig struct {
	Id    int    `xorm:"not null pk default 0 comment('??ID') INT(11)" json:"id" form:"id" csv:"id"`
	Name  string `xorm:"not null default '' comment('????') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Rank  int    `xorm:"not null default 0 comment('??') INT(8)" json:"rank" form:"rank" csv:"rank"`
	Boxid int    `xorm:"not null default 0 comment('??ID') INT(11)" json:"boxid" form:"boxid" csv:"boxid"`
	List  string `xorm:"not null default '' comment('?????') VARCHAR(3000)" json:"list" form:"list" csv:"list"`
}
