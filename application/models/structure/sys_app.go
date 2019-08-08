package structure

type SysApp struct {
	AppId       int    `xorm:"not null pk autoincr comment('应用编号') INT(10)" json:"app_id" form:"app_id" csv:"app_id"`
	AppEname    string `xorm:"comment('应用Code') VARCHAR(50)" json:"app_ename" form:"app_ename" csv:"app_ename"`
	AppName     string `xorm:"comment('应用名称') VARCHAR(100)" json:"app_name" form:"app_name" csv:"app_name"`
	AppImg      string `xorm:"comment('图片') VARCHAR(200)" json:"app_img" form:"app_img" csv:"app_img"`
	AppOrder    int    `xorm:"not null default 0 comment('排序') INT(11)" json:"app_order" form:"app_order" csv:"app_order"`
	AppTreeShow int    `xorm:"not null default 1 comment('是否在导航中显示') TINYINT(1)" json:"app_tree_show" form:"app_tree_show" csv:"app_tree_show"`
	Status      int    `xorm:"not null default 0 comment('状态') TINYINT(1)" json:"status" form:"status" csv:"status"`
}
