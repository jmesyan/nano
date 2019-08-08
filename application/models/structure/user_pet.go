package structure

type UserPet struct {
	Uid    int `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Pid    int `xorm:"not null pk default 0 comment('宠物id') INT(11)" json:"pid" form:"pid" csv:"pid"`
	Plevel int `xorm:"not null default 0 comment('宠物等级') INT(11)" json:"plevel" form:"plevel" csv:"plevel"`
	Pscore int `xorm:"not null default 0 comment('宠物评分') INT(10)" json:"pscore" form:"pscore" csv:"pscore"`
}
