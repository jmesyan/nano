package structure

type LogGamePlays struct {
	Date  int `xorm:"not null pk default 0 comment('天') INT(10)" json:"date" form:"date" csv:"date"`
	Appid int `xorm:"not null pk default 0 comment('包') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Uid   int `xorm:"not null pk default 0 comment('用户id') INT(10)" json:"uid" form:"uid" csv:"uid"`
	Gid   int `xorm:"not null pk default 0 comment('游戏id') SMALLINT(5)" json:"gid" form:"gid" csv:"gid"`
	Num   int `xorm:"not null default 1 comment('打的次数') INT(10)" json:"num" form:"num" csv:"num"`
}
