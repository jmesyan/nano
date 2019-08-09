package structure

type UserGameGuidance struct {
	Uid   int    `xorm:"not null pk comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Stage string `xorm:"not null default '' comment('引导阶段 , 分隔') VARCHAR(50)" json:"stage" form:"stage" csv:"stage"`
}
