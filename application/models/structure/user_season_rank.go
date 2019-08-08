package structure

type UserSeasonRank struct {
	Season     int    `xorm:"not null pk comment('赛季') INT(10)" json:"season" form:"season" csv:"season"`
	Uid        int    `xorm:"not null pk default 0 comment('用户编号') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Nickname   string `xorm:"not null default '' comment('昵称') VARCHAR(20)" json:"nickname" form:"nickname" csv:"nickname"`
	UseHead    int    `xorm:"not null default 0 comment('头像ID') INT(11)" json:"use_head" form:"use_head" csv:"use_head"`
	UseAvatar  int    `xorm:"not null default 0 comment('头像框id') INT(11)" json:"use_avatar" form:"use_avatar" csv:"use_avatar"`
	GradeStars int    `xorm:"not null default 0 comment('段位星级') INT(11)" json:"grade_stars" form:"grade_stars" csv:"grade_stars"`
	GradeScore int    `xorm:"not null default 0 comment('段位积分') INT(11)" json:"grade_score" form:"grade_score" csv:"grade_score"`
	GradeRank  int    `xorm:"not null default 0 comment('赛季排名') INT(10)" json:"grade_rank" form:"grade_rank" csv:"grade_rank"`
}
