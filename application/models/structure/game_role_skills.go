package structure

type GameRoleSkills struct {
	Quality    int `xorm:"not null pk default 0 comment('角色品阶') INT(5)" json:"quality" form:"quality" csv:"quality"`
	SkillLevel int `xorm:"not null pk default 0 comment('角色技能等级') INT(5)" json:"skill_level" form:"skill_level" csv:"skill_level"`
	RoleLevel  int `xorm:"not null default 0 comment('解锁技能需要的角色等级') INT(5)" json:"role_level" form:"role_level" csv:"role_level"`
	Stuff1     int `xorm:"not null default 0 comment('突破材料id1') INT(11)" json:"stuff1" form:"stuff1" csv:"stuff1"`
	Stuff2     int `xorm:"not null default 0 comment('突破材料id2') INT(11)" json:"stuff2" form:"stuff2" csv:"stuff2"`
}
