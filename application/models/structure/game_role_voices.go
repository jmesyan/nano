package structure

type GameRoleVoices struct {
	Rvid       int    `xorm:"not null pk autoincr comment('自增id') INT(11)" json:"rvid" form:"rvid" csv:"rvid"`
	Roid       int    `xorm:"not null default 0 comment('角色id') INT(11)" json:"roid" form:"roid" csv:"roid"`
	Text       string `xorm:"not null default '' comment('语音文字描述') VARCHAR(300)" json:"text" form:"text" csv:"text"`
	Context    string `xorm:"not null comment('文本内容') VARCHAR(300)" json:"context" form:"context" csv:"context"`
	ResVoice   int    `xorm:"not null default 0 comment('语音资源ID') INT(10)" json:"res_voice" form:"res_voice" csv:"res_voice"`
	Scenes     string `xorm:"not null default '0' comment('使用场景，格式 scen1,scen2，使用逗号分隔') VARCHAR(100)" json:"scenes" form:"scenes" csv:"scenes"`
	SkillLevel int    `xorm:"default 0 comment('解锁技能等级') INT(11)" json:"skill_level" form:"skill_level" csv:"skill_level"`
}
