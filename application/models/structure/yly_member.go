package structure

type YlyMember struct {
	Uid        int    `xorm:"not null pk autoincr comment('用户编号') INT(11)" json:"uid" form:"uid" csv:"uid"`
	Pid        string `xorm:"not null default '' comment('平台id') CHAR(100)" json:"pid" form:"pid" csv:"pid"`
	Username   string `xorm:"not null default '' comment('用户名') unique VARCHAR(64)" json:"username" form:"username" csv:"username"`
	Password   string `xorm:"not null default '' comment('密码') CHAR(32)" json:"password" form:"password" csv:"password"`
	Tel        string `xorm:"not null default '' comment('手机') VARCHAR(20)" json:"tel" form:"tel" csv:"tel"`
	Email      string `xorm:"not null default '' comment('邮箱') CHAR(32)" json:"email" form:"email" csv:"email"`
	Birthday   string `xorm:"not null default '' comment('生日') VARCHAR(50)" json:"birthday" form:"birthday" csv:"birthday"`
	RegIp      string `xorm:"not null default '' comment('注册ip') CHAR(15)" json:"reg_ip" form:"reg_ip" csv:"reg_ip"`
	RegDate    int    `xorm:"not null default 0 comment('注册时间') index INT(11)" json:"reg_date" form:"reg_date" csv:"reg_date"`
	Gender     string `xorm:"not null default 'm' comment('性别') ENUM('f','m')" json:"gender" form:"gender" csv:"gender"`
	Utype      string `xorm:"not null default '' comment('用户类型') CHAR(10)" json:"utype" form:"utype" csv:"utype"`
	Nickname   string `xorm:"not null default '' comment('昵称') CHAR(20)" json:"nickname" form:"nickname" csv:"nickname"`
	GroupId    int    `xorm:"not null default 5 comment('管理级别') TINYINT(3)" json:"group_id" form:"group_id" csv:"group_id"`
	Locale     string `xorm:"not null comment('语言') VARCHAR(20)" json:"locale" form:"locale" csv:"locale"`
	Avatar     string `xorm:"not null comment('头像') VARCHAR(255)" json:"avatar" form:"avatar" csv:"avatar"`
	Upuid      int    `xorm:"not null default 0 comment('上层uid，是哪个用户推荐过来的') INT(11)" json:"upuid" form:"upuid" csv:"upuid"`
	Ad         string `xorm:"not null default '' comment('用户来源，哪个广告') index CHAR(100)" json:"ad" form:"ad" csv:"ad"`
	LoginIp    string `xorm:"not null default '' comment('登录ip') CHAR(15)" json:"login_ip" form:"login_ip" csv:"login_ip"`
	LoginTimes int    `xorm:"not null default 0 comment('登录次数') INT(11)" json:"login_times" form:"login_times" csv:"login_times"`
	LoginDate  int    `xorm:"not null default 0 comment('最后一次登陆时间') index INT(10)" json:"login_date" form:"login_date" csv:"login_date"`
	LoginVer   string `xorm:"comment('最后一次登录版本') VARCHAR(10)" json:"login_ver" form:"login_ver" csv:"login_ver"`
	ForbidTime int    `xorm:"not null default 0 comment('封号时间') INT(10)" json:"forbid_time" form:"forbid_time" csv:"forbid_time"`
	Remark     string `xorm:"not null default '' comment('封号备注') VARCHAR(255)" json:"remark" form:"remark" csv:"remark"`
	Sign       string `xorm:"not null default '' comment('签名') VARCHAR(255)" json:"sign" form:"sign" csv:"sign"`
	Regcity    string `xorm:"not null default '' comment('城市') VARCHAR(45)" json:"regcity" form:"regcity" csv:"regcity"`
	Regarea    string `xorm:"not null default '' comment('区域') VARCHAR(100)" json:"regarea" form:"regarea" csv:"regarea"`
	NewUser    int    `xorm:"not null default 1 comment('新用户') TINYINT(1)" json:"new_user" form:"new_user" csv:"new_user"`
	Ps         string `xorm:"comment('pass') VARCHAR(100)" json:"ps" form:"ps" csv:"ps"`
	UserSalt   string `xorm:"not null default '' comment('salt') CHAR(6)" json:"user_salt" form:"user_salt" csv:"user_salt"`
	Appid      int    `xorm:"not null default 0 comment('appid') TINYINT(4)" json:"appid" form:"appid" csv:"appid"`
	Logincity  string `xorm:"not null comment('登录地') VARCHAR(50)" json:"logincity" form:"logincity" csv:"logincity"`
	Subutype   string `xorm:"not null default '' VARCHAR(50)" json:"subutype" form:"subutype" csv:"subutype"`
}
