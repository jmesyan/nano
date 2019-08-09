package models

import (
	"fmt"
	"github.com/jmesyan/nano/application/models/structure"
	"strings"
)

func GetUserProps(uid int, args ...int) []structure.UserPropsRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	props := []structure.UserPropsRelated{}
	sql := dbr.Table("user_props").Alias("a").Select("a.*, b.type,b.subtype,b.name,b.icon,b.sytime,b.hbtype,b.quality,b.exps,b.habit_role,b.effect,b.repeatid,b.bwtype,b.bquality,b.minnum,b.maxnum,b.overlay,b.desc,b.preferences").Join("INNER", []string{"game_props", "b"}, "a.pid=b.pid").Where("a.uid=?", uid)
	if state > 0 {
		switch state {
		case 1:
			//获取使用列表
			sql = sql.Where("a.used > ?", 0)
		case 2:
			//获取丢弃列表
			sql = sql.Where("a.abandond > ?", 0)
		case 3:
			//获取过期列表
			sql = sql.Where("a.expired > ?", 0)
		case 4:
			//获取有过期时间并且还有没使用的道具
			sql = sql.Where("a.extime > ? and a.available > ?", 0, 0)
		case 5:
			//获取表情列表
			sql = sql.Where("a.available > 0 and b.type = 8")
		case 6:
			//获取突破材料列表
			sql = sql.Where("a.available > 0 and b.type = 5")
		case 7:
			//获取头像框
			sql = sql.Where("a.available > 0 and b.type = 10")
		case 8:
			//获取头像
			sql = sql.Where("a.available > 0 and b.type = 12")
		case 9:
			//获取复活币
			sql = sql.Where("a.available > 0 and b.type = 13 and b.subtype=6")
		default:
			//获取可用列表
			sql = sql.Where("a.available > ?", 0)
		}
	} else {
		//获取可用列表
		sql = sql.Where("a.available > ?", 0)
	}
	err := sql.Find(&props)
	if err != nil {
		logger.Println(err.Error())
		return props
	}
	return props
}

func GetUserPropsNum(uid int) int {
	var total int
	_, err := dbr.Table("user_props").Alias("a").Select("count(a.upid) total").Join("INNER",
		[]string{"game_props", "b"}, "a.pid=b.pid").Where("a.uid=? and a.available>0 and b.type not in(1,3,4,8,9,10)",
		uid).Cols("cap").Get(&total)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return total
}

func GetUserRepeatRrops(uid int) []structure.UserPropsRelated {
	props := []structure.UserPropsRelated{}
	sql := dbr.Table("user_props").Alias("a").Select("a.*, b.type,b.subtype,b.name,b.icon,b.sytime,b.hbtype,b.quality,b.exps,b.habit_role,b.effect,b.repeatid,b.repeatnum,b.bwtype,b.bquality,b.minnum,b.maxnum,b.overlay,b.desc").Join("INNER", []string{"game_props", "b"}, "a.pid=b.pid").Where("a.uid=?", uid)
	sql = sql.Where("a.available > 0 and b.overlay=0 and b.repeatid>0 and b.repeatnum > 0")
	err := sql.Find(&props)
	if err != nil {
		logger.Println(err.Error())
		return props
	}
	return props
}

func UpdateUserProps(upid int, data *structure.UserProps) bool {
	data.Upid = upid
	_, err := dbr.Where("upid = ?", upid).MustCols("available").Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func DeleteUserProps(upid int) {
	dbr.Delete(&structure.UserProps{Upid: upid})
}

func InsertUserPropsLog(data *structure.LogUserProps) {
	dbr.Insert(data)
}

func GetUserRoles(uid int, args ...int) []structure.UserRolesRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	roles := []structure.UserRolesRelated{}
	sql := dbr.Table("user_roles").Alias("a").Select("a.*, b.name,b.icon,b.quality,b.bs_yz,b.bs_jb,b.bs_ld,b.bs_fj,b.bs_zd,b.bs_sz,b.bs_sd,b.skid,b.maxlevel,b.habit_wp,b.habit_role, b.res_art,b.res_voice,b.res_emoijs,b.res_vemoijs,b.res_lv,b.actors, b.desc, b.chathead").Join("INNER", []string{"game_roles", "b"}, "a.roid=b.roid").Where("a.uid=?", uid)
	if state > 0 {
		switch state {
		case 1:
			//获取体验角色
			sql = sql.Where("a.extime > ?", 0)
		case 2:
			//获取默认角色
			sql = sql.Where("a.isdefault = ?", 1)
		case 3:
			//获取出战角色
			sql = sql.Where("a.isfight = ?", 1)
		}
	}
	err := sql.Find(&roles)
	if err != nil {
		logger.Println(err.Error())
		return roles
	}
	return roles
}

