package structure

type GameUserfield struct {
	Uid           int     `xorm:"not null pk INT(10)" json:"uid" form:"uid" csv:"uid"`
	Changes       int64   `xorm:"not null default 0 comment('玩家金币总流水') BIGINT(20)" json:"changes" form:"changes" csv:"changes"`
	Startdumb     int     `xorm:"not null default 0 comment('禁言开始') index INT(10)" json:"startdumb" form:"startdumb" csv:"startdumb"`
	Dumblimit     int     `xorm:"not null default 0 comment('禁言结束') INT(10)" json:"dumblimit" form:"dumblimit" csv:"dumblimit"`
	State         int     `xorm:"not null default 0 comment('0:正常;1:禁止私聊;2:禁止聊天;3:禁止旁观;4:禁止进入游戏;5:冻结') TINYINT(3)" json:"state" form:"state" csv:"state"`
	Online        int     `xorm:"not null default 0 comment('在线') TINYINT(1)" json:"online" form:"online" csv:"online"`
	OnlineTimes   int     `xorm:"not null default 0 comment('在线总时长') INT(11)" json:"online_times" form:"online_times" csv:"online_times"`
	GameTimes     int     `xorm:"not null default 0 comment('游戏总时长') INT(11)" json:"game_times" form:"game_times" csv:"game_times"`
	Ismotor       int     `xorm:"not null default 0 comment('是否机器人') index TINYINT(1)" json:"ismotor" form:"ismotor" csv:"ismotor"`
	PayGolds      int64   `xorm:"not null default 0 comment('充值总房卡') BIGINT(20)" json:"pay_golds" form:"pay_golds" csv:"pay_golds"`
	PayStones     int64   `xorm:"not null default 0 comment('充值总宝石') BIGINT(20)" json:"pay_stones" form:"pay_stones" csv:"pay_stones"`
	PayFirst      float32 `xorm:"not null default 0.00 comment('首充金额') FLOAT(11,2)" json:"pay_first" form:"pay_first" csv:"pay_first"`
	RoomCard      int     `xorm:"not null default 0 comment('房卡数量') INT(11)" json:"room_card" form:"room_card" csv:"room_card"`
	Ticket        int     `xorm:"not null default 0 comment('礼券数量') INT(11)" json:"ticket" form:"ticket" csv:"ticket"`
	Stones        int     `xorm:"not null default 0 comment('宝石') INT(11)" json:"stones" form:"stones" csv:"stones"`
	CardUse       int     `xorm:"not null default 0 comment('消耗房卡总计') INT(11)" json:"card_use" form:"card_use" csv:"card_use"`
	TicketUse     int     `xorm:"not null default 0 comment('消耗礼券总计') INT(11)" json:"ticket_use" form:"ticket_use" csv:"ticket_use"`
	StonesUse     int     `xorm:"not null default 0 comment('宝石使用总计') INT(11)" json:"stones_use" form:"stones_use" csv:"stones_use"`
	Level         int     `xorm:"not null default 0 comment('等级') INT(10)" json:"level" form:"level" csv:"level"`
	Experience    int     `xorm:"not null default 0 comment('经验值') INT(10)" json:"experience" form:"experience" csv:"experience"`
	VipLevel      int     `xorm:"not null default 0 comment('vip等级') index(vip_level) TINYINT(3)" json:"vip_level" form:"vip_level" csv:"vip_level"`
	VipPay        float32 `xorm:"not null default 0.00 comment('在当前VIP充值金额') FLOAT(11,2)" json:"vip_pay" form:"vip_pay" csv:"vip_pay"`
	VipCard       int     `xorm:"not null default 0 comment('vip消耗的房卡') INT(11)" json:"vip_card" form:"vip_card" csv:"vip_card"`
	VipTime       int     `xorm:"not null default 0 comment('成为VIP的时间') INT(10)" json:"vip_time" form:"vip_time" csv:"vip_time"`
	VipExpire     int     `xorm:"not null default 0 comment('过期时间') index(vip_level) INT(10)" json:"vip_expire" form:"vip_expire" csv:"vip_expire"`
	VipRounds     int     `xorm:"not null default 0 comment('当前局数') INT(10)" json:"vip_rounds" form:"vip_rounds" csv:"vip_rounds"`
	VipLstate     int     `xorm:"not null default 0 comment('最后一次 1升2保3降') TINYINT(3)" json:"vip_lstate" form:"vip_lstate" csv:"vip_lstate"`
	VipLtime      int     `xorm:"not null default 0 comment('最后一次 升保降时间') INT(10)" json:"vip_ltime" form:"vip_ltime" csv:"vip_ltime"`
	VipExps       int     `xorm:"not null default 0 comment('vip经验') INT(10)" json:"vip_exps" form:"vip_exps" csv:"vip_exps"`
	UseProp       int     `xorm:"not null default 0 comment('使用的道具') INT(10)" json:"use_prop" form:"use_prop" csv:"use_prop"`
	UseChat       int     `xorm:"not null default 0 comment('使用聊天框') INT(10)" json:"use_chat" form:"use_chat" csv:"use_chat"`
	UseAvatar     int     `xorm:"not null default 0 comment('使用头像框') INT(10)" json:"use_avatar" form:"use_avatar" csv:"use_avatar"`
	UseHead       int     `xorm:"not null default 0 comment('使用头像') INT(11)" json:"use_head" form:"use_head" csv:"use_head"`
	Btimes        int     `xorm:"not null default 0 comment('累计红包奖励次数') INT(5)" json:"btimes" form:"btimes" csv:"btimes"`
	Bonus         int64   `xorm:"not null default 0 comment('累计红包奖励数量') BIGINT(20)" json:"bonus" form:"bonus" csv:"bonus"`
	Golds         int64   `xorm:"not null default 0 comment('金币') BIGINT(20)" json:"golds" form:"golds" csv:"golds"`
	BonusGolds    int64   `xorm:"not null default 0 comment('金币变化量') BIGINT(20)" json:"bonus_golds" form:"bonus_golds" csv:"bonus_golds"`
	Lucky         int     `xorm:"not null default 0 INT(11)" json:"lucky" form:"lucky" csv:"lucky"`
	Curlianwins   int     `xorm:"not null default 0 comment('当前连赢局数') INT(11)" json:"curlianwins" form:"curlianwins" csv:"curlianwins"`
	PropsCap      int     `xorm:"not null default 0 comment('道具容纳能力') INT(8)" json:"props_cap" form:"props_cap" csv:"props_cap"`
	GradeSeason   int     `xorm:"not null default 0 comment('段位赛季') INT(10)" json:"grade_season" form:"grade_season" csv:"grade_season"`
	GradeStars    int     `xorm:"not null default 1 comment('段位星级') INT(10)" json:"grade_stars" form:"grade_stars" csv:"grade_stars"`
	GradeScore    int     `xorm:"not null default 0 comment('段位星级积分') INT(10)" json:"grade_score" form:"grade_score" csv:"grade_score"`
	GradeStarmax  int     `xorm:"not null default 0 comment('历史最高星级') INT(10)" json:"grade_starmax" form:"grade_starmax" csv:"grade_starmax"`
	FightRole     int     `xorm:"not null default 0 comment('出战角色id') INT(11)" json:"fight_role" form:"fight_role" csv:"fight_role"`
	BoxesCap      int     `xorm:"not null default 0 comment('有cd普通宝箱容纳能力') INT(8)" json:"boxes_cap" form:"boxes_cap" csv:"boxes_cap"`
	RolesAuth     int     `xorm:"not null default 0 comment('是否允许查看角色 0-允许 1-不允许') TINYINT(1)" json:"roles_auth" form:"roles_auth" csv:"roles_auth"`
	BattleAuth    int     `xorm:"not null default 0 comment('是否允许查看战绩 0-允许 1-不允许') TINYINT(1)" json:"battle_auth" form:"battle_auth" csv:"battle_auth"`
	Stage         string  `xorm:"not null default '4,1,3,2,5,6' comment('已完成的引导阶段') VARCHAR(50)" json:"stage" form:"stage" csv:"stage"`
	JipaiqiExpire int     `xorm:"not null default 0 comment('记牌器过期时间') INT(11)" json:"jipaiqi_expire" form:"jipaiqi_expire" csv:"jipaiqi_expire"`
	Appellation   int     `xorm:"not null default 0 comment('用户称号') INT(11)" json:"appellation" form:"appellation" csv:"appellation"`
}
