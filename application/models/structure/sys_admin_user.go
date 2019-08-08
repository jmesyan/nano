package structure

type SysAdminUser struct {
	UserId   int    `xorm:"not null pk comment('用户') INT(11)" json:"user_id" form:"user_id" csv:"user_id"`
	RoleId   int    `xorm:"not null pk comment('角色') INT(11)" json:"role_id" form:"role_id" csv:"role_id"`
	Status   int    `xorm:"not null default 1 comment('状态') TINYINT(1)" json:"status" form:"status" csv:"status"`
	UserGame string `xorm:"comment('游戏') VARCHAR(500)" json:"user_game" form:"user_game" csv:"user_game"`
}
