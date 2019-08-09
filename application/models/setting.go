package models

import (
	"fmt"
	"github.com/jmesyan/nano/application/models/structure"
	"github.com/sirupsen/logrus"
)

//获取用户语音资源
func GetGameRoleVoices() []structure.GameRoleVoices {
	voices := []structure.GameRoleVoices{}
	err := dbr.Find(&voices)
	if err != nil {
		logger.Println(err.Error())
		return voices
	}
	return voices
}
func InsertGameRoleVoices(data *structure.GameRoleVoices) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Rvid
}

func UpdateGameRoleVoices(rvid int, data *structure.GameRoleVoices) bool {
	_, err := dbr.Where("rvid = ?", rvid).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}

func DeleteGameRoleVoices(rvid int) {
	dbr.Delete(&structure.GameRoleVoices{Rvid: rvid})
}

func GetUserRoleVoices(roid, skillLevel int) []structure.GameRoleVoices {
	var voices []structure.GameRoleVoices
	err := dbr.Where("roid=? and skill_level <= ?", roid, skillLevel).Find(&voices)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return voices
}

func GetGameProps() []structure.GameProps {
	props := []structure.GameProps{}
	err := dbr.Find(&props)
	if err != nil {
		logger.Println(err.Error())
		return props
	}
	return props
}

func InsertGameProps(data *structure.GameProps) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Pid
}
func UpdateGameProps(pid int, data *structure.GameProps) bool {
	_, err := dbr.Where("pid = ?", pid).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteProps(pid int) {
	dbr.Delete(&structure.GameProps{Pid: pid})
}

func GetForbidDeviceuidByDevice(device string) int {
	dev := &structure.YlyForbidDeviceuid{Deviceuid: device}
	counts, err := dbr.Count(dev)
	if err != nil {
		logger.Println("GetForbidDeviceuidByDevice", err.Error())
		return 0
	}
	return int(counts)
}

func GoldBlac(key int, page int, pageSize int) []structure.GoldsBlacklist {
	var data []structure.GoldsBlacklist
	sql := dbr.Table("golds_blacklist")
	if key > 0 {
		sql = sql.Where("uid = ?", key)
	}
	sql = sql.OrderBy("ltime desc")
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func GoldBlacCount(key int) int {
	sql := dbr.Table("golds_blacklist").Select("count(uid) total")
	if key > 0 {
		sql = sql.Where("uid = ?", key)
	}
	totals := &PageTotal{}
	logrus.Info(totals)
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}
func InsertGoldBlac(data *structure.GoldsBlacklist) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Uid
}
func DeleteGoldBlack(uid int) {
	dbr.Delete(&structure.GoldsBlacklist{Uid: uid})
}

func GetGameRoomStat() []structure.GameRoomStat {
	data := []structure.GameRoomStat{}
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return data
	}
	return data
}

func UpdateGameRoomStat(gid int, rtype int, gametype int, data *structure.GameRoomStat) bool {
	data.Gid = gid
	data.Rtype = rtype
	data.Gametype = gametype
	_, err := dbr.Where("gid = ? and rtype = ? and gametype = ?", gid, rtype, gametype).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertGameRoomStat(data *structure.GameRoomStat) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Gid
}
func DeleteGameRoomStat(gid int, rtype int, gametype int) {
	dbr.Delete(&structure.GameRoomStat{Gid: gid, Rtype: rtype, Gametype: gametype})
}
func GetGameGoldsType() []*structure.GameGoldsType {
	data := []*structure.GameGoldsType{}
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return data
	}
	return data
}
func UpdateGameGoldsType(gsid int, data *structure.GameGoldsType) bool {
	data.Gsid = gsid
	_, err := dbr.Where("gsid = ?", gsid).Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertGameGoldsType(data *structure.GameGoldsType) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Gsid
}
func DeleteGameGoldsTyle(gsid int) {
	dbr.Delete(&structure.GameGoldsType{Gsid: gsid})
}

