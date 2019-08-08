package structure

type LogGameTax struct {
	Ldate      int   `xorm:"not null pk comment('日期') INT(8)" json:"ldate" form:"ldate" csv:"ldate"`
	Gid        int   `xorm:"not null pk default 0 comment('游戏id') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Rtype      int   `xorm:"not null pk default 0 comment('房间类型') TINYINT(3)" json:"rtype" form:"rtype" csv:"rtype"`
	AiSystax   int64 `xorm:"not null default 0 comment('机器人税收') BIGINT(20)" json:"ai_systax" form:"ai_systax" csv:"ai_systax"`
	UserSystax int64 `xorm:"not null default 0 comment('玩家税收') BIGINT(20)" json:"user_systax" form:"user_systax" csv:"user_systax"`
	Changes    int64 `xorm:"not null default 0 comment('机器人输赢') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
}
