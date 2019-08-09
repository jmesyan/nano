package structure

type GameProps struct {
	Pid         int    `xorm:"not null pk autoincr comment('道具id') INT(11)" json:"pid" form:"pid" csv:"pid"`
	Type        int    `xorm:"not null default 0 comment('道具类型 1-货币道具 2-经验材料 3-特效装备 4-人物体验卡 5-突破材料 6-改名卡 7-随机道具 8-表情 9-bgm 10-头像框 11-人物角色卡 12-头像 13-游戏道具 14-用户称号') index INT(5)" json:"type" form:"type" csv:"type"`
	Subtype     int    `xorm:"not null default 0 comment('道具子类型 0-通用 1-出牌特效 2-结算特效 3-终结特效 4-牌背 5-桌布 6-保连胜道具 7-单抽劵道具 8-加倍卡 9-记牌器') TINYINT(3)" json:"subtype" form:"subtype" csv:"subtype"`
	Name        string `xorm:"not null default '' comment('道具名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Icon        int    `xorm:"not null default 0 comment('icon-ID') INT(5)" json:"icon" form:"icon" csv:"icon"`
	Syxhtype    int    `xorm:"not null default 0 comment('使用消耗货币类型 1-金币 2-宝石') TINYINT(3)" json:"syxhtype" form:"syxhtype" csv:"syxhtype"`
	Syxhnum     int    `xorm:"not null default 0 comment('使用消耗货币数量') INT(5)" json:"syxhnum" form:"syxhnum" csv:"syxhnum"`
	Dqtype      int    `xorm:"not null default 0 comment('丢弃类型 0-不可丢弃 1-可丢弃') TINYINT(2)" json:"dqtype" form:"dqtype" csv:"dqtype"`
	Sytype      int    `xorm:"not null default 0 comment('使用类型 0-不可使用 1-可使用') TINYINT(2)" json:"sytype" form:"sytype" csv:"sytype"`
	Sytime      int    `xorm:"not null default 0 comment('使用期限 0-永久 单位分钟') INT(8)" json:"sytime" form:"sytime" csv:"sytime"`
	Hbtype      int    `xorm:"not null default 0 comment('货币类型') TINYINT(3)" json:"hbtype" form:"hbtype" csv:"hbtype"`
	Quality     int    `xorm:"not null default 0 comment('品级 1-4级') TINYINT(3)" json:"quality" form:"quality" csv:"quality"`
	Exps        int    `xorm:"not null default 0 comment('增加经验值') INT(8)" json:"exps" form:"exps" csv:"exps"`
	HabitRole   int    `xorm:"not null default 0 comment('偏好人物类型') INT(10)" json:"habit_role" form:"habit_role" csv:"habit_role"`
	Effect      int    `xorm:"not null default 0 comment('表现效果id') INT(11)" json:"effect" form:"effect" csv:"effect"`
	Repeatid    int    `xorm:"not null default 0 comment('重复获得替换道具ID') INT(11)" json:"repeatid" form:"repeatid" csv:"repeatid"`
	Repeatnum   int    `xorm:"not null default 0 comment('重复获得替换道具数量') INT(11)" json:"repeatnum" form:"repeatnum" csv:"repeatnum"`
	Ulroleid    int    `xorm:"not null default 0 comment('人物ID') INT(11)" json:"ulroleid" form:"ulroleid" csv:"ulroleid"`
	Ulrolelv    int    `xorm:"not null default 0 comment('人物等级') INT(8)" json:"ulrolelv" form:"ulrolelv" csv:"ulrolelv"`
	Bwtype      int    `xorm:"not null default 0 comment('袋子物品类型') INT(5)" json:"bwtype" form:"bwtype" csv:"bwtype"`
	Bquality    int    `xorm:"not null comment('袋子物品品级') TINYINT(3)" json:"bquality" form:"bquality" csv:"bquality"`
	Minnum      int    `xorm:"not null default 0 comment('袋子最小物品数量') INT(5)" json:"minnum" form:"minnum" csv:"minnum"`
	Maxnum      int    `xorm:"not null default 0 comment('袋子最大物品数量') INT(5)" json:"maxnum" form:"maxnum" csv:"maxnum"`
	Overlay     int    `xorm:"not null default 0 comment('0-不支持叠加 1-支持叠加') TINYINT(1)" json:"overlay" form:"overlay" csv:"overlay"`
	Preferences string `xorm:"not null default '' comment('偏好') VARCHAR(200)" json:"preferences" form:"preferences" csv:"preferences"`
	Desc        string `xorm:"not null default '' comment('描述') VARCHAR(100)" json:"desc" form:"desc" csv:"desc"`
	Ltime       int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime       int    `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
