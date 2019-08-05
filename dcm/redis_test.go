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
