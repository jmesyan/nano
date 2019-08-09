package structure

type YlyMotorPool struct {
	Uid      int    `xorm:"not null pk autoincr comment('用户id') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Nickname string `xorm:"not null default '' comment('用户昵称') VARCHAR(20)" json:"nickname" form:"nickname" csv:"nickname"`
	Avatar   string `xorm:"not null default '' comment('用户头像') VARCHAR(200)" json:"avatar" form:"avatar" csv:"avatar"`
	Use      int    `xorm:"not null default 0 TINYINT(1)" json:"use" form:"use" csv:"use"`
}
