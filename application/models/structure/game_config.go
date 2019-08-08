package structure

type GameConfig struct {
	Cid    int    `xorm:"not null pk autoincr comment('编号') INT(11)" json:"cid" form:"cid" csv:"cid"`
	Cvalue int    `xorm:"not null default 0 comment('值') INT(11)" json:"cvalue" form:"cvalue" csv:"cvalue"`
	Cdesc  string `xorm:"not null default '' comment('描述') VARCHAR(20)" json:"cdesc" form:"cdesc" csv:"cdesc"`
	Ctext  string `xorm:"comment('字符串值') TEXT" json:"ctext" form:"ctext" csv:"ctext"`
}
