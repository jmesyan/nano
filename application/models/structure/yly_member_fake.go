package structure

type YlyMemberFake struct {
	Uid     int `xorm:"not null pk default 0 comment('原UID') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Fakeuid int `xorm:"not null pk default 0 comment('目标uid') INT(11)" json:"fakeuid" form:"fakeuid" csv:"fakeuid"`
	Addtime int `xorm:"not null INT(11)" json:"addtime" form:"addtime" csv:"addtime"`
}
