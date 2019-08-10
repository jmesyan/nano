package cache

import (
	"encoding/json"
	"fmt"
	"github.com/jmesyan/nano/application/models"
	"github.com/jmesyan/nano/application/models/structure"
	"github.com/jmesyan/nano/utils"
	"gopkg.in/redis.v4"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	CacheManager *cacheManager
	logger       = log.New(os.Stderr, "cache", log.LstdFlags|log.Llongfile)
)

func init() {
	Week = 604800 * time.Second
	Day = 86400 * time.Second
	Hour = 3600 * time.Second
	Minute = 60 * time.Second
	CacheManager = newCacheManager()
}

var (
	Week   time.Duration
	Day    time.Duration
	Hour   time.Duration
	Minute time.Duration
)

type cacheManager struct {
	Config string
	Client *redis.Client
	Prefix string
}

func newCacheManager() *cacheManager {
	cm := &cacheManager{
		Prefix: "stargames",
	}

	addr := strings.Join([]string{"127.0.0.1", "6379"}, ":")
	redisOptions := &redis.Options{
		Addr: addr,
		DB:   0,
	}
	cm.Client = redis.NewClient(redisOptions)

	_, err := cm.Client.Ping().Result()
	if err != nil {
		log.Fatal("redis", err.Error())
	}
	if cm.Client == nil {
		log.Fatal("can't find the cache client")
	}
	return cm
}

func (cm *cacheManager) Set(key string, value interface{}, expiration time.Duration) bool {
	ckey := cm.Prefix + ":" + key
	err := cm.Client.Set(ckey, value, expiration).Err()
	if err != nil {
		logger.Printf("set key=>%s, error=>%s", ckey, err.Error())
		return false
	}
	return true
}
func (cm *cacheManager) Get(key string) string {
	ckey := cm.Prefix + ":" + key
	val, err := cm.Client.Get(ckey).Result()
	if err != nil {
		logger.Printf("get key=>%s, error=>%s", ckey, err.Error())
		return ""
	}
	return val
}

func (cm *cacheManager) GetPatternKeys(prefix string) []string {
	key := cm.CacheKey(prefix, "*")
	keys := cm.Client.Keys(cm.Prefix + ":" + key)
	return keys.Val()
}
func (cm *cacheManager) GetInt(key string) int {
	val := cm.Get(key)
	return utils.StringToInt(val)
}
func (cm *cacheManager) GetSet(key string, value interface{}) string {
	ckey := cm.Prefix + ":" + key
	val, err := cm.Client.GetSet(ckey, value).Result()
	if err != nil {
		logger.Println(err)
		return ""
	}
	return val
}