//金币场机器人管理
func GetGameAndroidLevel() []structure.GameAndroidLevel {
	data := []structure.GameAndroidLevel{}
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func GetGameAndroidLevelByGsid(gid int, rtype int) *structure.GameAndroidLevel {
	level := &structure.GameAndroidLevel{}
	//level := new(structure.GameAndroidLevel)
	sql := dbr.Where("gid = ? and rtype = ?", gid, rtype)
	_, err := sql.Get(level)

	if err != nil {
		logger.Println(err)
		return level
	}
	return level
}
func UpdateAndroidLevel(gid int, rtype int, data *structure.GameAndroidLevel) bool {
	data.Gid = gid
	data.Rtype = rtype
	_, err := dbr.Where("gid = ? and rtype = ?", gid, rtype).Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertAndroidLevel(data *structure.GameAndroidLevel) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Gid
}
func DeleteAndroidLevel(gid int, rtype int) {
	dbr.Delete(&structure.GameAndroidLevel{Gid: gid, Rtype: rtype})
}

//金币记牌器
func GetConfP2pJipaiqi() []structure.ConfP2pJipaiqi {
	data := []structure.ConfP2pJipaiqi{}
	err := dbr.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func UpdateGoldCoinCard(gameid int, days int, data *structure.ConfP2pJipaiqi) bool {
	data.Gameid = gameid
	data.Days = days
	_, err := dbr.Where("gameid = ? and days = ?", gameid, days).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGoldCoinCard(data *structure.ConfP2pJipaiqi) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Gameid
}
func DeleteGoldCoinCard(gameid int, days int) {
	dbr.Delete(&structure.ConfP2pJipaiqi{Gameid: gameid, Days: days})
}

//大厅服务器开始
func GetServerFeedback(day int) []map[string]ModelValue {
	sql := fmt.Sprintf("select day, server, count(distinct uid) users, sum(count) count from log_server_feedback where day>=%d group by day, server order by day desc, count desc", day)
	res := DB.Select(sql)
	return MapArrToMV(res)
}

//大厅服务器连接失败日志
func GetReportServerFeedback(start string, end string) []structure.ReportDayServerFeedback {
	data := []structure.ReportDayServerFeedback{}
	err := dbr.Where("ldate >= ? and ldate <= ?", start, end).Find(&data)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return data
}
func UpdateGameHallserver(hid int, data *structure.GameHallserver) bool {
	data.Hid = hid
	_, err := dbr.Where("hid = ?", hid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGameHallserver(data *structure.GameHallserver) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameHallserver(hid int) {
	dbr.Delete(&structure.GameHallserver{Hid: hid})
}

//审核开关
func GetAuditList() []structure.GameAudit {
	var audits []structure.GameAudit
	err := dbr.Table("game_audit").Find(&audits)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return audits
}

func GetGameHallServer() []structure.GameHallserver {
	var ret []structure.GameHallserver
	err := dbr.OrderBy("minusecard desc, maxusecard desc").Find(&ret)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return ret
}

func HallNotice(appid int) []structure.GameHallNotice {
	var list []structure.GameHallNotice
	err := dbr.Where("state = 1 and appid = ?", appid).Desc("id").Limit(50).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

//游戏管理
func GetGameType() []structure.GameType {
	var list []structure.GameType
	err := dbr.Asc("orderby").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameType(data *structure.GameType) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Gsid
}
func UpdateGameType(id int, data *structure.GameType) bool {
	data.Gsid = id
	_, err := dbr.Where("gsid = ?", id).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteGameType(id int) {
	dbr.Delete(&structure.GameType{Gsid: id})
}

func GetGamePets() []structure.GamePets {
	pets := []structure.GamePets{}
	err := dbr.Find(&pets)
	if err != nil {
		logger.Println(err.Error())
		return pets
	}
	return pets
}

func InsertGamePets(data *structure.GamePets) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Id
}
func UpdateGamePets(id int, data *structure.GamePets) bool {
	data.Id = id
	_, err := dbr.Where("id = ?", id).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteGamePets(id int) {
	dbr.Delete(&structure.GamePets{Id: id})
}

//游戏CDN管理
func GetGameCDNS() []structure.GameCdns {
	var list []structure.GameCdns
	err := dbr.Desc("bid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameCDNS(data *structure.GameCdns) string {
	_, err := dbr.Insert(data)
	if err != nil {
		return ""
	}
	return data.Bid
}
func UpdateGameCDNS(id string, data *structure.GameCdns) bool {
	data.Bid = id
	_, err := dbr.Where("bid = ?", id).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteGameCDNS(id string) {
	dbr.Delete(&structure.GameCdns{Bid: id})
}

//WEB服务器管理
func GetGameServers() []structure.GameServers {
	var list []structure.GameServers
	err := dbr.Asc("usetime").Asc("uses").Asc("hid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameServers(data *structure.GameServers) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Hid
}
func UpdateGameServers(id int, data *structure.GameServers) bool {
	data.Hid = id
	_, err := dbr.Where("hid = ?", id).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteGameServers(id int) {
	dbr.Delete(&structure.GameServers{Hid: id})
}

//服务器管理
func GetGameServers2() []structure.GameServers2 {
	var list []structure.GameServers2
	err := dbr.Desc("usetime").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameServers2(data *structure.GameServers2) string {
	_, err := dbr.Insert(data)
	if err != nil {
		return ""
	}
	return data.Ip
}
func UpdateGameServers2(ip string, data *structure.GameServers2) bool {
	//data.Ip = ip
	_, err := dbr.Where("ip = ?", ip).Update(data)
	if err != nil {
		return false
	}
	return true
}
func DeleteGameServers2(ip string) {
	dbr.Delete(&structure.GameServers2{Ip: ip})
}

//段位
func GetGameRankStat() []structure.GameRankStat {
	var list []structure.GameRankStat
	err := dbr.Asc("stars").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list

}
func UpdateGameRankStat(stars int, data *structure.GameRankStat) bool {
	data.Stars = stars
	_, err := dbr.Where("stars = ?", stars).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertGameRankStat(data *structure.GameRankStat) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Stars
}
func DeleteGameRankStat(stars int) {
	dbr.Delete(&structure.GameRankStat{Stars: stars})
}

func GetMail(key string, fromid int, toid int, isShow int, page int, pageSize int) []structure.YlyPost {
	var data []structure.YlyPost
	sql := dbr.Table("yly_post")
	if len(key) > 0 {
		sql = sql.Where("title like '%?%' or ps like '%?%'", key, key)
	}
	if fromid > 0 {
		sql = sql.Where("fromid =?", fromid)
	}
	if toid > 0 {
		sql = sql.Where("toid =?", toid)
	}
	if isShow > -1 {
		sql = sql.Where("isShow = ?", isShow)
	}
	sql = sql.OrderBy("pid desc")
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func GetMailCount(key string, fromid int, toid int, isShow int) int {
	sql := dbr.Table("yly_post").Select("count(1) total")
	if len(key) > 0 {
		sql = sql.Where("title like '%?%' or ps like '%?%'", key, key)
	}
	if fromid > 0 {
		sql = sql.Where("fromid = ?", fromid)
	}
	if toid > 0 {
		sql = sql.Where("toid = ?", toid)
	}
	if isShow > -1 {
		sql = sql.Where("isShow = ?", isShow)
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}
func GetYlyPostByPid(pid int) []structure.YlyPost {
	var list []structure.YlyPost
	err := dbr.Where("pid=?", pid).Limit(1).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertMail(data *structure.YlyPost) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Pid
}
func InsertAllMail(data []structure.YlyPost) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		return false
	}
	return true
}
func GetUserMailByToid(uid int) []structure.YlyPost {
	var list []structure.YlyPost
	err := dbr.Where("toid=?", uid).Limit(200).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func UpdateUserEmails(pid int, data *structure.YlyPost) bool {
	_, err := dbr.Where("pid = ?", pid).MustCols("status,type,wppack").Update(data)
	if err != nil {
		return false
	}
	return true
}
func UpdateUserAllEmails(pids []int, data map[string]interface{}) bool {
	_, err := dbr.Table("yly_post").In("pid", pids).Update(data)
	if err != nil {
		return false
	}
	return true
}
func GetUserMailByPid(pid int) *structure.YlyPost {
	data := &structure.YlyPost{Pid: pid}
	has, _ := dbr.Where("pid = ?", pid).Get(data)
	if has {
		return data
	}
	return nil

}
func DeleteMail(pid int) {
	dbr.Delete(&structure.YlyPost{Pid: pid})
}
func DeleteAllMail(pids []int) {
	dbr.In("pid", pids).Delete(&structure.YlyPost{})
}
func InsertUserSendPost(data []structure.UserSendPost) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func GetUserSendPost() []structure.UserSendPost {
	var list []structure.UserSendPost
	err := dbr.Limit(2000, 0).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func DeleteUserSendPost(uids []int) {
	dbr.In("uid", uids).Delete(&structure.UserSendPost{})
}
func GetGameSkills() []structure.GameSkills {
	var list []structure.GameSkills
	err := dbr.Desc("skid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func UpdateGameSkills(id int, data *structure.GameSkills) bool {
	data.Skid = id
	_, err := dbr.Where("skid=?", id).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGameSkills(data *structure.GameSkills) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Skid
}
func DeleteGameSkills(id int) {
	dbr.Delete(&structure.GameSkills{Skid: id})
}
func GetGameRoles() []structure.GameRoles {
	var list []structure.GameRoles
	err := dbr.Desc("roid").Asc("roid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func GetDefaultRoles() []structure.GameRoles {
	var list []structure.GameRoles
	err := dbr.Where("isdefault = ?", 1).Asc("roid").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func UpdateGameRoles(roid int, data *structure.GameRoles) bool {
	data.Roid = roid
	_, err := dbr.Where("roid=?", roid).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertGameRoles(data *structure.GameRoles) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Roid
}
func DeleteGameRoles(roid int) {
	dbr.Delete(&structure.GameRoles{Roid: roid})
}

func GetGameShopConfig(args ...int) []structure.GameShopConfig {
	channel := 0
	if len(args) > 0 {
		channel = args[0]
	}
	var list []structure.GameShopConfig
	sql := dbr.Desc("scid")
	if channel > 0 {
		sql = sql.Where("channel= ?", channel)
	}
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func UpdateGameShopConfig(id int, data *structure.GameShopConfig) bool {
	data.Scid = id
	_, err := dbr.Where("scid=?", id).AllCols().Update(data)
	if err != nil {
		return false
	}
	return true
}
func InsertGameShopConfig(data *structure.GameShopConfig) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Scid
}
func DeleteGameShopConfig(id int) {
	dbr.Delete(&structure.GameShopConfig{Scid: id})
}
func GetGameRoleExps() []structure.GameRoleExps {
	var list []structure.GameRoleExps
	err := dbr.Desc("quality").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func UpdateGameRoleExps(quality int, role_level int, data *structure.GameRoleExps) bool {
	_, err := dbr.Where("quality = ? and role_level = ? ", quality, role_level).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGameRoleExps(data *structure.GameRoleExps) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Quality
}
func DeleteGameRoleExps(quality int, role_level int) {
	dbr.Delete(&structure.GameRoleExps{Quality: quality, RoleLevel: role_level})
}
func GetGameRoleSkills() []structure.GameRoleSkills {
	var list []structure.GameRoleSkills
	err := dbr.Desc("quality").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func UpdateGameRoleSkills(quality int, skill_level int, data *structure.GameRoleSkills) bool {
	_, err := dbr.Where("quality = ? and skill_level = ? ", quality, skill_level).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGameRoleSkills(data *structure.GameRoleSkills) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Quality
}
func DeleteGameRoleSkills(quality int, skill_level int) {
	dbr.Delete(&structure.GameRoleSkills{Quality: quality, SkillLevel: skill_level})
}

func GetMatchSeasonAwards(args ...int) []structure.MatchSeasonAwards {
	season := 0
	if len(args) > 0 {
		season = args[0]
	}
	var list []structure.MatchSeasonAwards
	sql := dbr.OrderBy("season asc, srank asc")
	if season > 0 {
		sql = sql.Where("season = ?", season)
	}
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func UpdateMatchSeasonAwards(id int, data *structure.MatchSeasonAwards) bool {
	_, err := dbr.Where("aid=?", id).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func InsertMatchSeasonAwards(data *structure.MatchSeasonAwards) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Aid
}
func DeleteMatchSeasonAwards(id int) {
	dbr.Delete(&structure.MatchSeasonAwards{Aid: id})
}

func GetSeasonLastAwards(season int) string {
	var awards string
	_, err := dbr.Table("match_season_awards").Where("season=? and erank=?", season, 9999).Cols("pids").Get(&awards)
	if err != nil {
		logger.Println(err.Error())
	}
	return awards
}

func GetGameBoxes(args ...int) []structure.GameBoxes {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	var list []structure.GameBoxes
	sql := dbr.Desc("bid")
	if state > 0 {
		switch state {
		case 1: //游戏宝箱
			sql = sql.Where("btype=?", 1).Select("bid,name,btype,bnum,btime,frees,price,tab,objects")
		case 2: //抽卡宝箱
			sql = sql.Where("btype=?", 2).Select("bid,name,btype,bnum,btime,frees,price,tab,objects")
		}
	}
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func UpdateGameBoxes(bid int, data *structure.GameBoxes) bool {
	_, err := dbr.Where("bid=?", bid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func InsertGameBoxes(data *structure.GameBoxes) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Bid
}
func DeleteBoxesDelete(bid int) {
	dbr.Delete(&structure.GameBoxes{Bid: bid})
}

func GetGameWinStreak() []structure.GameWinStreak {
	var list []structure.GameWinStreak
	err := dbr.Table("game_win_streak").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameWinStreak(data *structure.GameWinStreak) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func UpdateGameWinStreak(wins int, data *structure.GameWinStreak) bool {
	_, err := dbr.Table("game_win_streak").Where("wins = ?", wins).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameWinStreak(wins int) {
	dbr.Delete(&structure.GameWinStreak{Wins: wins})
}

func GetGameTask(ttype int) []structure.GameTask {
	var list []structure.GameTask
	sql := dbr.Table("game_task")
	if ttype > 0 {
		sql = sql.Where("type = ?", ttype)
	}
	err := sql.Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameTask(data *structure.GameTask) int {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return data.Tkid
}
func UpdateGameTask(id int, data *structure.GameTask) bool {
	_, err := dbr.Table("game_task").Where("tkid = ?", id).AllCols().Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameTask(id int) {
	dbr.Delete(&structure.GameTask{Tkid: id})
}

func GetGameLiveness() []structure.GameLiveness {
	var list []structure.GameLiveness
	err := dbr.Table("game_liveness").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameLiveness(data *structure.GameLiveness) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func UpdateGameLiveness(livetype int, liveness int, data *structure.GameLiveness) bool {
	_, err := dbr.Table("game_liveness").Where("livetype = ? and liveness = ?", livetype, liveness).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameLiveness(livetype int, liveness int) {
	dbr.Delete(&structure.GameLiveness{Livetype: livetype, Liveness: liveness})
}

func GetGameRatioFactor() []structure.GameRatioFactor {
	var list []structure.GameRatioFactor
	err := dbr.Table("game_ratio_factor").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}
func InsertGameRatioFactor(data *structure.GameRatioFactor) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func UpdateGameRatioFactor(ratio int, data *structure.GameRatioFactor) bool {
	_, err := dbr.Table("game_ratio_factor").Where("ratio = ? ", ratio).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}
func DeleteGameRatioFactor(ratio int) {
	dbr.Delete(&structure.GameRatioFactor{Ratio: ratio})
}

//大厅公告管理
func GetHallNotice(htype int, page int, pageSize int) []structure.GameHallNotice {
	var data []structure.GameHallNotice
	sql := dbr.Table("game_hall_notice")
	if htype > 0 {
		sql = sql.Where("appid=?", htype)
	}
	sql = sql.OrderBy("id desc")
	offset := (page - 1) * pageSize
	limit := pageSize
	sql = sql.Limit(limit, offset)
	err := sql.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}
func GetHallNoticeCount(htype int) int {
	sql := dbr.Table("game_hall_notice").Select("count(1) total")
	if htype > 0 {
		sql = sql.Where("appid = ?", htype)
	}
	totals := &PageTotal{}
	_, err := sql.Get(totals)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return totals.Total
}

func GetHallNoticeById(id int) *structure.GameHallNotice {
	list := &structure.GameHallNotice{}
	_, err := dbr.Table("game_hall_notice").Where("id=?", id).Get(list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func InsertHallNotice(data *structure.GameHallNotice) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func UpdateHallNotice(id int, data *structure.GameHallNotice) bool {
	_, err := dbr.Where("id=?", id).AllCols().Update(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func DeleteHallNotice(id int) {
	dbr.Delete(&structure.GameHallNotice{Id: id})
}
