package structure

type GameGoldsType struct {
	Gsid      int    `xorm:"not null pk comment('游戏编号') INT(10)" json:"gsid" form:"gsid" csv:"gsid"`
	Gamename  string `xorm:"not null comment('游戏名称') VARCHAR(500)" json:"gamename" form:"gamename" csv:"gamename"`
	Shortkey  string `xorm:"not null default '' comment('英文短名') VARCHAR(20)" json:"shortkey" form:"shortkey" csv:"shortkey"`
	Gstype    int    `xorm:"not null default 0 comment('类型 1金币游戏 0房卡游戏 0代表2者都有') TINYINT(1)" json:"gstype" form:"gstype" csv:"gstype"`
	Gsclass   string `xorm:"not null comment('集合') VARCHAR(200)" json:"gsclass" form:"gsclass" csv:"gsclass"`
	Remark    string `xorm:"not null default '' comment('简单描述') VARCHAR(1000)" json:"remark" form:"remark" csv:"remark"`
	State     int    `xorm:"not null default 1 comment('状态 1可用 0不可用') TINYINT(1)" json:"state" form:"state" csv:"state"`
	Orderby   int    `xorm:"not null default 0 comment('排序') INT(11)" json:"orderby" form:"orderby" csv:"orderby"`
	Apiserver string `xorm:"not null default '' comment('游戏服务器地址') VARCHAR(500)" json:"apiserver" form:"apiserver" csv:"apiserver"`
	IsHot     int    `xorm:"not null default 0 comment('牛人在玩啥') TINYINT(1)" json:"is_hot" form:"is_hot" csv:"is_hot"`
	UiRootid  int    `xorm:"not null default 0 comment('UI区域编号') INT(1)" json:"ui_rootid" form:"ui_rootid" csv:"ui_rootid"`
	UiPosid   int    `xorm:"not null default 0 comment('UI位置编号') INT(1)" json:"ui_posid" form:"ui_posid" csv:"ui_posid"`
	UiColor   int    `xorm:"not null default 0 comment('UI颜色') INT(8)" json:"ui_color" form:"ui_color" csv:"ui_color"`
	Censerver string `xorm:"not null default '' comment('中心服务器地址') VARCHAR(100)" json:"censerver" form:"censerver" csv:"censerver"`
}
