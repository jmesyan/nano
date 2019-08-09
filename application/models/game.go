package models

import (
	"fmt"
	"github.com/jmesyan/nano/application/models/structure"
	"time"
)

func GetGameConfig(id int) *structure.GameConfig {
	config := &structure.GameConfig{Cid: id}
	has, _ := dbr.Where("uid=?", id).Get(config)
	if has {
		return config
	}
	return nil
}

func UpdateGameConfig(id, v int, t string) {
	sql := fmt.Sprintf("insert into game_config(cid,cvalue,cdesc) values(%d,%d,%s) ON DUPLICATE KEY update cvalue=%d,cdesc=%s", id, v, t, v, t)
	dbr.Exec(sql)
}

func GetAnnouncement() []structure.P2pAnnouncement {
	var announcement []structure.P2pAnnouncement
	err := dbr.Find(&announcement)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return announcement
}

func UpdateAnnouncement(aid int, data *structure.P2pAnnouncement) {
	dbr.Where("aid=?", aid).Update(data)
}

func InsertAnnouncement(data *structure.P2pAnnouncement) int {
	_, err := dbr.Insert(data)
	if err != nil {
		return 0
	}
	return data.Aid
}

func DeleteAnnouncement(aid int) {
	dbr.Delete(&structure.P2pAnnouncement{Aid: aid})
}

func InsertFeedback(data *structure.P2pFeedback) int {
	sql := fmt.Sprintf("select count(id) times from p2p_feedback where md5='%s'", data.Md5)
	res := DB.SelectOne(sql)
	ress := MapFaceToMV(res)
	if ress["times"].Intval > 0 {
		sql2 := fmt.Sprintf("update p2p_feedback set state=1, times = times + 1, last_time = %v, pic='%v'  where md5 = '%v'", data.LastTime, data.Pic, data.Md5)
		DB.Update(sql2)
	} else {
		dbr.Insert(data)
	}
	return 1
}
func InsertMotor(data *structure.YlyMotor) {
	dbr.Insert(data)
}

func CrontabReportMinuteOnline() int {
	sql := "call crontab_report_minute_online()"
	res := DB.SelectOne(sql)
	ret := MapFaceToMV(res)
	return ret["count"].Intval
}

func CrontabReportDay(day string) int {
	sql := fmt.Sprintf("call crontab_report_day(%s)", day)
	res := DB.SelectOne(sql)
	ret := MapFaceToMV(res)
	return ret["count"].Intval
}

func CrontabReportMonth(month string) int {
	sql := fmt.Sprintf("call crontab_report_month(%s)", month)
	res := DB.SelectOne(sql)
	ret := MapFaceToMV(res)
	return ret["count"].Intval
}

func GetGoldsRank() []structure.UserGoldsRank {
	var list []structure.UserGoldsRank
	err := dbr.Desc("golds").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func GetSeasonRank(season int) []structure.UserSeasonRank {
	var list []structure.UserSeasonRank
	err := dbr.Where("season = ?", season).OrderBy("grade_rank asc").Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func GetSeasonRankByUid(season, uid int) *structure.UserSeasonRank {
	info := &structure.UserSeasonRank{}
	ret, err := dbr.Where("season = ? and uid = ?", season, uid).Get(info)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return info
}
func UpdatePageEnters(pp string) int {
	day := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	sql := fmt.Sprintf("insert into report_page_count(day,page,enters) values(%v, %v, %v) on duplicate key update enters = enters+1 ", day, pp, 1)
	DB.Insert(sql)
	return 0
}
func UpdatePageDowns(pp string) int {
	day := time.Unix(time.Now().Unix(), 0).Format("20060504")
	sql := fmt.Sprintf("insert into report_page_count(day,page,downs) values(%v, %v, %v) on duplicate key update downs=downs+1", day, pp, 1)
	DB.Insert(sql)
	return 0
}

func GetUserGameInfo(uid int) *structure.UserGameInfo {
	info := &structure.UserGameInfo{}
	ret, err := dbr.Where("uid = ?", uid).Get(info)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return info
}

func InsertServerFeedBack(uid int, day, server, ip string) bool {
	sql := fmt.Sprintf("insert into log_server_feedback(day,server,uid,ip,count) values(%s,'%s',%d,'%s',1) ON DUPLICATE KEY update count=count+1,ip='%s'", day, server, uid, ip, ip)
	_, err := dbr.Exec(sql)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}
func GetOnlineNum(date int, start int, end int) []structure.LogOnlineNum {
	var data []structure.LogOnlineNum
	sql := dbr.Table("log_online_num")
	sql = sql.Where("day = ?", date)

	if start == 0 && end == 0 {
		sql = sql.Where("rtype >= ? and rtype <= ? ", start, end)
	}
	sql = sql.OrderBy("h asc,m asc")
	err := sql.Find(&data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data

}
