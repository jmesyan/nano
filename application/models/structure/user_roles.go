package structure

type UserRoles struct {
	Uroid       int    `xorm:"not null pk autoincr comment('用户角色id') INT(11)" json:"uroid" form:"uroid" csv:"uroid"`
	Uid         int    `xorm:"not null default 0 comment('用户id') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Roid        int    `xorm:"not null default 0 comment('角色id') INT(11)" json:"roid" form:"roid" csv:"roid"`
	Extime      int    `xorm:"not null default 0 comment('体验道具到期时间，为0的 时候为永久') INT(11)" json:"extime" form:"extime" csv:"extime"`
	RoleLevel   int    `xorm:"not null default 0 comment('角色等级') INT(5)" json:"role_level" form:"role_level" csv:"role_level"`
	RoleExps    int    `xorm:"not null default 0 comment('角色经验值') INT(11)" json:"role_exps" form:"role_exps" csv:"role_exps"`
	SkillLevel  int    `xorm:"not null default 0 comment('角色技能等级') INT(5)" json:"skill_level" form:"skill_level" csv:"skill_level"`
	Isdefault   int    `xorm:"not null default 0 comment('是否默认角色 0-否 1-是') TINYINT(1)" json:"isdefault" form:"isdefault" csv:"isdefault"`
	Isfight     int    `xorm:"not null default 0 comment('是否是出战角色 0-否 1-是') TINYINT(1)" json:"isfight" form:"isfight" csv:"isfight"`
	PropEffects string `xorm:"not null default '' comment('装备特效') VARCHAR(200)" json:"prop_effects" form:"prop_effects" csv:"prop_effects"`
	Ltime       int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime       int    `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
