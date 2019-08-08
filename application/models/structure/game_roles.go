package structure

type GameRoles struct {
	Roid       int    `xorm:"not null pk autoincr comment('角色ID') INT(11)" json:"roid" form:"roid" csv:"roid"`
	Name       string `xorm:"not null default '' comment('角色名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Icon       int    `xorm:"not null default 0 comment('Icon-id') INT(5)" json:"icon" form:"icon" csv:"icon"`
	Quality    int    `xorm:"not null default 0 comment('品阶') TINYINT(3)" json:"quality" form:"quality" csv:"quality"`
	BsYz       int    `xorm:"not null default 0 comment('牌型基础分-压制') TINYINT(3)" json:"bs_yz" form:"bs_yz" csv:"bs_yz"`
	BsJb       int    `xorm:"not null default 0 comment('牌型基础分-加倍') TINYINT(3)" json:"bs_jb" form:"bs_jb" csv:"bs_jb"`
	BsLd       int    `xorm:"not null default 0 comment('牌型基础分-连对') TINYINT(3)" json:"bs_ld" form:"bs_ld" csv:"bs_ld"`
	BsFj       int    `xorm:"not null default 0 comment('牌型基础分-飞机') TINYINT(3)" json:"bs_fj" form:"bs_fj" csv:"bs_fj"`
	BsZd       int    `xorm:"not null default 0 comment('牌型基础分-炸弹') TINYINT(3)" json:"bs_zd" form:"bs_zd" csv:"bs_zd"`
	BsSz       int    `xorm:"not null default 0 comment('牌型基础分-顺子') TINYINT(3)" json:"bs_sz" form:"bs_sz" csv:"bs_sz"`
	BsSd       int    `xorm:"not null default 0 comment('牌型基础分-三带') TINYINT(3)" json:"bs_sd" form:"bs_sd" csv:"bs_sd"`
	Skid       int    `xorm:"not null default 0 comment('技能ID') INT(11)" json:"skid" form:"skid" csv:"skid"`
	Maxlevel   int    `xorm:"not null default 0 comment('角色上限') INT(5)" json:"maxlevel" form:"maxlevel" csv:"maxlevel"`
	HabitWp    int    `xorm:"not null default 0 comment('偏好物品ID') INT(11)" json:"habit_wp" form:"habit_wp" csv:"habit_wp"`
	HabitRole  int    `xorm:"not null default 0 comment('偏好人物类型') INT(10)" json:"habit_role" form:"habit_role" csv:"habit_role"`
	ResArt     int    `xorm:"not null default 0 comment('美术路径-id') INT(5)" json:"res_art" form:"res_art" csv:"res_art"`
	ResLv      string `xorm:"not null default '' comment('技能突破解锁表情IDS，格式 技能等级1-表情ID1,技能等级2-表情ID2') VARCHAR(300)" json:"res_lv" form:"res_lv" csv:"res_lv"`
	ResVoice   int    `xorm:"not null default 0 comment('语音路径-id') INT(5)" json:"res_voice" form:"res_voice" csv:"res_voice"`
	ResEmoijs  string `xorm:"not null default '' comment('默认表情道具ID') VARCHAR(300)" json:"res_emoijs" form:"res_emoijs" csv:"res_emoijs"`
	ResVemoijs string `xorm:"not null default '' comment('解锁表情道具ID') VARCHAR(300)" json:"res_vemoijs" form:"res_vemoijs" csv:"res_vemoijs"`
	Actors     string `xorm:"not null default '' comment('配音人员') VARCHAR(100)" json:"actors" form:"actors" csv:"actors"`
	Chathead   int    `xorm:"not null default 0 comment('头像道具ID') INT(11)" json:"chathead" form:"chathead" csv:"chathead"`
	Chatavator int    `xorm:"not null default 0 comment('头像框道具ID') INT(11)" json:"chatavator" form:"chatavator" csv:"chatavator"`
	Taste      int    `xorm:"not null default 1 comment('0正常角色 1体验角色') TINYINT(3)" json:"taste" form:"taste" csv:"taste"`
	Isdefault  int    `xorm:"not null default 0 comment('是否默认角色 0-否 1-是') TINYINT(1)" json:"isdefault" form:"isdefault" csv:"isdefault"`
	Desc       string `xorm:"not null default '' comment('描述') VARCHAR(500)" json:"desc" form:"desc" csv:"desc"`
	Ltime      int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime      int    `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
