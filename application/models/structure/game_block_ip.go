package structure

type GameBlockIp struct {
	Ip          string `xorm:"not null pk default '' comment('禁用ＩＰ 支持正则') VARCHAR(100)" json:"ip" form:"ip" csv:"ip"`
	Regx        int    `xorm:"not null default 0 comment('1正则 0非正则') TINYINT(1)" json:"regx" form:"regx" csv:"regx"`
	ExpiredTime int    `xorm:"not null default 0 comment('过期时间 0不过期 ') INT(11)" json:"expired_time" form:"expired_time" csv:"expired_time"`
	Reason      string `xorm:"comment('禁封原因') VARCHAR(100)" json:"reason" form:"reason" csv:"reason"`
	GmId        int    `xorm:"not null default 0 comment('禁封IP的GM') INT(11)" json:"gm_id" form:"gm_id" csv:"gm_id"`
}
