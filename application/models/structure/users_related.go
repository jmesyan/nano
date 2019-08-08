package structure

type UsersRelated struct {
	//表game_userfield
	Uid         int    `xorm:"not null pk INT(10)" json:"uid" form:"uid" csv:"uid"`
	Changes     int64  `xorm:"not null default 0 comment('玩家金币总流水') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
	Startdumb   int    `xorm:"not null default 0 comment('禁言开始') index INT(10)" json:"startdumb" form:"startdumb" csv:"startdumb"`
	Dumblimit   int    `xorm:"not null default 0 comment('禁言结束') INT(10)" json:"dumblimit" form:"dumblimit" csv:"dumblimit"`
	State       int    `xorm:"not null default 0 comment('0:正常;1:禁止私聊;2:禁止聊天;3:禁止旁观;4:禁止进入游戏;5:冻结') TINYINT(3)" json:"state" form:"state" csv:"state"`
	Online      int    `xorm:"not null default 0 comment('在线') TINYINT(1)" json:"online" form:"online" csv:"online"`
	OnlineTimes int    `xorm:"not null default 0 comment('在线总时长') INT(11)" json:"online_times" form:"online_times" csv:"online_times"`
	GameTimes   int    `xorm:"not null default 0 comment('游戏总时长') INT(11)" json:"game_times" form:"game_times" csv:"game_times"`
	Ismotor     int    `xorm:"not null default 0 comment('是否机器人') index TINYINT(1)" json:"ismotor" form:"ismotor" csv:"ismotor"`
	PayGolds    int64  `xorm:"not null default 0 comment('充值总房卡') BIGINT(20)" json:"pay_golds" form:"pay_golds" csv:"pay_golds"`
	PayStones   int64  `xorm:"not null default 0 comment('充值总宝石') BIGINT(20)" json:"pay_stones" form:"pay_stones" csv:"pay_stones"`
	PayGolds2   int64  `xorm:"not null default 0 comment('充值总金币') BIGINT(20)" json:"pay_golds2" form:"pay_golds2" csv:"pay_golds2"`
	RoomCard    int    `xorm:"not null default 0 comment('房卡数量') INT(11)" json:"room_card" form:"room_card" csv:"room_card"`
	Ticket      int    `xorm:"not null default 0 comment('礼券数量') INT(11)" json:"ticket" form:"ticket" csv:"ticket"`
	Stones      int    `xorm:"not null default 0 comment('宝石') INT(11)" json:"stones" form:"stones" csv:"stones"`
	CardUse     int    `xorm:"not null default 0 comment('消耗房卡总计') INT(11)" json:"card_use" form:"card_use" csv:"card_use"`
	TicketUse   int    `xorm:"not null default 0 comment('消耗礼券总计') INT(11)" json:"ticket_use" form:"ticket_use" csv:"ticket_use"`
	StonesUse   int    `xorm:"not null default 0 comment('宝石使用总计') INT(11)" json:"stones_use" form:"stones_use" csv:"stones_use"`
	Level       int    `xorm:"not null default 0 comment('等级') INT(10)" json:"level" form:"level" csv:"level"`
	Experience  int    `xorm:"not null default 0 comment('经验值') INT(10)" json:"experience" form:"experience" csv:"experience"`
	VipLevel    int    `xorm:"not null default 0 comment('vip等级') index(vip_level) TINYINT(3)" json:"vip_level" form:"vip_level" csv:"vip_level"`
	VipPay      string `xorm:"not null default 0.00 comment('在当前VIP充值金额') DECIMAL(18,2)" json:"vip_pay" form:"vip_pay" csv:"vip_pay"`
	VipCard     int    `xorm:"not null default 0 comment('vip消耗的房卡') INT(11)" json:"vip_card" form:"vip_card" csv:"vip_card"`
	VipTime     int    `xorm:"not null default 0 comment('成为VIP的时间') INT(10)" json:"vip_time" form:"vip_time" csv:"vip_time"`
	VipExpire   int    `xorm:"not null default 0 comment('过期时间') index(vip_level) INT(10)" json:"vip_expire" form:"vip_expire" csv:"vip_expire"`
	VipRounds   int    `xorm:"not null default 0 comment('当前局数') INT(10)" json:"vip_rounds" form:"vip_rounds" csv:"vip_rounds"`
	VipLstate   int    `xorm:"not null default 0 comment('最后一次 1升2保3降') TINYINT(3)" json:"vip_lstate" form:"vip_lstate" csv:"vip_lstate"`
	VipLtime    int    `xorm:"not null default 0 comment('最后一次 升保降时间') INT(10)" json:"vip_ltime" form:"vip_ltime" csv:"vip_ltime"`
	VipExps     int    `xorm:"not null default 0 comment('vip经验') INT(10)" json:"vip_exps" form:"vip_exps" csv:"vip_exps"`
	UseProp     int    `xorm:"not null default 0 comment('使用的道具') INT(10)" json:"use_prop" form:"use_prop" csv:"use_prop"`
	UseChat     int    `xorm:"not null default 0 comment('使用聊天框') INT(10)" json:"use_chat" form:"use_chat" csv:"use_chat"`
	UseAvatar   int    `xorm:"not null default 0 comment('使用头像框') INT(10)" json:"use_avatar" form:"use_avatar" csv:"use_avatar"`
	Btimes      int    `xorm:"not null default 0 comment('累计红包奖励次数') INT(5)" json:"btimes" form:"btimes" csv:"btimes"`
	Bonus       int64  `xorm:"not null default 0 comment('累计红包奖励数量') BIGINT(20)" json:"bonus" form:"bonus" csv:"bonus"`
	Golds       int64  `xorm:"not null default 0 comment('金币') BIGINT(20)" json:"golds" form:"golds" csv:"golds"`
	BonusGolds  int64  `xorm:"not null default 0 comment('金币变化量') BIGINT(20)" json:"bonus_golds" form:"bonus_golds" csv:"bonus_golds"`
	Lucky       int    `xorm:"not null default 0 INT(11)" json:"lucky" form:"lucky" csv:"lucky"`
	Curlianwins int    `xorm:"not null default 0 comment('当前连赢局数') INT(11)" json:"curlianwins" form:"curlianwins" csv:"curlianwins"`
	PropsCap    int    `xorm:"not null default 0 comment('道具容纳能力') INT(8)" json:"props_cap" form:"props_cap" csv:"props_cap"`
	GradeSeason int    `xorm:"not null default 0 comment('段位赛季') INT(10)" json:"grade_season" form:"grade_season" csv:"grade_season"`
	GradeStars  int    `xorm:"not null default 1 comment('段位星级') INT(10)" json:"grade_stars" form:"grade_stars" csv:"grade_stars"`
	GradeScore  int    `xorm:"not null default 0 comment('段位星级积分') INT(10)" json:"grade_score" form:"grade_score" csv:"grade_score"`
	FightRole   int    `xorm:"not null default 0 comment('出战角色id') INT(11)" json:"fight_role" form:"fight_role" csv:"fight_role"`
	BoxesCap    int    `xorm:"not null default 0 comment('有cd普通宝箱容纳能力') INT(8)" json:"boxes_cap" form:"boxes_cap" csv:"boxes_cap"`
	//表yly_member
	Pid        string `xorm:"not null default '' comment('平台id') CHAR(100)" json:"pid" form:"pid" csv:"pid"`
	Username   string `xorm:"not null default '' comment('用户名') unique VARCHAR(64)" json:"username" form:"username" csv:"username"`
	Password   string `xorm:"not null default '' comment('密码') CHAR(32)" json:"password" form:"password" csv:"password"`
	Tel        string `xorm:"not null default '' comment('手机') VARCHAR(20)" json:"tel" form:"tel" csv:"tel"`
	Email      string `xorm:"not null default '' comment('邮箱') CHAR(32)" json:"email" form:"email" csv:"email"`
	Birthday   string `xorm:"not null default '' comment('生日') VARCHAR(50)" json:"birthday" form:"birthday" csv:"birthday"`
	RegIp      string `xorm:"not null default '' comment('注册ip') CHAR(15)" json:"reg_ip" form:"reg_ip" csv:"reg_ip"`
	RegDate    int    `xorm:"not null default 0 comment('注册时间') index INT(11)" json:"reg_date" form:"reg_date" csv:"reg_date"`
	Gender     string `xorm:"not null default 'm' comment('性别') ENUM('f','m')" json:"gender" form:"gender" csv:"gender"`
	Utype      string `xorm:"not null default '' comment('用户类型') CHAR(10)" json:"utype" form:"utype" csv:"utype"`
	Nickname   string `xorm:"not null default '' comment('昵称') CHAR(20)" json:"nickname" form:"nickname" csv:"nickname"`
	GroupId    int    `xorm:"not null default 5 comment('管理级别') TINYINT(3)" json:"group_id" form:"group_id" csv:"group_id"`
	Locale     string `xorm:"not null comment('语言') VARCHAR(20)" json:"locale" form:"locale" csv:"locale"`
	Avatar     string `xorm:"not null comment('头像') VARCHAR(255)" json:"avatar" form:"avatar" csv:"avatar"`
	Upuid      int    `xorm:"not null default 0 comment('上层uid，是哪个用户推荐过来的') INT(11)" json:"upuid" form:"upuid" csv:"upuid"`
	Ad         string `xorm:"not null default '' comment('用户来源，哪个广告') index CHAR(100)" json:"ad" form:"ad" csv:"ad"`
	LoginIp    string `xorm:"not null default '' comment('登录ip') CHAR(15)" json:"login_ip" form:"login_ip" csv:"login_ip"`
	LoginTimes int    `xorm:"not null default 0 comment('登录次数') INT(11)" json:"login_times" form:"login_times" csv:"login_times"`
	LoginDate  int    `xorm:"not null default 0 comment('最后一次登陆时间') index INT(10)" json:"login_date" form:"login_date" csv:"login_date"`
	ForbidTime int    `xorm:"not null default 0 comment('封号时间') INT(10)" json:"forbid_time" form:"forbid_time" csv:"forbid_time"`
	Remark     string `xorm:"not null default '' comment('封号备注') VARCHAR(255)" json:"remark" form:"remark" csv:"remark"`
	Sign       string `xorm:"not null default '' comment('签名') VARCHAR(255)" json:"sign" form:"sign" csv:"sign"`
	Regcity    string `xorm:"not null default '' comment('城市') VARCHAR(45)" json:"regcity" form:"regcity" csv:"regcity"`
	Regarea    string `xorm:"not null default '' comment('区域') VARCHAR(100)" json:"regarea" form:"regarea" csv:"regarea"`
	NewUser    int    `xorm:"not null default 1 comment('新用户') TINYINT(1)" json:"new_user" form:"new_user" csv:"new_user"`
	Ps         string `xorm:"comment('pass') VARCHAR(100)" json:"ps" form:"ps" csv:"ps"`
	UserSalt   string `xorm:"not null default '' comment('salt') CHAR(6)" json:"user_salt" form:"user_salt" csv:"user_salt"`
	Appid      int    `xorm:"not null default 0 comment('appid') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Logincity  string `xorm:"not null comment('登录地') VARCHAR(50)" json:"logincity" form:"logincity" csv:"logincity"`
	Subutype   string `xorm:"not null default '' VARCHAR(50)" json:"subutype" form:"subutype" csv:"subutype"`
}
