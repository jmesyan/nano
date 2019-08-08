package structure

type MatchSeasonAwards struct {
	Aid    int    `xorm:"not null pk autoincr comment('奖励id') INT(11)" json:"aid" form:"aid" csv:"aid"`
	Season int    `xorm:"not null default 0 comment('赛季ID') INT(11)" json:"season" form:"season" csv:"season"`
	Srank  int    `xorm:"not null default 0 comment('开始排名') INT(8)" json:"srank" form:"srank" csv:"srank"`
	Erank  int    `xorm:"not null default 0 comment('结束排名 9999-直到最后排名') INT(8)" json:"erank" form:"erank" csv:"erank"`
	Pids   string `xorm:"not null default '0' comment('获得道具ids, 格式 pid1-num1, pid2-num2') VARCHAR(100)" json:"pids" form:"pids" csv:"pids"`
}
