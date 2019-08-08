package structure

type SysRole struct {
	RoleId        int    `xorm:"not null pk autoincr comment('角色') INT(10)" json:"role_id" form:"role_id" csv:"role_id"`
	RoleName      string `xorm:"comment('角色名称') VARCHAR(100)" json:"role_name" form:"role_name" csv:"role_name"`
	RoleEname     string `xorm:"comment('角色代码') VARCHAR(50)" json:"role_ename" form:"role_ename" csv:"role_ename"`
	RoleFuncnames string `xorm:"comment('角色功能') VARCHAR(3000)" json:"role_funcnames" form:"role_funcnames" csv:"role_funcnames"`
	RoleFuncids   string `xorm:"comment('角色功能代码') VARCHAR(3000)" json:"role_funcids" form:"role_funcids" csv:"role_funcids"`
	Status        int    `xorm:"not null default 0 comment('状态') TINYINT(1)" json:"status" form:"status" csv:"status"`
}
