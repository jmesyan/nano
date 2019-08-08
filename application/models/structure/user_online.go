package structure

type UserOnline struct {
	UserId   int    `xorm:"not null pk comment('在线用户编号') INT(11)" json:"user_id" form:"user_id" csv:"user_id"`
	UserUrl  string `xorm:"comment('当前ＵＲＬ') VARCHAR(200)" json:"user_url" form:"user_url" csv:"user_url"`
	UrlIp    string `xorm:"VARCHAR(45)" json:"url_ip" form:"url_ip" csv:"url_ip"`
	LastTime int    `xorm:"not null default 0 comment('最后一次操作时间') INT(11)" json:"last_time" form:"last_time" csv:"last_time"`
}
