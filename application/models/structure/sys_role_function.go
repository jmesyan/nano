package structure

type SysRoleFunction struct {
	RoleId int    `xorm:"not null pk comment('角色编号') INT(11)" json:"role_id" form:"role_id" csv:"role_id"`
	FuncId int    `xorm:"not null pk comment('功能编号') INT(11)" json:"func_id" form:"func_id" csv:"func_id"`
	FuncOp string `xorm:"comment('功能操作') VARCHAR(100)" json:"func_op" form:"func_op" csv:"func_op"`
}
