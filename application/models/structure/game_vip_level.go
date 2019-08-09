package structure

type GameVipLevel struct {
	VipLevel      int     `xorm:"not null pk default 0 comment('vip等级') INT(11)" json:"vip_level" form:"vip_level" csv:"vip_level"`
	GameRecharges float32 `xorm:"not null default 0.00 comment('累积充值数量') FLOAT(11,2)" json:"game_recharges" form:"game_recharges" csv:"game_recharges"`
	GameReliefs   int     `xorm:"not null default 0 comment('每日救济金次数') INT(10)" json:"game_reliefs" form:"game_reliefs" csv:"game_reliefs"`
	GameRolegifts int     `xorm:"not null default 0 comment('每天给角色送礼次数') INT(10)" json:"game_rolegifts" form:"game_rolegifts" csv:"game_rolegifts"`
	GameBoxes     int     `xorm:"not null default 0 comment('给用户增加宝箱位个数') INT(10)" json:"game_boxes" form:"game_boxes" csv:"game_boxes"`
	OpenBoxes     int     `xorm:"not null default 0 comment('开启vip宝箱个数') INT(10)" json:"open_boxes" form:"open_boxes" csv:"open_boxes"`
	Wppack        string  `xorm:"not null default '' comment('道具礼包ids-格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
}
