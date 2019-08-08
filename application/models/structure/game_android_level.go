package structure

type GameAndroidLevel struct {
	Gid      int    `xorm:"not null pk default 0 comment('游戏id') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Rtype    int    `xorm:"not null pk default 0 comment('rtype 1 初级 2 中级') TINYINT(4)" json:"rtype" form:"rtype" csv:"rtype"`
	Basepond int    `xorm:"not null default 0 comment('åˆå§‹æ± ') INT(10)" json:"basepond" form:"basepond" csv:"basepond"`
	Curpond  int    `xorm:"not null default 0 comment('å½“å‰æ± ') INT(10)" json:"curpond" form:"curpond" csv:"curpond"`
	Pond     string `xorm:"not null default '' comment('库存值') VARCHAR(100)" json:"pond" form:"pond" csv:"pond"`
	LevelMin string `xorm:"not null default '' comment('机器人最小等级') VARCHAR(100)" json:"level_min" form:"level_min" csv:"level_min"`
	LevelMax string `xorm:"not null default '' comment('机器人最大等级') VARCHAR(100)" json:"level_max" form:"level_max" csv:"level_max"`
}
