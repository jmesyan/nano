package structure

type GameHallNotice struct {
	Id      int    `xorm:"not null pk autoincr INT(11)" json:"id" form:"id" csv:"id"`
	Content string `xorm:"not null comment('内容') TEXT" json:"content" form:"content" csv:"content"`
	State   int    `xorm:"not null default 0 comment('状态 1 显示  0 不显示') TINYINT(3)" json:"state" form:"state" csv:"state"`
	Ltime   int    `xorm:"not null default 0 comment('添加时间') INT(10)" json:"ltime" form:"ltime" csv:"ltime"`
	Admin   int    `xorm:"not null default 0 comment('管理员') INT(10)" json:"admin" form:"admin" csv:"admin"`
	Appid   int    `xorm:"not null default 0 comment('appid') INT(11)" json:"appid" form:"appid" csv:"appid"`
}
