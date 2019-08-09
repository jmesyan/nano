package models

import (
	"fmt"
	"github.com/jmesyan/nano/application/models/structure"
	"strings"
	"time"
)

type AdminInfo struct {
	RoleId      int
	Games       map[int]bool
	Permissions map[string]map[string]bool
}

func AdminRole(uid int) int {
	user := &structure.SysAdminUser{UserId: uid, Status: 1}
	has, _ := dbr.Where("user_id=? and status=?", uid, 1).Get(user)
	if has {
		return user.RoleId
	}
	return 0
}

func Time() int {
	now := time.Now().Unix()
	return Int64Toint(now)
}

func UserFuncOps(uid int) []map[string]interface{} {
	sql := fmt.Sprintf("select func_ename,func_op from sys_admin_user as a inner join sys_role as b on a.role_id=b.role_id and b.status=1 and a.status=1 inner join sys_role_function as c on c.role_id=a.role_id inner join sys_app_function as d on d.func_id=c.func_id and d.status=1 where user_id=%d", uid)
	return DB.Select(sql)

}

func hasPerssion(op, f string) bool {
	if strings.Contains(op, f) {
		return true
	}
	return false
}

func GetAdmin(uid int) *AdminInfo {
	role_id := AdminRole(uid)
	if role_id == 0 {
		return nil
	}
	info := &AdminInfo{
		RoleId:      role_id,
		Permissions: make(map[string]map[string]bool),
	}
	userops := UserFuncOps(uid)

	for _, v := range userops {
		ename := IfcToString(v["func_ename"])
		op := IfcToString(v["func_op"])
		info.Permissions[ename] = make(map[string]bool)
		info.Permissions[ename]["view"] = hasPerssion(op, "view")
		info.Permissions[ename]["add"] = hasPerssion(op, "add")
		info.Permissions[ename]["edit"] = hasPerssion(op, "edit")
		info.Permissions[ename]["delete"] = hasPerssion(op, "delete")
	}
	user_games := UserGames(uid)
	info.Games = make(map[int]bool)
	for _, game := range user_games {
		info.Games[game.GameId] = true
	}
	if role_id == 1006 {
		info.Games[0] = true
	}
	return info
}

func DeleteAdmin(user_id int) {
	dbr.Delete(&structure.SysAdminUser{UserId: user_id})
}

func GetAllUser() []structure.YlyUser {
	var list []structure.YlyUser
	err := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"},
		"b.uid = a.uid").Where("b.ismotor = 0 and b.state != 5").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func GetUserByID(uid int) *structure.YlyUser {
	user := &structure.YlyUser{}
	ret, err := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"},
		"b.uid = a.uid").Where("a.uid = ?", uid).Get(user)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return user
}

func GetMemberByID(uid int) *structure.YlyMember {
	user := &structure.YlyMember{}
	ret, err := dbr.Where("uid=?", uid).Get(user)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return user
}

func GetUserByName(username string) *structure.YlyUser {
	user := &structure.YlyUser{}
	ret, err := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"},
		"b.uid = a.uid").Where("a.username = ?", username).Get(user)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return user
}

func GetUserByNickname(nickname string) *structure.YlyUser {
	user := &structure.YlyUser{}
	ret, err := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"},
		"b.uid = a.uid").Where("a.nickname = ?", nickname).Get(user)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return user
}

func UpdateUser(user_id int, data map[string]interface{}) bool {
	affected, err := dbr.Table("yly_member").Where("uid=?", user_id).Update(data)
	if err != nil {
		logger.Println(err.Error())
	}
	if affected > 0 {
		return true
	}
	return false
}

func DeleteUser(user_id int) {
	dbr.Delete(&structure.YlyMember{Uid: user_id})
	dbr.Delete(&structure.GameUserfield{Uid: user_id})
}

