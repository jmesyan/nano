package structure

type YlyMotorGoldPool struct {
	Day      int   `xorm:"not null pk default 0 INT(11)" json:"day" form:"day" csv:"day"`
	Rtype    int   `xorm:"not null pk default 1 INT(10)" json:"rtype" form:"rtype" csv:"rtype"`
	Gameid   int   `xorm:"not null pk default 0 SMALLINT(5)" json:"gameid" form:"gameid" csv:"gameid"`
	Gold03   int64 `xorm:"default 0 BIGINT(20)" json:"gold_0_3" form:"gold_0_3" csv:"gold_0_3"`
	Gold38   int64 `xorm:"default 0 BIGINT(20)" json:"gold_3_8" form:"gold_3_8" csv:"gold_3_8"`
	Gold814  int64 `xorm:"default 0 BIGINT(20)" json:"gold_8_14" form:"gold_8_14" csv:"gold_8_14"`
	Gold1420 int64 `xorm:"default 0 BIGINT(20)" json:"gold_14_20" form:"gold_14_20" csv:"gold_14_20"`
	Gold2024 int64 `xorm:"default 0 BIGINT(20)" json:"gold_20_24" form:"gold_20_24" csv:"gold_20_24"`
}
