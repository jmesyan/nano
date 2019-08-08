package structure

type LogRecurringUser struct {
	Rid      int `xorm:"not null pk autoincr comment('自动ID') INT(11)" json:"rid" form:"rid" csv:"rid"`
	Uid      int `xorm:"not null default 0 comment('uid') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Lasttime int `xorm:"not null default 0 comment('上次登陆时间') INT(11)" json:"lasttime" form:"lasttime" csv:"lasttime"`
	Nowtime  int `xorm:"not null default 0 comment('本次登陆时间') INT(11)" json:"nowtime" form:"nowtime" csv:"nowtime"`
	State    int `xorm:"not null default 0 comment('状态 0未处理 1已处理') TINYINT(4)" json:"state" form:"state" csv:"state"`
	Linvitor int `xorm:"not null default 0 comment('老邀请人') INT(11)" json:"linvitor" form:"linvitor" csv:"linvitor"`
	Ninvitor int `xorm:"not null default 0 comment('新邀请人') INT(11)" json:"ninvitor" form:"ninvitor" csv:"ninvitor"`
	Statebig int `xorm:"not null default 0 comment('0大转盘未读取 1已结读取') TINYINT(4)" json:"statebig" form:"statebig" csv:"statebig"`
}