func InsertUserRoles(data *structure.UserRoles) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func DeleteUserRoles(uroid int) {
	dbr.Delete(&structure.UserRoles{Uroid: uroid})
}

func UpdateUserRoles(uroid int, data map[string]interface{}) bool {
	_, err := dbr.Table("user_roles").Where("uroid = ?", uroid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetUserBoxes(uid int, args ...int) []structure.UserBoxesRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	props := []structure.UserBoxesRelated{}
	sql := dbr.Table("user_boxes").Alias("a").Select("a.*,b.name,b.btype,b.btime,b.bnum,b.quality").Join("INNER", []string{"game_boxes", "b"}, "a.bid=b.bid").Where("a.uid=?", uid)
	if state > 0 {
		switch state {
		case 1:
			//获取未开启列表
			sql = sql.Where("a.state = ?", 0)
		case 2:
			//获取临时列表
			sql = sql.Where("a.state = ?", 2)
		default:
			//获取可用列表
			sql = sql.Where("a.state = ? or a.state=? or a.state=?", 0, 2, 3)
		}
	} else {
		//获取可用列表
		sql = sql.Where("a.state = ? or a.state=? or a.state=?", 0, 2, 3)
	}
	err := sql.Asc("ltime").Find(&props)
	if err != nil {
		logger.Println(err.Error())
		return props
	}
	return props
}

func DeleteUserBox(ubid int) {
	dbr.Delete(&structure.UserBoxes{Ubid: ubid})
}

func UpdateUserBoxes(ubid int, data map[string]interface{}) bool {
	_, err := dbr.Table("user_boxes").Where("ubid = ?", ubid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetFeedback(ftype int, stime int, etime int, page int, pageSize int) []structure.P2pFeedback {
	var list []structure.P2pFeedback
	sql := dbr.Table("p2p_feedback").Where("last_time>?", stime).Where("last_time<?", etime)
	if ftype > -1 {
		sql = sql.Where("type=?", ftype)
	}
	sql = sql.OrderBy("id desc")
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetFeedbackCount(ftype int, stime int, etime int) int {
	sql := dbr.Table("p2p_feedback").Select("count(1) total").Where("last_time>?", stime).Where("last_time<?", etime)
	if ftype > -1 {
		sql = sql.Where("type=?", ftype)
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}
func UpdateFeedback(id int) {
	data := &structure.P2pFeedback{State: 2}
	dbr.Table("p2p_feedback").Where("id=?", id).Update(data)
}

func GetForbidDeviceuid() []structure.YlyForbidDeviceuid {
	var data []structure.YlyForbidDeviceuid
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func InsertForbidDeviceuid(data *structure.YlyForbidDeviceuid) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteForbidDeviceuid(id int) {
	dbr.Delete(&structure.YlyForbidDeviceuid{Id: id})
}

func GetGameBlockIP() []structure.GameBlockIp {
	var data []structure.GameBlockIp
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func GetCode(uid int, nickname string, stime int64, etime int64, gid int) []int {
	tables := "log_create_user_" + string(gid)
	sql := fmt.Sprintf("select a.nroomid from %s as a inner join yly_member as b on a.uid = b.uid where a.ltime >= %v and a.ltime <= %v and ", tables, stime, etime)
	var con string
	if uid > 0 {
		sql = sql + fmt.Sprintf(" (a.uid = %v ", uid)
		con = " or "
	}
	if nickname != "" {
		if con != "" {
			sql = sql + fmt.Sprintf(" %v b.nickname like '%%v%%'", con, nickname)
		} else {
			sql = sql + fmt.Sprintf(" (b.nickname like '%%v%%'", nickname)
		}
		con = " or "
	}
	if con != "" {
		sql = sql + " )"
	}
	res := DB.Select(sql)
	data := MapArrToMV(res)
	arr := []int{}
	for _, info := range data {
		arr = append(arr, info["nroomid"].Intval)
	}
	return arr

}
func GetDayRound(stime int64, etime int64, code []int, mytype int, start int, limit int, ridx int, gidStr string, gidInt int) []map[string]ModelValue {
	var table string
	if gidInt < 10000 {
		table = "log_create_user_" + gidStr
	} else {
		table = "log_create_user_match"
	}
	sql := fmt.Sprintf("select a.lid,a.nroomid,GROUP_CONCAT(a.uid) as uids,b.ltime,b.uid as owner,a.cur,a.round,a.is_over,b.state,b.gsid,b.tid, b.cid, b.feeuid, b.buytype,b.type btype,b.endtime from %s as a left join log_create_rooms as b on a.lid = b.lid  where  a.ltime >= %v and a.ltime < %v ", table, stime, etime)

	if len(code) > 0 {
		instr := strings.Replace(strings.Trim(fmt.Sprint(code), "[]"), " ", ",", -1)
		sql = sql + fmt.Sprintf(" and b.code in (%s) ", instr)
	}
	switch mytype {
	case 1:
		sql += " and a.cur > 1 "
		break
	case 2:
		sql += " and a.is_over = 0 "
		break
	case 3:
		sql += " and a.is_over = 1 "
		break
	}
	if ridx > 0 {
		sql += fmt.Sprintf("and b.gsid = '%v_1_%v' ", gidInt, ridx)
	}
	sql += fmt.Sprintf("group by a.lid order by a.ltime desc limit %v,%v ", start, limit)
	res := DB.Select(sql)
	data := MapArrToMV(res)
	return data
}
func GetDayRoundCount(stime int64, etime int64, code []int, mytype int, ridx int, gidStr string, gidInt int) int {
	var table string
	if gidInt < 10000 {
		table = "log_create_user_" + gidStr
	} else {
		table = "log_create_user_match"
	}
	sql := fmt.Sprintf("select count(1) as acount from %s as a left join log_create_rooms as b on a.lid = b.lid where  a.ltime >= %v and a.ltime < %v ", table, stime, etime)

	if len(code) > 0 {
		instr := strings.Replace(strings.Trim(fmt.Sprint(code), "[]"), " ", ",", -1)
		sql = sql + fmt.Sprintf(" and b.code in (%s) ", instr)
	}
	switch mytype {
	case 1:
		sql += " and a.cur > 1 "
		break
	case 2:
		sql += " and a.is_over = 0 "
		break
	case 3:
		sql += " and a.is_over = 1 "
		break
	}
	if ridx > 0 {
		sql += fmt.Sprintf("and b.gsid = '%v_1_%v' ", gidInt, ridx)
	}
	sql += " group by a.lid "
	res := DB.SelectOne(sql)
	data := MapFaceToMV(res)
	return data["acount"].Intval
}

func InsertGameBlockIP(data *structure.GameBlockIp) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func UpdateGameBlockIP(ip string, data *structure.GameBlockIp) bool {
	_, err := dbr.Table("game_block_ip").Where("ip = ?", ip).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameBlockIP(ip string) {
	dbr.Delete(&structure.GameBlockIp{Ip: ip})
}

func GetDumbs(begin int, end int, uid int, state int, oc int, os string, page int, pageSize int) []structure.UsersRelated {
	var list []structure.UsersRelated
	sql := dbr.Table("game_userfield").Alias("u").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid").Where("u.state!=0")
	if begin > 0 {
		sql = sql.Where("u.startdumb >= ?", begin)
	}
	if end > 0 {
		sql = sql.Where("u.dumblimit <= ?", end)
	}
	if uid > 0 {
		sql = sql.Where("u.uid = ?", uid)
	}
	if state > 0 {
		sql = sql.Where("u.state = ?", state)
	}
	if oc == 2 {
		sql = sql.OrderBy("m.login_date " + os)
	} else {
		sql = sql.OrderBy("u.uid " + os)
	}
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetDumbCount(begin int, end int, uid int, state int) int {
	sql := dbr.Table("game_userfield").Alias("u").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid").Select("count(1) total").Where("u.state!=0")
	if begin > 0 {
		sql = sql.Where("u.startdumb >= ?", begin)
	}
	if end > 0 {
		sql = sql.Where("u.dumblimit <= ?", end)
	}
	if uid > 0 {
		sql = sql.Where("u.uid = ?", uid)
	}
	if state > 0 {
		sql = sql.Where("u.state = ?", state)
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}
func UpdateDumb(uid int, data map[string]interface{}, state int) bool {
	sql := dbr.Table("game_userfield").Where("uid = ?", uid)
	if state >= 0 {
		sql = sql.Where("state = ?", state)
	}
	_, err := sql.Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func GetDumbAll(state int) []structure.GameUserfield {
	var list []structure.GameUserfield
	sql := dbr.Table("game_userfield").Alias("u").Select("u.uid uid").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid")
	if state < 0 {
		sql = sql.Where("state != 0")
	} else {
		sql = sql.Where("state = ?", state)
	}
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

//计算用户抽卡宝箱已经抽卡次数
func GetUserLotteryBoxTimes(uid, bid int, begin, end int64) int {
	lottery := structure.LogLotteryCards{}
	counts, err := dbr.Where("uid=? and bid=? and ltime >= ? and ltime < ?", uid, bid, begin, end).Count(&lottery)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return Int64Toint(counts)
}

func InsertUserLotteryBox(data *structure.LogLotteryCards) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func InsertUserBox(data *structure.UserBoxes) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetForbids(begin int, end int, key string, oc int, os string, page int, pageSize int, is_forbid int) []structure.UsersRelated {
	var list []structure.UsersRelated
	sql := dbr.Table("game_userfield").Alias("u").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid")
	if is_forbid == 1 {
		sql = sql.Where("(m.group_id=13 or u.state=2)")
	}
	if begin > 0 {
		sql = sql.Where("u.forbid_time >= ?", begin)
	}
	if end > 0 {
		sql = sql.Where("u.forbid_time <= ?", end)
	}
	if len(key) > 0 {
		if StringToInt(key) > 0 {
			sql = sql.Where("m.uid = ?", StringToInt(key))
		} else if CheckEmail(key) {
			sql = sql.Where("m.email = ?", key)
		} else if CheckIP(key) {
			sql = sql.Where("(m.login_ip = ? or m.reg_ip = ?)", key, key)
		} else if CheckCN(key) {
			sql = sql.Where("(m.nickname like '%?%' or m.logincity like '%?%')", key, key)
		} else {
			sql = sql.Where("(m.username like '%?%' or m.nickname like '%?%')", key, key)
		}
	}
	if oc == 2 {
		sql = sql.OrderBy("m.login_date " + os)
	} else if oc == 8 {
		sql = sql.OrderBy("u.pay " + os)
	} else if oc == 9 {
		sql = sql.OrderBy("u.pay_num " + os)
	} else if oc == 12 {
		sql = sql.OrderBy("u.room_card " + os)
	} else if oc == 17 {
		sql = sql.OrderBy("u.ticket " + os)
	} else if is_forbid > 0 {
		sql = sql.OrderBy("u.forbid_time desc")
	} else {
		sql = sql.OrderBy("u.uid " + os)
	}
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetForbidCount(begin int, end int, key string, is_forbid int) int {
	sql := dbr.Table("game_userfield").Alias("u").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid").Select("count(1) total").Where("m.uid>0")
	if is_forbid == 1 {
		sql = sql.Where("(m.group_id=13 or u.state=2)")
	}
	if begin > 0 {
		sql = sql.Where("u.reg_date >= ?", begin)
	}
	if end > 0 {
		sql = sql.Where("u.reg_date <= ?", end)
	}
	if len(key) > 0 {
		if StringToInt(key) > 0 {
			sql = sql.Where("m.uid = ?", StringToInt(key))
		} else if CheckEmail(key) {
			sql = sql.Where("m.email = ?", key)
		} else if CheckIP(key) {
			sql = sql.Where("(m.login_ip = ? or m.reg_ip = ?)", key, key)
		} else if CheckCN(key) {
			sql = sql.Where("(m.nickname like '%?%' or m.logincity like '%?%')", key, key)
		} else {
			sql = sql.Where("(m.username like '%?%' or m.nickname like '%?%')", key, key)
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
func GetForbidByInvitor(uid int) []structure.YlyMemberWeixin {
	var list []structure.YlyMemberWeixin
	sql := dbr.Table("yly_member_weixin").Alias("a").Join("INNER", []string{"yly_member", "b"}, "a.invitor=b.uid").Select("a.uid")
	sql = sql.Where("a.invitor=? and b.group_id != 13", uid)
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetUserForbid() []structure.YlyMember {
	var list []structure.YlyMember
	sql := dbr.Table("game_userfield").Alias("u").Join("INNER", []string{"yly_member", "m"}, "u.uid=m.uid").Select("m.uid")
	sql = sql.Where("m.group_id=13 or u.state=2")
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertForbidLog(data map[string]interface{}) bool {
	_, err := dbr.Table("log_forbid").Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func GetForbidList(uid int, sdate int, edate int, page int, pageSize int) []structure.LogForbid {
	var list []structure.LogForbid
	sql := dbr.Table("log_forbid")
	if uid > 0 {
		sql = sql.Where("uid = ?", uid)
	}
	if sdate > 0 {
		sql = sql.Where("unforbid_time >= ?", sdate)
	}
	if edate > 0 {
		sql = sql.Where("unforbid_time <= ?", edate)
	}
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetForbidListCount(uid int, sdate int, edate int) int {
	sql := dbr.Table("log_forbid").Select("count(1) total")
	if uid > 0 {
		sql = sql.Where("uid = ?", uid)
	}
	if sdate > 0 {
		sql = sql.Where("unforbid_time >= ?", sdate)
	}
	if edate > 0 {
		sql = sql.Where("unforbid_time <= ?", edate)
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}

func GetOpeningTimeBoxes(uid, nowtime int) []*structure.UserBoxesRelated {
	boxes := []*structure.UserBoxesRelated{}
	err := dbr.Table("user_boxes").Alias("a").
		Select("a.*,b.name,b.btype,b.btime,b.bnum,b.quality").
		Join("INNER", []string{"game_boxes", "b"}, "a.bid=b.bid").
		Where("a.uid=?", uid).
		Where("b.btime > 0 and a.state != 1 and a.bstart = 1 and a.optime > ?", nowtime).
		Find(&boxes)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return boxes
}
