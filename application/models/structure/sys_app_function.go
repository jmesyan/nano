package structure

type SysAppFunction struct {
	FuncId    int    `xorm:"not null pk autoincr comment('功能编号') INT(10)" json:"func_id" form:"func_id" csv:"func_id"`
	AppId     int    `xorm:"not null comment('应用编号') INT(11)" json:"app_id" form:"app_id" csv:"app_id"`
	FuncName  string `xorm:"comment('功能名称') VARCHAR(50)" json:"func_name" form:"func_name" csv:"func_name"`
	FuncEname string `xorm:"comment('功能代码') VARCHAR(100)" json:"func_ename" form:"func_ename" csv:"func_ename"`
	FuncUrl   string `xorm:"comment('地址') VARCHAR(200)" json:"func_url" form:"func_url" csv:"func_url"`
	FuncImg   string `xorm:"comment('图标') VARCHAR(200)" json:"func_img" form:"func_img" csv:"func_img"`
	FuncOrder int    `xorm:"not null default 0 comment('排序') INT(11)" json:"func_order" form:"func_order" csv:"func_order"`
	Status    int    `xorm:"not null default 1 comment('状态') INT(3)" json:"status" form:"status" csv:"status"`
}
