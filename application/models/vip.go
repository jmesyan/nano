package models

import "github.com/jmesyan/nano/application/models/structure"

func GetGameVipLevel() []*structure.GameVipLevel {
	var list []*structure.GameVipLevel
	err := dbr.Asc("vip_level").Find(&list)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return list
}

func InsertGameVipLevel(data *structure.GameVipLevel) bool {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func UpdateGameVipLevel(id int, data *structure.GameVipLevel) bool {
	_, err := dbr.Where("vip_level=?", id).AllCols().Update(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func DeleteGameVipLevel(id int) {
	dbr.Delete(&structure.GameVipLevel{VipLevel: id})
}

func InsertLogVipLevel(data *structure.LogVipLevel) {
	_, err := dbr.Insert(data)
	if err != nil {
		logger.Println(err)
	}
}

func GetLogVipLevel(uid int) []*structure.LogVipLevel {
	var list []*structure.LogVipLevel
	err := dbr.Where("uid=?", uid).Find(&list)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return list
}

func UpdateLogVipLevel(uid, level int, data map[string]interface{}) bool {
	_, err := dbr.Table("log_vip_level").Where("uid=? and nlevel=?", uid, level).Update(data)
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}
