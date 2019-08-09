package structure

type YlyPost struct {
	Pid      int    `xorm:"not null pk autoincr INT(11)" json:"pid" form:"pid" csv:"pid"`
	Fromid   int    `xorm:"not null default 0 comment('寄件人') index INT(11)" json:"fromid" form:"fromid" csv:"fromid"`
	Toid     int    `xorm:"not null default 0 comment('收件人') index INT(11)" json:"toid" form:"toid" csv:"toid"`
	Fname    string `xorm:"not null default '' comment('寄件人姓名') VARCHAR(30)" json:"fname" form:"fname" csv:"fname"`
	Sendtime int    `xorm:"not null default 0 comment('寄件时间') index INT(11)" json:"sendtime" form:"sendtime" csv:"sendtime"`
	Readtime int    `xorm:"not null pk default 0 comment('收件时间') index INT(11)" json:"readtime" form:"readtime" csv:"readtime"`
	Dateline int    `xorm:"not null default 0 comment('失效时间') INT(11)" json:"dateline" form:"dateline" csv:"dateline"`
	Ps       string `xorm:"not null comment('信件内容') TEXT" json:"ps" form:"ps" csv:"ps"`
	State    int    `xorm:"not null default 0 comment('0:寄出;1:读取;2:需要付钱取物品;3：物品已经取出') TINYINT(2)" json:"state" form:"state" csv:"state"`
	Title    int    `xorm:"not null comment('1.签到奖励 2.赛季奖励 3.段位排行 4.金币排行 5.维护补偿 6.背包已满 7.福利回馈 8.节日有礼 9.天降福利 10.充值回馈 11.回归好礼 12.累计登陆 13.好友邀请') TINYINT(2)" json:"title" form:"title" csv:"title"`
	Subject  string `xorm:"not null default '' comment('主题') VARCHAR(300)" json:"subject" form:"subject" csv:"subject"`
	Admin    int    `xorm:"not null default 0 comment('管理员') INT(11)" json:"admin" form:"admin" csv:"admin"`
	Status   int    `xorm:"not null default 0 comment('邮件数量类型 0:单人邮件 1:群发邮件 2:系统邮件') TINYINT(2)" json:"status" form:"status" csv:"status"`
	Type     int    `xorm:"not null default 0 comment('邮件属性类型 0-普通邮件 1-道具邮件') TINYINT(3)" json:"type" form:"type" csv:"type"`
	Attach   string `xorm:"default '' comment('附件') VARCHAR(255)" json:"attach" form:"attach" csv:"attach"`
	Golds    int    `xorm:"not null default 0 INT(11)" json:"golds" form:"golds" csv:"golds"`
	Golds2   int    `xorm:"not null default 0 INT(11)" json:"golds2" form:"golds2" csv:"golds2"`
	Wppack   string `xorm:"not null default '' comment('道具礼包ids,格式 pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Isshow   int    `xorm:"not null default 1 comment('是否显示 0否 1是') TINYINT(1)" json:"isShow" form:"isShow" csv:"isShow"`
}
