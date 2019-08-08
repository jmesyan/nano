package structure

type LogUserProps struct {
	Lid       int `xorm:"not null pk autoincr comment('流水ID') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Uid       int `xorm:"not null default 0 comment('用户ID') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Upid      int `xorm:"not null default 0 comment('用户道具ID') INT(11)" json:"upid" form:"upid" csv:"upid"`
	State     int `xorm:"not null default 0 comment('操作 1-使用 2-丢弃 3-过期 4-获得 5-购买') INT(5)" json:"state" form:"state" csv:"state"`
	Available int `xorm:"not null default 0 comment('可用的数量') INT(8)" json:"available" form:"available" csv:"available"`
	Used      int `xorm:"not null default 0 comment('已经使用数量') INT(8)" json:"used" form:"used" csv:"used"`
	Abandond  int `xorm:"not null default 0 comment('丢弃的数量') INT(8)" json:"abandond" form:"abandond" csv:"abandond"`
	Expired   int `xorm:"not null default 0 comment('过期的数量') INT(8)" json:"expired" form:"expired" csv:"expired"`
	Ltime     int `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
