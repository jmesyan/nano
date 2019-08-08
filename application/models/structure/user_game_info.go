package structure

type UserGameInfo struct {
	Uid           int `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Totalrounds   int `xorm:"not null default 0 comment('总局数') INT(11)" json:"totalrounds" form:"totalrounds" csv:"totalrounds"`
	Winrounds     int `xorm:"not null default 0 comment('胜利局数') INT(11)" json:"winrounds" form:"winrounds" csv:"winrounds"`
	Curlianwin    int `xorm:"not null default 0 comment('当前连胜') INT(11)" json:"curlianwin" form:"curlianwin" csv:"curlianwin"`
	Maxlianwin    int `xorm:"not null default 0 comment('最高连胜') INT(11)" json:"maxlianwin" form:"maxlianwin" csv:"maxlianwin"`
	Farmerturns   int `xorm:"not null default 0 comment('农民拿到牌权次数') INT(11)" json:"farmerturns" form:"farmerturns" csv:"farmerturns"`
	Totalturns    int `xorm:"not null default 0 comment('所有出牌权数') INT(11)" json:"totalturns" form:"totalturns" csv:"totalturns"`
	Doubleturns   int `xorm:"not null default 0 comment('加倍局数') INT(11)" json:"doubleturns" form:"doubleturns" csv:"doubleturns"`
	Opratetime    int `xorm:"not null default 0 comment('出牌时间和') INT(11)" json:"opratetime" form:"opratetime" csv:"opratetime"`
	Opratecounts  int `xorm:"not null default 0 comment('出牌次数') INT(11)" json:"opratecounts" form:"opratecounts" csv:"opratecounts"`
	Yazhicounts   int `xorm:"not null default 0 comment('压制次数') INT(11)" json:"yazhicounts" form:"yazhicounts" csv:"yazhicounts"`
	Landerrounds  int `xorm:"not null default 0 comment('地主次数') INT(11)" json:"landerrounds" form:"landerrounds" csv:"landerrounds"`
	Maxgraderanks int `xorm:"not null default 0 comment('最高段位排名') INT(10)" json:"maxgraderanks" form:"maxgraderanks" csv:"maxgraderanks"`
}
