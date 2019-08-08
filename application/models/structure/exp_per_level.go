package structure

type ExpPerLevel struct {
	Level int `xorm:"not null default 0 comment('等级') INT(11)" json:"level" form:"level" csv:"level"`
	Exp   int `xorm:"not null default 0 comment('所需经验') INT(11)" json:"exp" form:"exp" csv:"exp"`
}
