package structure

type GameMjRooms struct {
	Code        int    `xorm:"not null pk comment('房间号') INT(11)" json:"code" form:"code" csv:"code"`
	Gsid        string `xorm:"not null default '' comment('游戏服务器') VARCHAR(20)" json:"gsid" form:"gsid" csv:"gsid"`
	Tid         int    `xorm:"not null default 0 comment('桌子') INT(11)" json:"tid" form:"tid" csv:"tid"`
	Uid         int    `xorm:"not null default 0 comment('创建者') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Cards       int    `xorm:"not null default 0 comment('消耗房卡数量') INT(11)" json:"cards" form:"cards" csv:"cards"`
	Rounds      int    `xorm:"not null default 0 comment('局数') INT(11)" json:"rounds" form:"rounds" csv:"rounds"`
	Type        int    `xorm:"not null default 0 comment('类型 1襄阳经典 2襄阳全频道 3宜城跑恰摸') INT(11)" json:"type" form:"type" csv:"type"`
	Ltime       int    `xorm:"not null default 0 comment('创建时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Lid         int    `xorm:"not null default 0 comment('日志编号') INT(11)" json:"lid" form:"lid" csv:"lid"`
	Stype       int    `xorm:"not null default 0 comment('飞不飞') TINYINT(4)" json:"stype" form:"stype" csv:"stype"`
	Times       int    `xorm:"not null comment('百搭封顶 巧嘛几倍的花') TINYINT(4)" json:"times" form:"times" csv:"times"`
	Baidad      int    `xorm:"not null default 0 TINYINT(4)" json:"baidad" form:"baidad" csv:"baidad"`
	Huamaima    int    `xorm:"not null default 0 SMALLINT(4)" json:"huamaima" form:"huamaima" csv:"huamaima"`
	Bdwcount    int    `xorm:"not null default 0 comment('王最多出几张') TINYINT(4)" json:"bdwcount" form:"bdwcount" csv:"bdwcount"`
	Wft         int    `xorm:"not null default 0 comment('王压2') TINYINT(4)" json:"wft" form:"wft" csv:"wft"`
	Fd          int    `xorm:"not null default 0 comment('封顶数') SMALLINT(4)" json:"fd" form:"fd" csv:"fd"`
	Wdc         int    `xorm:"not null default 0 comment('王单出') TINYINT(4)" json:"wdc" form:"wdc" csv:"wdc"`
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
	Cid         int    `xorm:"not null comment('家族ID') INT(11)" json:"cid" form:"cid" csv:"cid"`
	Difen       int    `xorm:"not null default 0 comment('底分') INT(11)" json:"difen" form:"difen" csv:"difen"`
	Huangfan    int    `xorm:"not null default 0 INT(11)" json:"huangfan" form:"huangfan" csv:"huangfan"`
	Gid         int    `xorm:"not null default 0 comment('游戏ID') INT(11)" json:"gid" form:"gid" csv:"gid"`
	Cur         int    `xorm:"not null default 0 comment('当前局数') INT(8)" json:"cur" form:"cur" csv:"cur"`
	IsOver      int    `xorm:"not null default 0 comment('桌子状态') TINYINT(3)" json:"is_over" form:"is_over" csv:"is_over"`
	Uid1        int    `xorm:"not null comment('用户1') INT(11)" json:"uid1" form:"uid1" csv:"uid1"`
	Uid2        int    `xorm:"not null default 0 comment('用户2') INT(11)" json:"uid2" form:"uid2" csv:"uid2"`
	Uid3        int    `xorm:"not null default 0 comment('用户3') INT(11)" json:"uid3" form:"uid3" csv:"uid3"`
	Uid4        int    `xorm:"not null default 0 comment('用户4') INT(11)" json:"uid4" form:"uid4" csv:"uid4"`
	Uid5        int    `xorm:"not null default 0 comment('用户5') INT(11)" json:"uid5" form:"uid5" csv:"uid5"`
	Uid6        int    `xorm:"not null default 0 comment('用户6') INT(11)" json:"uid6" form:"uid6" csv:"uid6"`
	Bijiaotype  int    `xorm:"not null default 0 INT(8)" json:"bijiaotype" form:"bijiaotype" csv:"bijiaotype"`
	Door        int    `xorm:"not null comment('二麻') INT(11)" json:"door" form:"door" csv:"door"`
	Feetype     int    `xorm:"not null default 1 comment('1-扣家族卡 2-扣成员卡 3 扣成员卡并绑定代理星卡') TINYINT(3)" json:"feetype" form:"feetype" csv:"feetype"`
	Qixiaodui   int    `xorm:"not null default 0 comment('七小对') INT(5)" json:"qixiaodui" form:"qixiaodui" csv:"qixiaodui"`
}
