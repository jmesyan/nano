package structure

type LogUserGoldsStatedesc struct {
	Id       int    `xorm:"not null pk comment('状态') INT(11)" json:"id" form:"id" csv:"id"`
	Desc     string `xorm:"comment('描述') VARCHAR(50)" json:"desc" form:"desc" csv:"desc"`
	OrderNum int    `xorm:"default 0 comment('排序') INT(10)" json:"order_num" form:"order_num" csv:"order_num"`
	IsShow   int    `xorm:"default 1 comment('是否显示') TINYINT(4)" json:"is_show" form:"is_show" csv:"is_show"`
}
