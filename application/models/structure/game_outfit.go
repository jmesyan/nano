package structure

type GameOutfit struct {
	Uroid  int `xorm:"not null pk autoincr comment('自增id 装备id') INT(11)" json:"uroid" form:"uroid" csv:"uroid"`
	Uid    int `xorm:"not null default 0 comment('用户id') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Roid   int `xorm:"not null default 0 comment('角色id') INT(11)" json:"roid" form:"roid" csv:"roid"`
	Extime int `xorm:"not null default 0 comment('体验道具到期时间，为0的 时候为永久') INT(11)" json:"extime" form:"extime" csv:"extime"`
	Exp    int `xorm:"not null default 0 comment('拥有该装备之后多获得的经验值') INT(11)" json:"exp" form:"exp" csv:"exp"`
	Ltime  int `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
