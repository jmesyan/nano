package structure

type GameCdns struct {
	Bid   string `xorm:"not null pk default '' comment('BUNDLEID') VARCHAR(50)" json:"bid" form:"bid" csv:"bid"`
	Cdn1  string `xorm:"not null default '' comment('正常CDN') VARCHAR(100)" json:"cdn1" form:"cdn1" csv:"cdn1"`
	Cdn2  string `xorm:"not null default '' comment('审核CDN') VARCHAR(100)" json:"cdn2" form:"cdn2" csv:"cdn2"`
	Ver2  string `xorm:"not null default '' comment('审核版本号') VARCHAR(10)" json:"ver2" form:"ver2" csv:"ver2"`
	Flag2 int    `xorm:"not null comment('审核开关') INT(11)" json:"flag2" form:"flag2" csv:"flag2"`
}
