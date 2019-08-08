package structure

type LogPayCallbackFailure struct {
	Cbid    int    `xorm:"not null pk autoincr comment('自动ID') INT(11)" json:"cbid" form:"cbid" csv:"cbid"`
	Type    int    `xorm:"not null default 0 comment('类型') TINYINT(4)" json:"type" form:"type" csv:"type"`
	Apiurl  string `xorm:"not null default '' comment('接口') VARCHAR(200)" json:"apiurl" form:"apiurl" csv:"apiurl"`
	Apidata string `xorm:"not null comment('提交内容') TEXT" json:"apidata" form:"apidata" csv:"apidata"`
	Ltime   int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	State   int    `xorm:"not null default 0 comment('处理状态 0未处理 1已经处理') TINYINT(4)" json:"state" form:"state" csv:"state"`
	Oid     int    `xorm:"not null default 0 comment('订单ID') INT(11)" json:"oid" form:"oid" csv:"oid"`
}
