package dcm

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"testing"
)

func TestRedisCluster(t *testing.T) {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005", "127.0.0.1:7006"},
	})
	//redisdb.Ping()
	str, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
	val := redisdb.Get("name").Val()
	fmt.Println(val)
}

func TestTrimPrefix(t *testing.T) {
	key := "dcm:well"
	nkey := strings.TrimLeft(key, "dcm:")
	fmt.Println(nkey)
}

func TestUserData(t *testing.T) {
	uid, k, _ := 1000, "feng", map[string]interface{}{"sdsd": 12121}
	//err := DCManager.StoreUserData(uid, k, v)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = DCManager.DelUserData(uid, k)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	var value map[string]interface{}
	err := DCManager.GetUserData(uid, k, &value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", value)
}
