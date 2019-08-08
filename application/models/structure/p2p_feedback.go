package structure

import (
	"time"
)

type P2pFeedback struct {
	Id           int       `xorm:"not null pk autoincr INT(10)" json:"id" form:"id" csv:"id"`
	Uid          int       `xorm:"not null default 0 comment('用户ID') INT(10)" json:"uid" form:"uid" csv:"uid"`
	Ip           string    `xorm:"not null default '' comment('ip') VARCHAR(15)" json:"ip" form:"ip" csv:"ip"`
	Address      string    `xorm:"not null default '' comment('地址') VARCHAR(50)" json:"address" form:"address" csv:"address"`
	Type         int       `xorm:"not null default 0 comment('类型 0 系统自动反馈 1 用户建议 2 游戏bug') TINYINT(1)" json:"type" form:"type" csv:"type"`
	Text         string    `xorm:"not null comment('反馈内容') TEXT" json:"text" form:"text" csv:"text"`
	Timeline     int       `xorm:"not null default 0 comment('时间') INT(10)" json:"timeline" form:"timeline" csv:"timeline"`
	Pic          string    `xorm:"not null default '' comment('图片') VARCHAR(100)" json:"pic" form:"pic" csv:"pic"`
	UserAnswers  int       `xorm:"not null default 0 comment('用户回复数') INT(10)" json:"user_answers" form:"user_answers" csv:"user_answers"`
	GmAnswers    int       `xorm:"not null default 0 comment('客服回复数') INT(10)" json:"gm_answers" form:"gm_answers" csv:"gm_answers"`
	State        int       `xorm:"not null default 1 comment('问题状态 1 未解决 2 已解决') TINYINT(1)" json:"state" form:"state" csv:"state"`
	CompleteTime time.Time `xorm:"not null comment('完成时间') DATETIME" json:"complete_time" form:"complete_time" csv:"complete_time"`
	GmId         int       `xorm:"not null default 0 comment('客服ID') INT(10)" json:"gm_id" form:"gm_id" csv:"gm_id"`
	Score        int       `xorm:"not null default 0 comment('评分') TINYINT(1)" json:"score" form:"score" csv:"score"`
	Ver          string    `xorm:"not null default '' VARCHAR(20)" json:"ver" form:"ver" csv:"ver"`
	Md5          string    `xorm:"not null default '' index CHAR(32)" json:"md5" form:"md5" csv:"md5"`
	Times        int       `xorm:"not null default 0 INT(10)" json:"times" form:"times" csv:"times"`
	LastTime     int       `xorm:"not null default 0 INT(10)" json:"last_time" form:"last_time" csv:"last_time"`
}
