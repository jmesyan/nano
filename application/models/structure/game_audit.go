package structure

type GameAudit struct {
	Platform string `xorm:"not null pk default '' comment('平台') VARCHAR(20)" json:"platform" form:"platform" csv:"platform"`
	Version  string `xorm:"not null pk default '' comment('版本') VARCHAR(10)" json:"version" form:"version" csv:"version"`
}
