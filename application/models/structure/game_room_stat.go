package structure

type GameRoomStat struct {
	Gid            int    `xorm:"not null pk default 8 comment('遊戲id') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Rtype          int    `xorm:"not null pk default 1 comment('房間類型') INT(10)" json:"rtype" form:"rtype" csv:"rtype"`
	Gametype       int    `xorm:"not null pk default 0 comment('玩法 0上海敲麻玩法／1上海百搭玩法／2上海清混碰／3上海拉西胡／4红中麻将／5金山拦廿胡／6松江搓花玩法') TINYINT(3)" json:"gametype" form:"gametype" csv:"gametype"`
	S2g            int    `xorm:"not null default 0 comment('1分对应多少金币') INT(10)" json:"s2g" form:"s2g" csv:"s2g"`
	Systax         int64  `xorm:"not null default 0 comment('系统税收') BIGINT(20)" json:"systax" form:"systax" csv:"systax"`
	RoomName       string `xorm:"default ' ' comment('房间名') VARCHAR(512)" json:"room_name" form:"room_name" csv:"room_name"`
	Goldsless      int    `xorm:"not null default 0 comment('踢出条件') INT(10)" json:"goldsless" form:"goldsless" csv:"goldsless"`
	Goldsmin       int    `xorm:"default 0 comment('最小进入') INT(10)" json:"goldsmin" form:"goldsmin" csv:"goldsmin"`
	Goldsmax       int    `xorm:"not null default 0 comment('金币上限') INT(10)" json:"goldsmax" form:"goldsmax" csv:"goldsmax"`
	Wins           int    `xorm:"not null default 0 comment('掉落宝箱所需的连胜次数') INT(3)" json:"wins" form:"wins" csv:"wins"`
	Playernum      int    `xorm:"not null default 0 comment('玩家数量') INT(3)" json:"playernum" form:"playernum" csv:"playernum"`
	Androidnum     int    `xorm:"not null default 0 comment('初始机器人数量') INT(10)" json:"androidnum" form:"androidnum" csv:"androidnum"`
	Increasenum    int    `xorm:"not null default 0 comment('机器增量') INT(10)" json:"increasenum" form:"increasenum" csv:"increasenum"`
	Timefactor1    int    `xorm:"not null default 0 comment('时间参数1') INT(10)" json:"timefactor1" form:"timefactor1" csv:"timefactor1"`
	Timefactor2    int    `xorm:"not null default 0 comment('时间参数2(千分之)') INT(10)" json:"timefactor2" form:"timefactor2" csv:"timefactor2"`
	Goldfatcot3    int    `xorm:"not null default 0 comment('金币参数3') INT(10)" json:"goldfatcot3" form:"goldfatcot3" csv:"goldfatcot3"`
	Goldfatcot4    int    `xorm:"not null default 0 comment('金币参数4(千分之)') INT(10)" json:"goldfatcot4" form:"goldfatcot4" csv:"goldfatcot4"`
	Roomgoldfactor int    `xorm:"not null default 0 comment('房间等级金币权重') INT(10)" json:"roomgoldfactor" form:"roomgoldfactor" csv:"roomgoldfactor"`
	Roomtimefactor int    `xorm:"not null default 0 comment('房间等级时间权重(千分之)') INT(10)" json:"roomtimefactor" form:"roomtimefactor" csv:"roomtimefactor"`
}
