package structure

type LogAdminOp struct {
	Id      int    `xorm:"not null pk autoincr INT(11)" json:"id" form:"id" csv:"id"`
	Uid     int    `xorm:"default 0 comment('用户编号') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Url     string `xorm:"default '' comment('请求ＵＲＬ') VARCHAR(255)" json:"url" form:"url" csv:"url"`
	Request string `xorm:"comment('操作内容') TEXT" json:"request" form:"request" csv:"request"`
	Ip      string `xorm:"comment('ip') VARCHAR(20)" json:"ip" form:"ip" csv:"ip"`
	Ltime   int    `xorm:"default 0 comment('操作时间') index INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
