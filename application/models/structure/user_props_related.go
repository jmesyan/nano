package structure

type UserPropsRelated struct {
	//用户props属性
	Upid      int `xorm:"not null pk default 0 comment('用户道具ID') INT(11)" json:"upid" form:"upid" csv:"upid"`
	Uid       int `xorm:"not null comment('用户id') index(uid,pid) INT(11)" json:"uid" form:"uid" csv:"uid"`
	Pid       int `xorm:"not null comment('道具id') index(uid,pid) INT(11)" json:"pid" form:"pid" csv:"pid"`
	Syxhtype  int `xorm:"not null default 0 comment('使用消耗货币类型 1-金币 2-宝石') TINYINT(3)" json:"syxhtype" form:"syxhtype" csv:"syxhtype"`
	Syxhnum   int `xorm:"not null default 0 comment('使用消耗货币数量') INT(5)" json:"syxhnum" form:"syxhnum" csv:"syxhnum"`
	Dqtype    int `xorm:"not null default 0 comment('丢弃类型 0-不可丢弃 1-可丢弃') TINYINT(2)" json:"dqtype" form:"dqtype" csv:"dqtype"`
	Sytype    int `xorm:"not null default 0 comment('使用类型 0-不可使用 1-可使用') TINYINT(2)" json:"sytype" form:"sytype" csv:"sytype"`
	Extime    int `xorm:"not null default 0 comment('过期时间 0永不过期') INT(11)" json:"extime" form:"extime" csv:"extime"`
	Ulroleid  int `xorm:"not null default 0 comment('体验人物ID') INT(11)" json:"ulroleid" form:"ulroleid" csv:"ulroleid"`
	Ulrolelv  int `xorm:"not null default 0 comment('体验人物等级') INT(8)" json:"ulrolelv" form:"ulrolelv" csv:"ulrolelv"`
	Available int `xorm:"not null default 0 comment('可用的数量') INT(8)" json:"available" form:"available" csv:"available"`
	Used      int `xorm:"not null default 0 comment('已经使用数量') INT(8)" json:"used" form:"used" csv:"used"`
	Abandond  int `xorm:"not null default 0 comment('丢弃的数量') INT(8)" json:"abandond" form:"abandond" csv:"abandond"`
	Expired   int `xorm:"not null default 0 comment('过期的数量') INT(8)" json:"expired" form:"expired" csv:"expired"`
	Ltime     int `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime     int `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
	//系统props属性
	Type      int    `xorm:"not null default 0 comment('道具类型 1-货币道具 2-经验材料 3-特效装备 4-人物体验卡 5-突破材料 6-改名卡 7-随机道具') index INT(5)" json:"type" form:"type" csv:"type"`
	Subtype   int    `xorm:"not null default 0 comment('道具子类型') TINYINT(3)" json:"subtype" form:"subtype" csv:"subtype"`
	Name      string `xorm:"not null default '' comment('道具名称') VARCHAR(100)" json:"name" form:"name" csv:"name"`
	Icon      int    `xorm:"not null default 0 comment('icon-ID') INT(5)" json:"icon" form:"icon" csv:"icon"`
	Sytime    int    `xorm:"not null default 0 comment('使用期限 0-永久 单位分钟') INT(8)" json:"sytime" form:"sytime" csv:"sytime"`
	Hbtype    int    `xorm:"not null default 0 comment('货币类型') TINYINT(3)" json:"hbtype" form:"hbtype" csv:"hbtype"`
	Quality   int    `xorm:"not null default 0 comment('品级 1-4级') TINYINT(3)" json:"quality" form:"quality" csv:"quality"`
	Exps      int    `xorm:"not null default 0 comment('增加经验值') INT(8)" json:"exps" form:"exps" csv:"exps"`
	HabitRole int    `xorm:"not null default 0 comment('偏好人物类型') INT(10)" json:"habit_role" form:"habit_role" csv:"habit_role"`
	Effect    int    `xorm:"not null default 0 comment('表现效果id') INT(11)" json:"effect" form:"effect" csv:"effect"`
	Repeatid  int    `xorm:"not null default 0 comment('重复获得替换道具ID') INT(11)" json:"repeatid" form:"repeatid" csv:"repeatid"`
	Repeatnum int    `xorm:"not null default 0 comment('重复获得替换道具数量') INT(11)" json:"repeatnum" form:"repeatnum" csv:"repeatnum"`
	Bwtype    int    `xorm:"not null default 0 comment('袋子物品类型') INT(5)" json:"bwtype" form:"bwtype" csv:"bwtype"`
	Bquality  int    `xorm:"not null comment('袋子物品品级') TINYINT(3)" json:"bquality" form:"bquality" csv:"bquality"`
	Minnum    int    `xorm:"not null default 0 comment('袋子最小物品数量') INT(5)" json:"minnum" form:"minnum" csv:"minnum"`
	Maxnum    int    `xorm:"not null default 0 comment('袋子最大物品数量') INT(5)" json:"maxnum" form:"maxnum" csv:"maxnum"`
	Overlay   int    `xorm:"not null default 0 comment('0-不支持叠加 1-支持叠加') TINYINT(1)" json:"overlay" form:"overlay" csv:"overlay"`
	Desc      string `xorm:"not null default '' comment('描述') VARCHAR(100)" json:"desc" form:"desc" csv:"desc"`
}
