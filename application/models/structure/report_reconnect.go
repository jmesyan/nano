package structure

type ReportReconnect struct {
	Ldate int `xorm:"not null pk default 0 comment('日期') INT(10)" json:"ldate" form:"ldate" csv:"ldate"`
	Type  int `xorm:"not null pk default 0 comment('1 客户端报错重连 2 数据错误重连 3 收不到后续消息 4 断网重连 5 消息计数对不上 6 链接超时 7 502  8 游戏中重连') TINYINT(3)" json:"type" form:"type" csv:"type"`
	Users int `xorm:"not null default 0 comment('用户数') INT(10)" json:"users" form:"users" csv:"users"`
}
