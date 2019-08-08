package structure

type ReportPageCount struct {
	Page   string `xorm:"not null pk comment('页面') VARCHAR(30)" json:"page" form:"page" csv:"page"`
	Day    int    `xorm:"not null pk comment('YMD') INT(11)" json:"day" form:"day" csv:"day"`
	Enters int    `xorm:"not null default 0 comment('进入次数') INT(11)" json:"enters" form:"enters" csv:"enters"`
	Downs  int    `xorm:"not null default 0 comment('下载次数') INT(11)" json:"downs" form:"downs" csv:"downs"`
}