func (cm *cacheManager) Del(keys ...string) bool {
	for i, v := range keys {
		keys[i] = cm.Prefix + ":" + v
	}
	_, err := cm.Client.Del(keys...).Result()
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func (cm *cacheManager) DelKeys(keys ...string) bool {
	_, err := cm.Client.Del(keys...).Result()
	if err != nil {
		logger.Println(err)
		return false
	}
	return true
}

func (cm *cacheManager) CacheKey(prefix string, params ...interface{}) string {
	key := prefix
	len := len(params)
	if len > 0 {
		for i := 0; i < len; i++ {
			param := params[i]
			switch t := param.(type) {
			case string:
				key += "_" + t
			case int:
				key += "_" + strconv.Itoa(t)
			}
		}
	}
	return key
}

type MaintenanceInfo struct {
	Type int `json:"type"`
	T    int `json:"t"`
	Time int `json:"time"`
}

func (cm *cacheManager) GetMaintence() *MaintenanceInfo {
	key := cm.CacheKey("maintenance")
	val := cm.Get(key)
	var info *MaintenanceInfo
	if val == "" {
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	return info
}

func (cm *cacheManager) RemoveMaintence() bool {
	key := cm.CacheKey("maintenance")
	return cm.Del(key)
}

func (cm *cacheManager) RemoveServerManintence(gsid string) bool {
	key := cm.CacheKey("maintenance", gsid)
	return cm.Del(key)
}

func (cm *cacheManager) GetUser(uid int) *structure.YlyUser {
	key := cm.CacheKey("GetUser", uid)
	val := cm.Get(key)
	var info *structure.YlyUser
	if val == "" {
		info = models.GetUserByID(uid)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	return info
}

func (cm *cacheManager) RemoveUser(uid int) bool {
	key := cm.CacheKey("GetUser", uid)
	return cm.Del(key)
}

func (cm *cacheManager) GetGameConfig(id int) *structure.GameConfig {
	key := cm.CacheKey("GetGameConfig", id)
	val := cm.Get(key)
	info := &structure.GameConfig{}
	if val == "" {
		info = models.GetGameConfig(id)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return info
		}
		cm.Set(key, val, Day)
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	return info
}

func (cm *cacheManager) RemoveGameConfig(id int) bool {
	key := cm.CacheKey("GetGameConfig", id)
	return cm.Del(key)
}

func (cm *cacheManager) GetAnnouncement() map[int]structure.P2pAnnouncement {
	key := cm.CacheKey("GetAnnouncement")
	val := cm.Get(key)
	var info []structure.P2pAnnouncement
	ret := make(map[int]structure.P2pAnnouncement)
	if val == "" {
		info = models.GetAnnouncement()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Aid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Aid] = v
	}
	return ret
}

func (cm *cacheManager) RemoveAnnouncement() bool {
	key := cm.CacheKey("GetAnnouncement")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameProps() map[int]structure.GameProps {
	key := cm.CacheKey("GetGameProps")
	val := cm.Get(key)
	var info []structure.GameProps
	ret := make(map[int]structure.GameProps)
	if val == "" {
		info = models.GetGameProps()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Pid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Pid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveGameProps() bool {
	key := cm.CacheKey("GetGameProps")
	cm.Del(key)

	keys := cm.GetPatternKeys("GetUserProps")
	return cm.DelKeys(keys...)
}

func (cm *cacheManager) GetUserProps(uid int, args ...int) map[int]structure.UserPropsRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	key := cm.CacheKey("GetUserProps", uid, state)
	val := cm.Get(key)
	var info []structure.UserPropsRelated
	ret := make(map[int]structure.UserPropsRelated)
	if val == "" {
		info = models.GetUserProps(uid, state)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Upid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Upid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveUserProps(uid int) bool {
	keys := cm.GetPatternKeys(fmt.Sprintf("GetUserProps_%d", uid))
	return cm.DelKeys(keys...)
}

func (cm *cacheManager) GetUserRoles(uid int, args ...int) map[int]structure.UserRolesRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	key := cm.CacheKey("GetUserRoles", uid, state)
	val := cm.Get(key)
	var info []structure.UserRolesRelated
	ret := make(map[int]structure.UserRolesRelated)
	if val == "" {
		info = models.GetUserRoles(uid, state)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Uroid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Uroid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveUserRoles(uid int) bool {
	for i := 0; i <= 3; i++ {
		key := cm.CacheKey("GetUserRoles", uid, i)
		cm.Del(key)
	}
	return true
}

func (cm *cacheManager) GetGameGoldsType() map[int]*structure.GameGoldsType {
	key := cm.CacheKey("GetGameGoldsType")
	val := cm.Get(key)
	var info []*structure.GameGoldsType
	ret := make(map[int]*structure.GameGoldsType)
	if val == "" {
		info = models.GetGameGoldsType()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Gsid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Gsid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveGameGoldsType() bool {
	key := cm.CacheKey("GetGameGoldsType")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameRoomStat() map[int]map[int]map[int]structure.GameRoomStat {
	key := cm.CacheKey("GetGameRoomStat")
	val := cm.Get(key)
	var info []structure.GameRoomStat
	ret := make(map[int]map[int]map[int]structure.GameRoomStat)
	if val == "" {
		info = models.GetGameRoomStat()

		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			if _, ok := ret[v.Gid]; !ok {
				ret[v.Gid] = make(map[int]map[int]structure.GameRoomStat)
			}
			if _, ok := ret[v.Gid][v.Rtype]; !ok {
				ret[v.Gid][v.Rtype] = make(map[int]structure.GameRoomStat)
			}

			ret[v.Gid][v.Rtype][v.Gametype] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		if _, ok := ret[v.Gid]; !ok {
			ret[v.Gid] = make(map[int]map[int]structure.GameRoomStat)
		}
		if _, ok := ret[v.Gid][v.Rtype]; !ok {
			ret[v.Gid][v.Rtype] = make(map[int]structure.GameRoomStat)
		}
		ret[v.Gid][v.Rtype][v.Gametype] = v
	}
	return ret
}
func (cm *cacheManager) GetRoomStat() map[int]map[int]map[string]interface{} {
	list := CacheManager.GetGameRoomStat()
	ret := make(map[int]map[int]map[string]interface{})
	for k1, v1 := range list {
		for k2, v2 := range v1 {
			for _, v3 := range v2 {
				if _, ok := ret[k1]; !ok {
					ret[k1] = make(map[int]map[string]interface{})
				}
				if _, ok := ret[k1][k2]; !ok {
					ret[k1][k2] = make(map[string]interface{})
				}
				ret[k1][k2]["rtype"] = v3.Rtype
				ret[k1][k2]["gametype"] = v3.Gametype
				ret[k1][k2]["room_name"] = v3.RoomName
				ret[k1][k2]["goldsless"] = v3.Goldsless
				ret[k1][k2]["goldsmin"] = v3.Goldsmin
				ret[k1][k2]["goldsmax"] = v3.Goldsmax
				ret[k1][k2]["s2g"] = v3.S2g
				ret[k1][k2]["systax"] = v3.Systax
			}
		}
	}
	return ret
}
func (cm *cacheManager) RemoveGameRoomStat() bool {
	key := cm.CacheKey("GetGameRoomStat")
	return cm.Del(key)
}

func (cm *cacheManager) GetMobileUser(uid int) *structure.YlyMemberMobile {
	key := cm.CacheKey("GetMobileUser", uid)
	val := cm.Get(key)
	var info *structure.YlyMemberMobile
	if val == "" {
		info = models.GetMobileUser(uid)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	return info
}

func (cm *cacheManager) RemoveMobileUser(uid int) bool {
	key := cm.CacheKey("GetMobileUser", uid)
	return cm.Del(key)
}

func (cm *cacheManager) CetForbidDeviceuidByDevice(deviceuid string) int {
	key := cm.CacheKey("CetForbidDeviceuidByDevice", deviceuid)
	val := cm.Get(key)
	var info int
	if val == "" {
		info = models.GetForbidDeviceuidByDevice(deviceuid)
		val = utils.IntToString(info)
		cm.Set(key, val, Day)
		return info
	}
	info = utils.StringToInt(val)
	return info
}
func (cm *cacheManager) GetAuditMap(args ...string) int {
	var platform, ver = "", ""
	if len(args) > 0 {
		platform = args[0]
	}
	if len(args) > 1 {
		ver = args[1]
	}

	key := cm.CacheKey("GetAuditList")
	val := cm.Get(key)
	var info []structure.GameAudit
	if val == "" {
		info = models.GetAuditList()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return 0
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return 0
		}
	}
	data := make(map[string]string)
	for _, value := range info {
		if len(platform) == 0 {
			data[value.Platform] = value.Version
			continue
		}
		if platform != value.Platform {
			continue
		}
		data[value.Version] = "1"
	}
	if _, ok := data[ver]; ok && len(ver) > 0 {
		return 1
	}
	return 0
}

func (cm *cacheManager) RemoveGameAudit() bool {
	key := cm.CacheKey("GetAuditList")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameHallServer() map[int]structure.GameHallserver {
	key := cm.CacheKey("GetGameHallServer")
	val := cm.Get(key)
	var info []structure.GameHallserver
	if val == "" {
		info = models.GetGameHallServer()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameHallserver)
	for _, server := range info {
		list[server.Hid] = server
	}
	return list
}

func (cm *cacheManager) RemoveGameHallServer() bool {
	key := cm.CacheKey("GetGameHallServer")
	return cm.Del(key)
}

func (cm *cacheManager) GetHallNotice(appid int) []structure.GameHallNotice {
	key := cm.CacheKey("GetHallNotice", appid)
	val := cm.Get(key)
	var info []structure.GameHallNotice
	if val == "" {
		info = models.HallNotice(appid)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	return info
}

func (cm *cacheManager) RemoveHallNotice(appid int) bool {
	key := cm.CacheKey("GetHallNotice", appid)
	return cm.Del(key)
}

func (cm *cacheManager) GetGameType() map[int]structure.GameType {
	key := cm.CacheKey("GetGameType")
	val := cm.Get(key)
	var info []structure.GameType
	if val == "" {
		info = models.GetGameType()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameType)
	for _, value := range info {
		list[value.Gsid] = value

	}
	return list
}

func (cm *cacheManager) RemoveGameType() bool {
	key := cm.CacheKey("GetGameType")
	return cm.Del(key)
}

func (cm *cacheManager) GetGamePets() map[int]structure.GamePets {
	key := cm.CacheKey("GetGamePets")
	val := cm.Get(key)
	var info []structure.GamePets
	ret := make(map[int]structure.GamePets)
	if val == "" {
		info = models.GetGamePets()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Id] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Id] = v
	}
	return ret
}
func (cm *cacheManager) RemoveGamePets() bool {
	key := cm.CacheKey("GetGamePets")
	return cm.Del(key)
}
func (cm *cacheManager) GetConfP2pJipaiqi() map[int]map[int]structure.ConfP2pJipaiqi {
	key := cm.CacheKey("ConfP2pJipaiqi")
	val := cm.Get(key)
	var info []structure.ConfP2pJipaiqi
	ret := make(map[int]map[int]structure.ConfP2pJipaiqi)
	if val == "" {
		info = models.GetConfP2pJipaiqi()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			if _, ok := ret[v.Gameid]; !ok {
				ret[v.Gameid] = make(map[int]structure.ConfP2pJipaiqi)
			}
			ret[v.Gameid][v.Days] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		if _, ok := ret[v.Gameid]; !ok {
			ret[v.Gameid] = make(map[int]structure.ConfP2pJipaiqi)
		}
		ret[v.Gameid][v.Days] = v
	}
	return ret
}
func (cm *cacheManager) RemoveConfP2pJipaiqi() bool {
	key := cm.CacheKey("ConfP2pJipaiqi")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameCDNS() map[string]structure.GameCdns {
	key := cm.CacheKey("GetGameCDNS")
	val := cm.Get(key)
	var info []structure.GameCdns
	if val == "" {
		info = models.GetGameCDNS()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[string]structure.GameCdns)
	for _, value := range info {
		list[value.Bid] = value

	}
	return list
}
func (cm *cacheManager) RemoveGameCDNS() bool {
	key := cm.CacheKey("GetGameCDNS")
	return cm.Del(key)
}
func (cm *cacheManager) GetUserFakeList() map[int]map[int]structure.YlyMemberFake {
	key := cm.CacheKey("GetUserFakeList")
	val := cm.Get(key)
	var info []structure.YlyMemberFake
	if val == "" {
		info = models.GetUserFakeList()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]map[int]structure.YlyMemberFake)
	for _, value := range info {
		if _, ok := list[value.Uid]; !ok {
			list[value.Uid] = make(map[int]structure.YlyMemberFake)
		}
		list[value.Uid][value.Fakeuid] = value

	}
	return list
}
func (cm *cacheManager) RemoveUserFakeList() bool {
	key := cm.CacheKey("GetUserFakeList")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameServers() map[int]structure.GameServers {
	key := cm.CacheKey("GetGameServers")
	val := cm.Get(key)
	var info []structure.GameServers
	if val == "" {
		info = models.GetGameServers()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameServers)
	for _, value := range info {
		list[value.Hid] = value

	}
	return list
}
func (cm *cacheManager) RemoveGameServers() bool {
	key := cm.CacheKey("GetGameServers")
	return cm.Del(key)
}

func (cm *cacheManager) GetOnlineUsers() map[int]structure.YlyOnline {
	key := cm.CacheKey("GetOnlineUsers")
	val := cm.Get(key)
	var info []structure.YlyOnline
	if val == "" {
		info = models.GetOnlineUsers()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Minute)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.YlyOnline)
	for _, v := range info {
		list[v.Userid] = v
	}
	return list
}
func (cm *cacheManager) RemoveOnlineUsers() bool {
	key := cm.CacheKey("GetOnlineUsers")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameRankStat() map[int]structure.GameRankStat {
	key := cm.CacheKey("GetGameRankStat")
	val := cm.Get(key)
	var info []structure.GameRankStat
	if val == "" {
		info = models.GetGameRankStat()
		logger.Printf("%+v", info)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]structure.GameRankStat)
	for _, val := range info {
		list[val.Stars] = val
	}
	return list
}

func (cm *cacheManager) RemoveGameRankStat() bool {
	key := cm.CacheKey("GetGameRankStat")
	return cm.Del(key)
}
func (cm *cacheManager) CacheLockStart(key string, seconds int) bool {
	key = "Lock:" + key
	value := cm.Get(key)
	if value != "" {
		return false
	}
	cm.Set(key, 1, time.Duration(seconds)*time.Second)
	return true
}

func (cm *cacheManager) CacheLockEnd(key string) {
	key = "Lock:" + key
	cm.Del(key)
}

//用户邮件
func (cm *cacheManager) GetUserEmails(uid int) []structure.YlyPost {
	key := cm.CacheKey("GetUserEmails", uid)
	val := cm.Get(key)

	var info []structure.YlyPost
	if val == "" {
		info = models.GetUserMailByToid(uid)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	return info
}
func (cm *cacheManager) RemoveUserEmails(uid int) bool {
	key := cm.CacheKey("GetUserEmails", uid)
	return cm.Del(key)
}

//技能管理
func (cm *cacheManager) GetGameSkills() map[int]structure.GameSkills {
	key := cm.CacheKey("GetGameSkills")
	val := cm.Get(key)
	var info []structure.GameSkills
	if val == "" {
		info = models.GetGameSkills()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]structure.GameSkills)
	for _, val := range info {
		list[val.Skid] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameSkills() bool {
	key := cm.CacheKey("GetGameSkills")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameRoles() map[int]structure.GameRoles {
	key := cm.CacheKey("GetGameRoles")
	val := cm.Get(key)
	var info []structure.GameRoles
	if val == "" {
		info = models.GetGameRoles()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]structure.GameRoles)
	for _, val := range info {
		list[val.Roid] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameRoles() bool {
	key := cm.CacheKey("GetGameRoles")
	cm.Del(key)
	keys := cm.GetPatternKeys("GetUserRoles")
	return cm.DelKeys(keys...)
}

func (cm *cacheManager) GetGameShopConfig(args ...int) map[int]structure.GameShopConfig {
	channel := 0
	if len(args) > 0 {
		channel = args[0]
	}
	key := cm.CacheKey("GetGameShopConfig", channel)
	val := cm.Get(key)
	var info []structure.GameShopConfig
	if val == "" {
		info = models.GetGameShopConfig()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]structure.GameShopConfig)
	for _, val := range info {
		list[val.Scid] = val
	}
	return list
}

func (cm *cacheManager) GetMatchGradeSeason() int {
	key := cm.CacheKey("GetMatchGradeSeason")
	val := cm.Get(key)
	var season int
	if val == "" {
		season = models.GetMatchGradeSeason()
		cm.Set(key, season, Week)
	} else {
		season = utils.StringToInt(val)
	}
	return season
}

func (cm *cacheManager) RemoveMatchGradeSeason() bool {
	key := cm.CacheKey("GetMatchGradeSeason")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameRoleExps() map[int]map[int]structure.GameRoleExps {
	key := cm.CacheKey("GetGameRoleExps")
	val := cm.Get(key)
	var info []structure.GameRoleExps
	if val == "" {
		info = models.GetGameRoleExps()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]map[int]structure.GameRoleExps)
	for _, val := range info {
		if _, ok := list[val.Quality]; !ok {
			list[val.Quality] = make(map[int]structure.GameRoleExps)
		}
		list[val.Quality][val.RoleLevel] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameRoleExps() bool {
	key := cm.CacheKey("GetGameRoleExps")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameRoleSkills() map[int]map[int]structure.GameRoleSkills {
	key := cm.CacheKey("GetGameRoleSkills")
	val := cm.Get(key)
	var info []structure.GameRoleSkills
	if val == "" {
		info = models.GetGameRoleSkills()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]map[int]structure.GameRoleSkills)
	for _, val := range info {
		if _, ok := list[val.Quality]; !ok {
			list[val.Quality] = make(map[int]structure.GameRoleSkills)
		}
		list[val.Quality][val.SkillLevel] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameRoleSkills() bool {
	key := cm.CacheKey("GetGameRoleSkills")
	return cm.Del(key)
}

//赛季结算奖励管理
func (cm *cacheManager) GetMatchSeasonAwards(args ...int) map[int]map[int]structure.MatchSeasonAwards {
	season := 0
	if len(args) > 0 {
		season = args[0]
	}
	key := cm.CacheKey("GetMatchSeasonAwards", season)
	val := cm.Get(key)
	var info []structure.MatchSeasonAwards
	if val == "" {
		info = models.GetMatchSeasonAwards(season)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]map[int]structure.MatchSeasonAwards)
	for _, val := range info {
		if _, ok := list[val.Season]; !ok {
			list[val.Season] = make(map[int]structure.MatchSeasonAwards)
		}
		list[val.Season][val.Aid] = val
	}
	return list
}
func (cm *cacheManager) RemoveMatchSeasonAwards() bool {
	season := cm.GetMatchGradeSeason()
	for i := 0; i <= season; i++ {
		key := cm.CacheKey("GetMatchSeasonAwards", i)
		cm.Del(key)
	}
	return true
}
func (cm *cacheManager) GetGameBoxes(args ...int) map[int]structure.GameBoxes {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	key := cm.CacheKey("GetGameBoxes", state)
	val := cm.Get(key)
	var info []structure.GameBoxes
	if val == "" {
		info = models.GetGameBoxes(state)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}

	list := make(map[int]structure.GameBoxes)
	for _, val := range info {
		list[val.Bid] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameBoxes() bool {
	for i := 0; i <= 2; i++ {
		key := cm.CacheKey("GetGameBoxes", i)
		cm.Del(key)
	}
	keys := cm.GetPatternKeys("GetUserBoxes")
	return cm.DelKeys(keys...)
}

func (cm *cacheManager) GetUserBoxes(uid int, args ...int) map[int]structure.UserBoxesRelated {
	state := 0
	if len(args) > 0 {
		state = args[0]
	}
	key := cm.CacheKey("GetUserBoxes", uid, state)
	val := cm.Get(key)
	var info []structure.UserBoxesRelated
	ret := make(map[int]structure.UserBoxesRelated)
	if val == "" {
		info = models.GetUserBoxes(uid, state)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Ubid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Ubid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveUserBoxes(uid int) bool {
	for i := 0; i <= 2; i++ {
		key := cm.CacheKey("GetUserBoxes", uid, i)
		cm.Del(key)
	}
	return true
}

func (cm *cacheManager) GetOnlineNum(keys string, date int, start int, end int) []structure.LogOnlineNum {
	key := cm.CacheKey(keys)
	val := cm.Get(key)
	var info []structure.LogOnlineNum
	if val == "" {
		info = models.GetOnlineNum(date, start, end)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	return info
}
func (cm *cacheManager) GetForbidDeviceuid() map[string]structure.YlyForbidDeviceuid {
	key := cm.CacheKey("GetForbidDeviceuid")
	val := cm.Get(key)
	var info []structure.YlyForbidDeviceuid
	if val == "" {
		info = models.GetForbidDeviceuid()

		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[string]structure.YlyForbidDeviceuid)
	for _, val := range info {
		list[val.Deviceuid] = val
	}
	return list
}
func (cm *cacheManager) RemoveForbidDeviceuid() bool {
	key := cm.CacheKey("GetForbidDeviceuid")
	return cm.Del(key)
}
func (cm *cacheManager) GetGameRoleVoices() map[int]structure.GameRoleVoices {
	key := cm.CacheKey("GetGameRoleVoices")
	val := cm.Get(key)
	var info []structure.GameRoleVoices
	ret := make(map[int]structure.GameRoleVoices)
	if val == "" {
		info = models.GetGameRoleVoices()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		for _, v := range info {
			ret[v.Rvid] = v
		}
		return ret
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	for _, v := range info {
		ret[v.Rvid] = v
	}
	return ret
}
func (cm *cacheManager) RemoveGameRoleVoices() bool {
	key := cm.CacheKey("GetGameRoleVoices")
	cm.Del(key)

	keys := cm.GetPatternKeys("GetUserRoleVoices")
	return cm.DelKeys(keys...)
}

func (cm *cacheManager) GetUserRoleVoices(roid, skillLevel int) []structure.GameRoleVoices {
	key := cm.CacheKey("GetUserRoleVoices", roid, skillLevel)
	val := cm.Get(key)
	var info []structure.GameRoleVoices
	if val == "" {
		info = models.GetUserRoleVoices(roid, skillLevel)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
		} else {
			cm.Set(key, val, Day)
		}
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		logger.Println(err)
	}
	return info
}

func (cm *cacheManager) RemoveUserRoleVoices(roid, skillLevel int) bool {
	key := cm.CacheKey("GetUserRoleVoices", roid, skillLevel)
	return cm.Del(key)
}

func (cm *cacheManager) GetGameBlockIP() map[string]structure.GameBlockIp {
	key := cm.CacheKey("GetGameBlockIP")
	val := cm.Get(key)
	var info []structure.GameBlockIp
	if val == "" {
		info = models.GetGameBlockIP()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[string]structure.GameBlockIp)
	for _, val := range info {
		list[val.Ip] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameBlockIP() bool {
	key := cm.CacheKey("GetGameBlockIP")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameWinStreak() map[int]structure.GameWinStreak {
	key := cm.CacheKey("GetGameWinStreak")
	val := cm.Get(key)
	var info []structure.GameWinStreak
	if val == "" {
		info = models.GetGameWinStreak()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameWinStreak)
	for _, val := range info {
		list[val.Wins] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameWinStreak() bool {
	key := cm.CacheKey("GetGameWinStreak")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameTask(ttype int) map[int]structure.GameTask {
	key := cm.CacheKey("GetGameTask", ttype)
	val := cm.Get(key)
	var info []structure.GameTask
	if val == "" {
		info = models.GetGameTask(ttype)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameTask)
	for _, val := range info {
		list[val.Tkid] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameTask() bool {
	for i := 0; i <= 2; i++ {
		key := cm.CacheKey("GetGameTask", i)
		cm.Del(key)
	}
	return true
}

func (cm *cacheManager) GetGameLiveness() map[int]map[int]structure.GameLiveness {
	key := cm.CacheKey("GetGameLiveness")
	val := cm.Get(key)
	var info []structure.GameLiveness
	if val == "" {
		info = models.GetGameLiveness()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]map[int]structure.GameLiveness)
	for _, val := range info {
		if _, ok := list[val.Livetype]; !ok {
			list[val.Livetype] = make(map[int]structure.GameLiveness)
		}
		list[val.Livetype][val.Liveness] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameLiveness() bool {
	key := cm.CacheKey("GetGameLiveness")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameRatioFactor() map[int]structure.GameRatioFactor {
	key := cm.CacheKey("GetGameRatioFactor")
	val := cm.Get(key)
	var info []structure.GameRatioFactor
	if val == "" {
		info = models.GetGameRatioFactor()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]structure.GameRatioFactor)
	for _, val := range info {
		list[val.Ratio] = val
	}
	return list
}
func (cm *cacheManager) RemoveGameRatioFactor() bool {
	key := cm.CacheKey("GetGameRatioFactor")
	return cm.Del(key)
}

func (cm *cacheManager) GetGameVipLevel() map[int]*structure.GameVipLevel {
	key := cm.CacheKey("GetGameVipLevel")
	val := cm.Get(key)
	var info []*structure.GameVipLevel
	if val == "" {
		info = models.GetGameVipLevel()
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]*structure.GameVipLevel)
	for _, v := range info {
		list[v.VipLevel] = v
	}
	return list
}

func (cm *cacheManager) RemoveGameVipLevel() bool {
	key := cm.CacheKey("GetGameVipLevel")
	return cm.Del(key)
}

func (cm *cacheManager) GetLogVipLevel(uid int) map[int]*structure.LogVipLevel {
	key := cm.CacheKey("GetLogVipLevel", uid)
	val := cm.Get(key)
	var info []*structure.LogVipLevel
	if val == "" {
		info = models.GetLogVipLevel(uid)
		val, err := json.Marshal(info)
		if err != nil {
			logger.Println(err)
			return nil
		}
		cm.Set(key, val, Day)
	} else {
		err := json.Unmarshal([]byte(val), &info)
		if err != nil {
			logger.Println(err)
			return nil
		}
	}
	list := make(map[int]*structure.LogVipLevel)
	for _, v := range info {
		list[v.Nlevel] = v
	}
	return list
}

func (cm *cacheManager) RemoveLogVipLevel(uid int) bool {
	key := cm.CacheKey("GetLogVipLevel", uid)
	return cm.Del(key)
}
