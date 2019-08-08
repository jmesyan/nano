package structure

type UserSendPost struct {
	Uid int `xorm:"not null pk comment('用户uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
}
