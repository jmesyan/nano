package structure

type SysAdminUserGame struct {
	UserId int `xorm:"not null pk INT(11)" json:"user_id" form:"user_id" csv:"user_id"`
	GameId int `xorm:"not null pk INT(11)" json:"game_id" form:"game_id" csv:"game_id"`
}
