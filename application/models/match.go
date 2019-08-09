package models

import "github.com/jmesyan/nano/application/models/structure"

func GetMatchGradeSeason() int {
	var season int
	_, err := dbr.Table("match_season").Where("sid=?", 1).Select("max(season) season").Cols("season").Get(&season)
	if err != nil {
		logger.Println(err.Error())
	}
	return season
}

func UpdateMatchGradeSeason(season int, data map[string]interface{}) bool {
	_, err := dbr.Table("match_season").Where("sid = ? and season=?", 1, season).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func GetSeasonRankState(season int) int {
	var state int
	_, err := dbr.Table("match_season").Where("sid=? and season=?", 1, season).Cols("isrank").Get(&state)
	if err != nil {
		logger.Println(err.Error())
		return 0
	}
	return state
}

func GetSeasonScoreRanks(season, num int) []structure.UserMatchScore {
	var list []structure.UserMatchScore
	err := dbr.Table("user_match_score").Where("season=?", season).Desc("score").Limit(num).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func GetUserMatchScore(season, uid int) *structure.UserMatchScore {
	data := &structure.UserMatchScore{}
	_, err := dbr.Where("season=? and uid=?", season, uid).Get(data)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return data
}

func UpdateUserMatchScore(season, uid int, data map[string]interface{}) bool {
	_, err := dbr.Table("user_match_score").Where("season=? and uid=?", season, uid).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

func UpdateSeasonNonRanks(season int, data map[string]interface{}) bool {
	_, err := dbr.Table("user_match_score").Where("season=? and rank=0", season).Update(data)
	if err != nil {
		logger.Println(err.Error())
		return false
	}
	return true
}

//获取排位赛个人记录
func GetMatchGradeUserRecord(uid, limit int) []structure.LogCreateUserRankLiushui {
	list := []structure.LogCreateUserRankLiushui{}
	err := dbr.Where("uid=?", uid).Desc("lid").Limit(limit).Find(&list)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	return list
}

func GetGradeRankStars(season, rank int) int {
	var stars int
	ret, err := dbr.Table("user_season_rank").Where("season=? and grade_rank=?", season, rank).Cols("grade_stars").Get(&stars)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return 0
	}
	return stars
}

func GetGradeLastRank(season int) *structure.UserSeasonRank {
	info := &structure.UserSeasonRank{}
	ret, err := dbr.Table("user_season_rank").Where("season=? ", season).Desc("grade_rank").Get(info)
	if !ret || err != nil {
		if err != nil {
			logger.Println(err)
		}
		return nil
	}
	return info
}
