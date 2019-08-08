package structure

type GamePets struct {
	Id          int    `xorm:"not null pk autoincr comment('宠物id') INT(11)" json:"id" form:"id" csv:"id"`
	Name        string `xorm:"not null default '' comment('宠物名称') VARCHAR(20)" json:"name" form:"name" csv:"name"`
	RoomImg     string `xorm:"not null default '' comment('大厅立绘') VARCHAR(100)" json:"room_img" form:"room_img" csv:"room_img"`
	RoomAnimate string `xorm:"not null default '' comment('大厅动画') VARCHAR(100)" json:"room_animate" form:"room_animate" csv:"room_animate"`
	GameImg     string `xorm:"not null default '' comment('游戏内立绘') VARCHAR(100)" json:"game_img" form:"game_img" csv:"game_img"`
	GameAnimate string `xorm:"not null default '' comment('游戏内动画') VARCHAR(100)" json:"game_animate" form:"game_animate" csv:"game_animate"`
	Mutual      int    `xorm:"not null default 0 comment('交互气泡') INT(11)" json:"mutual" form:"mutual" csv:"mutual"`
	Effect      string `xorm:"not null default '' comment('特效描述') VARCHAR(200)" json:"effect" form:"effect" csv:"effect"`
	EffectNum   int    `xorm:"not null default 0 comment('特效倍数') INT(11)" json:"effect_num" form:"effect_num" csv:"effect_num"`
}
