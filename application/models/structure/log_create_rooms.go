package structure

type LogCreateRooms struct {
	Lid         int    `xorm:"not null pk autoincr INT(11)" json:"lid" form:"lid" csv:"lid"`
	Code        int    `xorm:"not null comment('房间号') index INT(11)" json:"code" form:"code" csv:"code"`
	Gsid        string `xorm:"not null default '' comment('游戏服务器') VARCHAR(20)" json:"gsid" form:"gsid" csv:"gsid"`
	Tid         int    `xorm:"not null default 0 comment('桌子') INT(11)" json:"tid" form:"tid" csv:"tid"`
	Uid         int    `xorm:"not null default 0 comment('创建者') index INT(11)" json:"uid" form:"uid" csv:"uid"`
	Cards       int    `xorm:"not null default 0 comment('消耗房卡数量') INT(11)" json:"cards" form:"cards" csv:"cards"`
	Ltime       int    `xorm:"not null default 0 comment('创建时间') index index(ltime_2) INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Rounds      int    `xorm:"not null default 0 comment('局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Type        int    `xorm:"not null default 0 comment('类型 1襄阳经典 2襄阳全频道 3宜城跑恰摸') INT(11)" json:"type" form:"type" csv:"type"`
	State       int    `xorm:"not null default 0 comment('状态 (-1创建) 0开始 1结束不退卡 2结束退卡  3创建失败退卡 5/6特殊退卡') index(ltime_2) index TINYINT(4)" json:"state" form:"state" csv:"state"`
	Rtype       int    `xorm:"not null default 0 comment('房间类型 1:6人桌 2:9人桌') TINYINT(4)" json:"rtype" form:"rtype" csv:"rtype"`
	Endtime     int    `xorm:"not null default 0 comment('结束时间') index INT(11)" json:"endtime" form:"endtime" csv:"endtime"`
	Stype       int    `xorm:"not null default 0 comment('飞不飞') TINYINT(4)" json:"stype" form:"stype" csv:"stype"`
	Times       int    `xorm:"not null comment('百搭封顶 巧嘛几倍的花') TINYINT(4)" json:"times" form:"times" csv:"times"`
	Baidad      int    `xorm:"not null default 0 TINYINT(4)" json:"baidad" form:"baidad" csv:"baidad"`
	Huamaima    int    `xorm:"not null default 0 SMALLINT(4)" json:"huamaima" form:"huamaima" csv:"huamaima"`
	Bdwcount    int    `xorm:"not null default 0 TINYINT(4)" json:"bdwcount" form:"bdwcount" csv:"bdwcount"`
	Wft         int    `xorm:"not null default 0 TINYINT(4)" json:"wft" form:"wft" csv:"wft"`
	Fd          int    `xorm:"not null default 0 SMALLINT(4)" json:"fd" form:"fd" csv:"fd"`
	Wdc         int    `xorm:"not null default 0 TINYINT(4)" json:"wdc" form:"wdc" csv:"wdc"`
	Qianggang   int    `xorm:"not null default 0 TINYINT(4)" json:"qianggang" form:"qianggang" csv:"qianggang"`
	Choujiang   int    `xorm:"not null default 0 TINYINT(4)" json:"choujiang" form:"choujiang" csv:"choujiang"`
	Checkip     int    `xorm:"not null default 0 TINYINT(4)" json:"checkIp" form:"checkIp" csv:"checkIp"`
	Checkgps    int    `xorm:"not null default 0 TINYINT(4)" json:"checkGps" form:"checkGps" csv:"checkGps"`
	Bnotdianpao int    `xorm:"not null default 0 comment('百搭点炮不可胡') TINYINT(4)" json:"bnotdianpao" form:"bnotdianpao" csv:"bnotdianpao"`
	Huanum      int    `xorm:"not null default 0 TINYINT(4)" json:"huanum" form:"huanum" csv:"huanum"`
	Flyvalue    int    `xorm:"not null default 0 TINYINT(4)" json:"flyvalue" form:"flyvalue" csv:"flyvalue"`
	Qmnotchi    int    `xorm:"not null default 0 TINYINT(4)" json:"qmnotchi" form:"qmnotchi" csv:"qmnotchi"`
	Paobaida    int    `xorm:"not null default 0 TINYINT(4)" json:"paobaida" form:"paobaida" csv:"paobaida"`
	Mid         int    `xorm:"not null default 0 comment('比赛ID') INT(11)" json:"mid" form:"mid" csv:"mid"`
	Buytype     int    `xorm:"not null default 0 comment('购买类型 0房主 1AA 2代开 3大赢家') TINYINT(4)" json:"buytype" form:"buytype" csv:"buytype"`
	Hid         int    `xorm:"not null default 0 comment('小助手') INT(11)" json:"hid" form:"hid" csv:"hid"`
	Cid         int    `xorm:"not null default 0 comment('家族ID') index index(ltime_2) INT(11)" json:"cid" form:"cid" csv:"cid"`
	Feeuid      int    `xorm:"not null default 0 comment('大赢家扣费id') INT(11)" json:"feeuid" form:"feeuid" csv:"feeuid"`
	Difen       int    `xorm:"not null default 0 comment('底分') INT(11)" json:"difen" form:"difen" csv:"difen"`
	Huangfan    int    `xorm:"not null default 0 comment('底分') INT(11)" json:"huangfan" form:"huangfan" csv:"huangfan"`
	Gid         int    `xorm:"not null default 0 comment('游戏id') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Bijiaotype  int    `xorm:"not null default 0 INT(8)" json:"bijiaotype" form:"bijiaotype" csv:"bijiaotype"`
	Door        int    `xorm:"not null comment('二麻') INT(11)" json:"door" form:"door" csv:"door"`
	Feetype     int    `xorm:"not null default 1 comment('1-扣家族卡(家族)2-扣成员卡(家族) 3 扣成员卡') TINYINT(3)" json:"feetype" form:"feetype" csv:"feetype"`
	Qixiaodui   int    `xorm:"not null default 0 comment('七小对') INT(5)" json:"qixiaodui" form:"qixiaodui" csv:"qixiaodui"`
}
