package structure

type UserGoldsRank struct {
	Uid       int    `xorm:"not null pk default 0 comment('用户编号') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Nickname  string `xorm:"not null default '' comment('昵称') VARCHAR(20)" json:"nickname" form:"nickname" csv:"nickname"`
	UseHead   int    `xorm:"not null default 0 comment('头像ID') INT(11)" json:"use_head" form:"use_head" csv:"use_head"`
	UseAvatar int    `xorm:"not null default 0 comment('头像框ID') INT(11)" json:"use_avatar" form:"use_avatar" csv:"use_avatar"`
	Golds     int    `xorm:"not null comment('金币') INT(11)" json:"golds" form:"golds" csv:"golds"`
}
