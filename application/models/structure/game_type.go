package structure

type GameType struct {
	Gsid      int    `xorm:"not null pk comment('游戏编号') INT(10)" json:"gsid" form:"gsid" csv:"gsid"`
	Gstype    int    `xorm:"not null default 0 comment('类型 0 房卡游戏  1 金币游戏') TINYINT(4)" json:"gstype" form:"gstype" csv:"gstype"`
	Gamename  string `xorm:"not null comment('游戏名称') VARCHAR(500)" json:"gamename" form:"gamename" csv:"gamename"`
	Shortkey  string `xorm:"not null comment('短名') CHAR(15)" json:"shortkey" form:"shortkey" csv:"shortkey"`
	Guild     int    `xorm:"not null TINYINT(3)" json:"guild" form:"guild" csv:"guild"`
	Asversion int    `xorm:"not null comment('flash版本') TINYINT(1)" json:"asversion" form:"asversion" csv:"asversion"`
	Remark    string `xorm:"not null default '' comment('备注') VARCHAR(1000)" json:"remark" form:"remark" csv:"remark"`
	State     int    `xorm:"not null default 1 comment('状态 1可用 不可用') TINYINT(1)" json:"state" form:"state" csv:"state"`
	Orderby   int    `xorm:"not null default 0 comment('排序') TINYINT(3)" json:"orderby" form:"orderby" csv:"orderby"`
}
