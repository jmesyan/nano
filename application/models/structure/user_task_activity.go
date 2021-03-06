package structure

type UserTaskActivity struct {
	Uid    int    `xorm:"not null pk default 0 comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Type   int    `xorm:"not null pk default 0 comment('任务类型 1-首充 2-累积充值 3-新手7日登陆 4-每日签到 5-新手专属任务') TINYINT(3)" json:"type" form:"type" csv:"type"`
	Gid    int    `xorm:"not null pk default 0 comment('游戏ID 0-大厅任务 1000-段位赛任务  >1000各游戏任务') INT(10)" json:"gid" form:"gid" csv:"gid"`
	Mstype int    `xorm:"not null pk default 0 comment('任务具体类型  1-累计玩X局游戏 2-累计X游戏时长(分钟) 3-单局赢X金币次数 4-地主身份胜利X次 5-农民身份胜利X次 6-连赢X局 7-累计致对手破产X次 8-累计赢X金币 9-获得X段位分 10-段位升X星 11-升X段 12-获得连胜宝箱X 13-开启连胜宝箱X 14-累计获胜X次 15-累计充值X元 16-单笔充值X元 17-获得救济X次 18-普通月卡领奖X次 19-登录游戏X次 20-游戏分享X次 21-手机绑定 22-实名认证 23-给人物送礼X次 24-人物升级 25-人物突破 26-累计打出炸弹 27-单局打出炸弹 28-获得人物 29-满级人物 30-连续登陆游戏') INT(10)" json:"mstype" form:"mstype" csv:"mstype"`
	Msnum  int    `xorm:"not null pk default 0 comment('任务数量') INT(10)" json:"msnum" form:"msnum" csv:"msnum"`
	Acnum  int    `xorm:"not null default 0 comment('达成数量') INT(10)" json:"acnum" form:"acnum" csv:"acnum"`
	State  int    `xorm:"not null default 0 comment('任务状态 -1-未完成 0-完成待领取 1-完成已领取') INT(5)" json:"state" form:"state" csv:"state"`
	Wppack string `xorm:"not null default '' comment('达成目标获得物品，格式pid1-num1, pid2-num2') VARCHAR(300)" json:"wppack" form:"wppack" csv:"wppack"`
	Stones int    `xorm:"not null default 0 comment('补签耗钻') INT(10)" json:"stones" form:"stones" csv:"stones"`
	Ltime  int    `xorm:"not null default 0 comment('添加时间') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Utime  int    `xorm:"not null default 0 comment('更新时间') INT(11)" json:"utime" form:"utime" csv:"utime"`
}
