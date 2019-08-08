package structure

type GameRatioFactor struct {
	Ratio  int `xorm:"not null pk comment('全局倍率') INT(10)" json:"ratio" form:"ratio" csv:"ratio"`
	Factor int `xorm:"not null comment('系数') INT(10)" json:"factor" form:"factor" csv:"factor"`
}
