package structure

type GameTask struct {
	Tkid     int    `xorm:"not null pk autoincr comment('每日任务id') INT(11)" json:"tkid" form:"tkid" csv:"tkid"`
	Type     int    `xorm:"not null default 0 comment('任务类型 1-每日任务 2-成就任务') TINYINT(1)" json:"type" form:"type" csv:"type"`
	Gptype   int    `xorm:"not null default 0 comment('分组类型 1-独立任务 2-分组任务') TINYINT(1)" json:"gptype" form:"gptype" csv:"gptype"`
	Gpid     int    `xorm:"not null default 0 comment('任务组id 0-未分组') INT(11)" json:"gpid" form:"gpid" csv:"gpid"`
	Gid      int    `xorm:"not null default 0 comment('游戏ID 0-大厅任务 1000-段位赛任务  >1000各游戏任务') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Title    string `xorm:"not null default '' comment('任务标题') VARCHAR(100)" json:"title" form:"title" csv:"title"`
	Mstype   int    `xorm:"not null default 0 comment('任务具体类型  1-累计玩X局游戏 2-累计X游戏时长(分钟) 3-单局赢X金币次数 4-地主身份胜利X次 5-农民身份胜利X次 6-连赢X局 7-累计致对手破产X次 8-累计赢X金币 9-获得X段位分 10-段位升X星 11-升X段 12-获得连胜宝箱X 13-开启连胜宝箱X 14-累计获胜X次 15-累计充值X元 16-单笔充值X元 17-获得救济X次 18-普通月卡领奖X次 19-登录游戏X次 20-游戏分享X次 21-手机绑定 22-实名认证 23-给人物送礼X次 24-人物升级 25-人物突破 26-累计打出炸弹 27-单局打出炸弹 28-获得人物 29-满级人物 30-连续登陆游戏 31-历史最高段位') TINYINT(3)" json:"mstype" form:"mstype" csv:"mstype"`
	Msnum    int    `xorm:"not null default 0 comment('任务数量') INT(10)" json:"msnum" form:"msnum" csv:"msnum"`
	Msnum2   int    `xorm:"not null default 0 comment('任务数量2') INT(10)" json:"msnum2" form:"msnum2" csv:"msnum2"`
	Liveness int    `xorm:"not null default 0 comment('达成目标获得活跃度') INT(8)" json:"liveness" form:"liveness" csv:"liveness"`
	Wppack   string `xorm:"not null default '' comment('达成目录获得物品，格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Objects  string `xorm:"not null comment('成就任务数量及奖励') TEXT" json:"objects" form:"objects" csv:"objects"`
	Desc     string `xorm:"not null default '' comment('描述') VARCHAR(100)" json:"desc" form:"desc" csv:"desc"`
}
