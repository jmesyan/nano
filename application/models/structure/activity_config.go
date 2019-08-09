package structure

type ActivityConfig struct {
	Id   int     `xorm:"not null pk autoincr comment('id') INT(11)" json:"id" form:"id" csv:"id"`
	Acid int     `xorm:"not null comment('活动id') index INT(11)" json:"acid" form:"acid" csv:"acid"`
	Type int     `xorm:"not null default 0 comment('配置类型') TINYINT(1)" json:"type" form:"type" csv:"type"`
	F1   int     `xorm:"not null default 0 comment('配置字段1') INT(11)" json:"f1" form:"f1" csv:"f1"`
	F2   int     `xorm:"not null default 0 comment('配置字段2') INT(11)" json:"f2" form:"f2" csv:"f2"`
	F3   int     `xorm:"not null default 0 comment('配置字段3') INT(11)" json:"f3" form:"f3" csv:"f3"`
	F4   int     `xorm:"not null default 0 comment('配置字段4') INT(11)" json:"f4" form:"f4" csv:"f4"`
	F5   int     `xorm:"not null default 0 comment('配置字段5') INT(11)" json:"f5" form:"f5" csv:"f5"`
	F6   int     `xorm:"not null default 0 comment('配置字段6') INT(11)" json:"f6" form:"f6" csv:"f6"`
	F7   int     `xorm:"not null default 0 comment('配置字段7') INT(11)" json:"f7" form:"f7" csv:"f7"`
	F8   int     `xorm:"not null default 0 comment('配置字段8') INT(11)" json:"f8" form:"f8" csv:"f8"`
	F9   int     `xorm:"not null default 0 comment('配置字段9') INT(11)" json:"f9" form:"f9" csv:"f9"`
	F10  int     `xorm:"not null default 0 comment('配置字段10') INT(11)" json:"f10" form:"f10" csv:"f10"`
	F11  float32 `xorm:"not null default 0.00 comment('配置字段11') FLOAT(11,2)" json:"f11" form:"f11" csv:"f11"`
	F12  float32 `xorm:"not null default 0.00 comment('配置字段12') FLOAT(11,2)" json:"f12" form:"f12" csv:"f12"`
	F13  float32 `xorm:"not null default 0.00 comment('配置字段13') FLOAT(11,2)" json:"f13" form:"f13" csv:"f13"`
	F14  float32 `xorm:"not null default 0.00 comment('配置字段14') FLOAT(11,2)" json:"f14" form:"f14" csv:"f14"`
	F15  string  `xorm:"not null default '' comment('配置字段15') VARCHAR(300)" json:"f15" form:"f15" csv:"f15"`
	F16  string  `xorm:"not null default '' comment('配置字段16') VARCHAR(300)" json:"f16" form:"f16" csv:"f16"`
	F17  string  `xorm:"not null default '' comment('配置字段17') VARCHAR(300)" json:"f17" form:"f17" csv:"f17"`
	F18  string  `xorm:"not null comment('配置字段18') TEXT" json:"f18" form:"f18" csv:"f18"`
	F19  string  `xorm:"not null comment('配置字段19') TEXT" json:"f19" form:"f19" csv:"f19"`
	F20  string  `xorm:"not null comment('配置字段20') TEXT" json:"f20" form:"f20" csv:"f20"`
}
