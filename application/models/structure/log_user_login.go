package structure

type LogUserLogin struct {
	Id     int    `xorm:"not null pk autoincr comment('编号') INT(11)" json:"id" form:"id" csv:"id"`
	Uid    int    `xorm:"not null default 0 comment('用户编号') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Online int    `xorm:"not null default 0 comment('状态 0:上线;1:下线') TINYINT(3)" json:"online" form:"online" csv:"online"`
	Utype  int    `xorm:"not null default 0 comment('登陆类型 0:fb;1:mobile;2web,3admin,4模拟登录') TINYINT(3)" json:"utype" form:"utype" csv:"utype"`
	Ldate  int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ldate" form:"ldate" csv:"ldate"`
	Ip     string `xorm:"not null default '' comment('ip') VARCHAR(15)" json:"ip" form:"ip" csv:"ip"`
}
