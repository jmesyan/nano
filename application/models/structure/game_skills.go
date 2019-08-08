package structure

type GameSkills struct {
	Skid      int    `xorm:"not null pk autoincr comment('技能ID') INT(11)" json:"skid" form:"skid" csv:"skid"`
	Name      string `xorm:"not null default '' comment('技能名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Icon      int    `xorm:"not null default 0 comment('Icon-id') INT(5)" json:"icon" form:"icon" csv:"icon"`
	Trigger   int    `xorm:"not null default 0 comment('是否重复触发 0-不触发 1-触发') TINYINT(1)" json:"trigger" form:"trigger" csv:"trigger"`
	Condition int    `xorm:"not null default 0 comment('触发条件') TINYINT(3)" json:"condition" form:"condition" csv:"condition"`
	Effect1   int    `xorm:"not null default 0 comment('等级1效果值(百分之)') INT(5)" json:"effect1" form:"effect1" csv:"effect1"`
	Effect2   int    `xorm:"not null default 0 comment('等级2效果值(百分之)') INT(5)" json:"effect2" form:"effect2" csv:"effect2"`
	Effect3   int    `xorm:"not null default 0 comment('等级3效果值(百分之)') INT(5)" json:"effect3" form:"effect3" csv:"effect3"`
	Effect4   int    `xorm:"not null default 0 comment('等级4效果值(百分之)') INT(5)" json:"effect4" form:"effect4" csv:"effect4"`
	Effect5   int    `xorm:"not null default 0 comment('等级5效果值(百分之)') INT(5)" json:"effect5" form:"effect5" csv:"effect5"`
	Effect6   int    `xorm:"not null default 0 comment('等级6效果值(百分之)') INT(5)" json:"effect6" form:"effect6" csv:"effect6"`
	Extra     int    `xorm:"not null default 0 comment('额外影响字段') INT(5)" json:"extra" form:"extra" csv:"extra"`
	Remark    string `xorm:"not null default '' comment('备注') VARCHAR(100)" json:"remark" form:"remark" csv:"remark"`
	Desc      string `xorm:"not null default '' comment('描述') VARCHAR(100)" json:"desc" form:"desc" csv:"desc"`
	Ltime     int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime     int    `xorm:"not null default 0 comment('修改时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
