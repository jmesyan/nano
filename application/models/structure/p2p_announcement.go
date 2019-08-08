package structure

type P2pAnnouncement struct {
	Aid     int    `xorm:"not null pk autoincr comment('公告') INT(11)" json:"aid" form:"aid" csv:"aid"`
	Title   string `xorm:"not null default '' comment('标题') VARCHAR(200)" json:"title" form:"title" csv:"title"`
	Content string `xorm:"not null comment('内容') TEXT" json:"content" form:"content" csv:"content"`
	Ltime   int    `xorm:"not null default 0 comment('时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
}