func UpdateUserField(uid int, data map[string]interface{}) bool {
	_, err := dbr.Table("game_userfield").Where("uid = ?", uid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func LogUserLogin(login *structure.LogUserLogin) int64 {
	ret, err := dbr.InsertOne(login)
	if err != nil {
		return 0
	}
	return ret
}

func UserFunction(role_id int) []map[string]ModelValue {
	sql := fmt.Sprintf("select a.func_id, b.app_id, func_name, func_ename, func_url, func_img, app_ename, app_name, app_img, func_op from sys_role_function as a inner join sys_app_function as b on a.func_id=b.func_id and b.status=1 inner join sys_app as c on c.app_id = b.app_id  and c.status=1 where role_id=%d order by app_order asc, func_order asc", role_id)
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func Roles() []structure.SysRole {
	var roles []structure.SysRole
	err := dbr.Where("status = ?", 1).Find(&roles)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return roles
}

func RolesAll() []structure.SysRole {
	var roles []structure.SysRole
	err := dbr.Find(&roles)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return roles
}

func Role(role_id int) *structure.SysRole {
	user := &structure.SysRole{RoleId: role_id}
	has, _ := dbr.Where("role_id = ?", role_id).Get(user)
	if has {
		return user
	}
	return nil
}

func DeleteApp(app_id int) {
	dbr.Delete(&structure.SysApp{AppId: app_id})
}
func UpdateApp(app_id int, app *structure.SysApp) {
	dbr.Where("app_id=?", app_id).Update(app)
}
func InsertApp(app *structure.SysApp) {
	dbr.Insert(app)
}
func App(app_id int) *structure.SysApp {
	app := &structure.SysApp{}
	_, err := dbr.Where("app_id=?", app_id).Get(app)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return app
}

func Apps2() []map[string]ModelValue {
	sql := "select app_id, app_ename, app_name, app_img, app_order, status from sys_app as a where exists(select app_id from sys_app_function where app_id=a.app_id) order by app_order asc"
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func Funcs(appid, rid int) []map[string]ModelValue {
	sql := fmt.Sprintf("select a.func_id, func_ename, func_name, func_img, func_order, status, role_id from sys_app_function as a left join sys_role_function as b on a.func_id=b.func_id and role_id=%d where app_id=%d order by func_order asc", rid, appid)
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func InsertRole(role *structure.SysRole) int {
	_, err := dbr.Insert(role)
	if err != nil {
		return 0
	}
	return role.RoleId
}

func UpdateRole(role_id int, role *structure.SysRole) bool {
	role.RoleId = role_id
	_, err := dbr.Where("role_id = ?", role_id).Update(role)
	if err != nil {
		return false
	}
	return true
}

func DeleteRoleFuncs(role_id int) bool {
	sql := fmt.Sprintf("delete from sys_role_function where role_id=%d", role_id)
	return DB.Delete(sql)
}

func InsertRoleFuncs(rolefuncs *structure.SysRoleFunction) {
	dbr.Insert(rolefuncs)
}

func DeleteRole(role_id int) bool {
	sql := fmt.Sprintf("delete from sys_role where role_id=%d", role_id)
	return DB.Delete(sql)
}

func Admins() []map[string]ModelValue {
	sql := "select a.user_id, nickname, email, role_name, reg_date, role_funcnames from sys_admin_user as a inner join sys_role as b on a.role_id = b.role_id and a.status=1 inner join yly_member as c on a.user_id=c.uid where b.role_id not in (1011,1012)order by b.role_id asc"
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func Admin(user_id int) map[string]ModelValue {
	sql := fmt.Sprintf("select a.user_id, email, nickname, a.role_id from sys_admin_user as a inner join sys_role as b on a.role_id = b.role_id and a.status=1 inner join yly_member as c on a.user_id=c.uid where a.user_id=%d", user_id)
	res := DB.SelectOne(sql)
	return MapFaceToMV(res)
}

func InsertAdmin(admin *structure.SysAdminUser) bool {
	_, err := dbr.Insert(admin)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func UpdateAdmin(user_id int, admin *structure.SysAdminUser) bool {
	_, err := dbr.Where("user_id = ?", user_id).Update(admin)
	if err != nil {
		return false
	}
	return true
}

func DeleteAdminGame(user_id int) {
	dbr.Delete(&structure.SysAdminUserGame{UserId: user_id})
}

func UserGames(uid int) []structure.SysAdminUserGame {
	var user []structure.SysAdminUserGame
	dbr.Where("user_id=?", uid).Find(&user)
	return user
}

func UserCheck(key string) *structure.YlyUser {
	var user *structure.YlyUser
	uid := StringToInt(key)
	if uid > 0 {
		user = GetUserByID(uid)
	} else {
		user = GetUserByName(key)
	}
	return user
}

func Functions() []map[string]ModelValue {
	sql := "select func_id, a.app_id, func_name, func_ename, func_url, func_img, func_order, a.status as func_status,app_name, app_ename, app_img, app_order, b.status as app_status from sys_app_function as a inner join sys_app as b on a.app_id=b.app_id order by app_order asc, func_order asc"
	res := DB.Select(sql)
	return MapArrToMV(res)
}
func Napps() []map[string]ModelValue {
	sql := "select app_id, app_ename, app_name, app_img, app_order, status from sys_app as a where not exists(select app_id from sys_app_function where app_id=a.app_id) order by app_order asc"
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func Func(func_id int) *structure.SysAppFunction {
	fun := &structure.SysAppFunction{FuncId: func_id}
	_, err := dbr.Where("func_id=?", func_id).Get(fun)
	if err != nil {
		return nil
	}
	return fun
}

func InsertFunc(funcs *structure.SysAppFunction) {
	dbr.Insert(funcs)
}

func UpdateFunc(func_id int, funcs *structure.SysAppFunction) {
	dbr.Where("func_id=?", func_id).Update(funcs)
}

func DeleteFunc(func_id int) {
	dbr.Delete(&structure.SysAppFunction{FuncId: func_id})
}

func DeleteFuncs(app_id int) {
	dbr.Delete(&structure.SysAppFunction{AppId: app_id})
}

func Apps() []structure.SysApp {
	var fun []structure.SysApp
	dbr.Find(&fun)
	return fun
}

type PageTotal struct {
	Total int `json:"total"`
}

func Users(begin, end, key string, oc int, os string, page int, pageSize int) []structure.YlyUser {
	var users []structure.YlyUser
	begintime := StrToTime(begin, "2006-01-02")
	endtime := StrToTime(end, "2006-01-02")
	sql := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"}, "a.uid=b.uid")
	if begintime > 0 {
		sql = sql.Where("a.reg_date >= ?", begintime)
	}
	if endtime > 0 {
		sql = sql.Where("a.reg_date <= ?", endtime+86400)
	}
	if len(key) > 0 {
		if StringToInt(key) > 0 {
			sql = sql.Where("a.uid = ?", StringToInt(key))
		} else if CheckEmail(key) {
			sql = sql.Where("a.email = ?", key)
		} else {
			sql = sql.Where("a.username like '%?%' or a.email like '%?%' or a.nickname like '%?%'", key, key, key)
		}
	}
	if oc == 2 {
		sql = sql.OrderBy("a.login_date " + os)
	} else if oc == 15 {
		sql = sql.OrderBy("a.login_times " + os)
	} else if oc == 16 {
		sql = sql.OrderBy("a.login_date " + os)
	} else {
		sql = sql.OrderBy("a.uid " + os)
	}

	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&users)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return users
}

func UsersCount(begin, end, key string, oc int, os string) int {
	begintime := StrToTime(begin, "2006-01-02")
	endtime := StrToTime(end, "2006-01-02")
	sql := dbr.Table("yly_member").Alias("a").Select("count(a.uid) total").Join("INNER", []string{"game_userfield", "b"}, "a.uid=b.uid")
	if begintime > 0 {
		sql = sql.Where("a.reg_date >= ?", begintime)
	}
	if endtime > 0 {
		sql = sql.Where("a.reg_date <= ?", endtime+86400)
	}
	if len(key) > 0 {
		if StringToInt(key) > 0 {
			sql = sql.Where("a.uid = ?", StringToInt(key))
		} else if CheckEmail(key) {
			sql = sql.Where("a.email = ?", key)
		} else {
			sql = sql.Where("a.username like '%?%' or a.email like '%?%' or a.nickname like '%?%'", key, key, key)
		}
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}

func InsertUser(user *structure.YlyMember) bool {
	_, err := dbr.Insert(user)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetUserCount() int64 {
	count, err := dbr.Count(&structure.YlyUser{})
	if err != nil {
		logger.Println(err.Error())
	}
	return count
}

func GetUIDByADID(adid, duid string) int {
	uid := 0
	has, _ := dbr.Table("yly_member_mobile").Where("adid=? and deviceuid=?", adid, duid).Cols("uid").Get(&uid)
	if !has || uid < 1 {
		_, err := dbr.Table("yly_member_mobile").Where("adid=?", adid).Cols("uid").Get(&uid)
		if err != nil {
			logger.Println(err.Error())
			uid = 0
		}
	}
	return uid
}

func InsertUserfield(data *structure.GameUserfield) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertUserAchievement(data *structure.UserAchievement) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

//GetUserAchievement 获取用户成就
func GetUserAchievement(uid int) *structure.UserAchievement {
	info := &structure.UserAchievement{}
	ret, err := dbr.Where("uid = ?", uid).Get(info)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return info
}

//UpdateUserAchievement 更新用户成就
func UpdateUserAchievement(uid int, data map[string]interface{}) bool {
	_, err := dbr.Table("user_achievement").Where("uid = ?", uid).Update(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func InsertMobileLogin(data *structure.LogMobileLogin) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func InsertRegisterIp(ip string) {
	day := time.Now().Format("20060102")
	sql := fmt.Sprintf("insert into game_register_ip(day,ip,count) values(%s,%s,1) ON DUPLICATE KEY update count=count+1", day, ip)
	dbr.Exec(sql)
}

func GetWeixinUserUnionID(uid int) string {
	var unionid string
	_, err := dbr.Table("yly_member_weixin").Where("uid = ?", uid).Cols("unionid").Get(&unionid)
	if err != nil {
		logger.Println(err.Error())
		return ""
	}
	return unionid
}

func UpdateMemberOpenID(appid, uid int, openid string) {
	sql := fmt.Sprintf("insert into yly_member_openid(appid,uid,openid) values(%d,%d,%s) on duplicate key update openid=%s", appid, uid, openid, openid)
	dbr.Exec(sql)
}

func GetMobileUser(uid int) *structure.YlyMemberMobile {
	mobile := &structure.YlyMemberMobile{}
	_, err := dbr.Where("uid=?", uid).Get(mobile)
	if err != nil {
		return nil
	}
	return mobile
}

func InsertMobileLogin2(uid int, ip string) {
	sql := fmt.Sprintf("call insertMobileLoginLog(%d,%s)", uid, ip)
	dbr.Exec(sql)
}

func InsertConfigLogin(data *structure.LogConfigLogin) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetUserServerIp(uid int) string {
	sql := fmt.Sprintf("call getUserServerIP(%d)", uid)
	ret := DB.SelectOne(sql)
	if ret == nil {
		return ""
	}
	res := MapFaceToMV(ret)
	return res["ip"].Strval
}

func LogUserDayLogin(uid, day int) {
	sql := fmt.Sprintf("insert ignore into log_day_login set uid = %d, day = %d", uid, day)
	dbr.Exec(sql)
}

//模拟登录
func GetUserFakeList() []structure.YlyMemberFake {
	var list []structure.YlyMemberFake
	err := dbr.Asc("uid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertUserFake(data *structure.YlyMemberFake) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Uid
}
func UpdateUserFake(id int, data *structure.YlyMemberFake) bool {
	data.Uid = id
	_, err := dbr.Where("uid = ?", id).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteUserFake(uid int, fakeuid int) {
	dbr.Delete(&structure.YlyMemberFake{Uid: uid, Fakeuid: fakeuid})
}
func GetOnlineUsers() []structure.YlyOnline {
	var data []structure.YlyOnline
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}

//GetWeixinUserInvitor 获取微信邀请人
func GetWeixinUserInvitor(uid int) int {
	invitor := 0
	_, err := dbr.Table("yly_member_weixin").Where("uid=?", uid).Select("invitor").Cols("invitor").Get(&invitor)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return invitor
}

//InsertRecurringUser 插入回流用户
func InsertRecurringUser(data *structure.LogRecurringUser) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Rid
}

func GetMemberMobile(uid int, deviceuid string) *structure.YlyMemberMobile {
	data := &structure.YlyMemberMobile{}
	_, err := dbr.Where("uid=? and deviceuid=?", uid, deviceuid).Get(data)
	if err != nil {
		logger.Println(err)
		return data
	}
	return data
}

func UpdateMemberMobile(uid int, deviceuid string, data map[string]interface{}) bool {
	_, err := dbr.Table("yly_member_mobile").Where("uid=? and deviceuid=?", uid, deviceuid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func InsertMemeberMobile(data *structure.YlyMemberMobile) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetTestUsers(offset, pageSize int) []structure.YlyUser {
	var users []structure.YlyUser
	sql := dbr.Table("yly_member").Alias("a").Select("a.*, b.*").Join("INNER", []string{"game_userfield", "b"}, "a.uid=b.uid").OrderBy("a.uid asc")
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&users)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return users
}
