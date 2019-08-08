package structure

type UserAchievementForGift struct {
	Uid         int `xorm:"not null pk INT(11)" json:"uid" form:"uid" csv:"uid"`
	Gid         int `xorm:"not null pk default 0 comment('游戏ID') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Rtype       int `xorm:"not null pk default 0 comment('房间类型') INT(11)" json:"rtype" form:"rtype" csv:"rtype"`
	Curlianwins int `xorm:"not null default 0 comment('当前连赢次数') INT(11)" json:"curlianwins" form:"curlianwins" csv:"curlianwins"`
}
