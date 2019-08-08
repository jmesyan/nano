package structure

type UserWincount struct {
	Date        int    `xorm:"not null pk comment('日期') INT(11)" json:"date" form:"date" csv:"date"`
	Uid         int    `xorm:"not null pk comment('用户ID') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gid         int    `xorm:"not null pk comment('用户当前连赢对应的游戏ID') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Maxlianwins int    `xorm:"default 0 comment('用户本日最大连赢') INT(11)" json:"maxlianwins" form:"maxlianwins" csv:"maxlianwins"`
	Curlianwins int    `xorm:"default 0 comment('用户当前连赢') INT(11)" json:"curlianwins" form:"curlianwins" csv:"curlianwins"`
	Time        int    `xorm:"comment('最后更新时间') INT(11)" json:"time" form:"time" csv:"time"`
	Ltime       int    `xorm:"not null default 0 INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	State       int    `xorm:"not null default 0 comment('0 未领取 1 已领取未发放  2 已领取已发') INT(2)" json:"state" form:"state" csv:"state"`
	Phone       string `xorm:"comment('手机号') VARCHAR(20)" json:"phone" form:"phone" csv:"phone"`
	Address     string `xorm:"comment('收货地址') VARCHAR(255)" json:"address" form:"address" csv:"address"`
	Name        string `xorm:"comment('姓名') VARCHAR(50)" json:"name" form:"name" csv:"name"`
	Oid         string `xorm:"comment('快递信息') VARCHAR(100)" json:"oid" form:"oid" csv:"oid"`
	Award       string `xorm:"default '' comment('奖励') VARCHAR(200)" json:"award" form:"award" csv:"award"`
	Rank        int    `xorm:"not null default 0 comment('名次') INT(11)" json:"rank" form:"rank" csv:"rank"`
	Atype       int    `xorm:"default 0 INT(11)" json:"atype" form:"atype" csv:"atype"`
	Anum        int    `xorm:"not null default 0 INT(11)" json:"anum" form:"anum" csv:"anum"`
}
