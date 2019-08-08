package structure

type UserGps struct {
	Uid      int     `xorm:"not null pk comment('uid') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Lng      float64 `xorm:"not null default 0.000000 comment('lng') DOUBLE(10,6)" json:"lng" form:"lng" csv:"lng"`
	Lat      float64 `xorm:"not null default 0.000000 comment('lat') DOUBLE(10,6)" json:"lat" form:"lat" csv:"lat"`
	Area     string  `xorm:"not null index VARCHAR(50)" json:"area" form:"area" csv:"area"`
	Address  string  `xorm:"not null default '' index VARCHAR(100)" json:"address" form:"address" csv:"address"`
	Ltime    int     `xorm:"not null default 0 comment('ltime') INT(11)" json:"ltime" form:"ltime" csv:"ltime"`
	Nlng     float64 `xorm:"not null default 0.000000 DOUBLE(10,6)" json:"nlng" form:"nlng" csv:"nlng"`
	Nlat     float64 `xorm:"not null default 0.000000 DOUBLE(10,6)" json:"nlat" form:"nlat" csv:"nlat"`
	Narea    string  `xorm:"not null default '' VARCHAR(50)" json:"narea" form:"narea" csv:"narea"`
	Naddress string  `xorm:"not null default '' VARCHAR(100)" json:"naddress" form:"naddress" csv:"naddress"`
	Larea    string  `xorm:"not null default '' VARCHAR(100)" json:"larea" form:"larea" csv:"larea"`
	Laddress string  `xorm:"not null default '' VARCHAR(100)" json:"laddress" form:"laddress" csv:"laddress"`
}
