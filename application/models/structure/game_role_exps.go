package structure

type GameRoleExps struct {
	Quality   int `xorm:"not null pk default 0 comment('角色品阶') INT(5)" json:"quality" form:"quality" csv:"quality"`
	RoleLevel int `xorm:"not null pk default 0 comment('角色等级') INT(5)" json:"role_level" form:"role_level" csv:"role_level"`
	RoleExps  int `xorm:"not null default 0 comment('累积经验值') INT(10)" json:"role_exps" form:"role_exps" csv:"role_exps"`
	Score     int `xorm:"not null default 0 comment('增加分值倍率') INT(8)" json:"score" form:"score" csv:"score"`
	Exps      int `xorm:"not null default 0 comment('消耗经验值') INT(8)" json:"exps" form:"exps" csv:"exps"`
}
