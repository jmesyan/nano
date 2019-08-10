package game

import (
	"fmt"
	"github.com/jmesyan/nano/application/cache"
	"github.com/jmesyan/nano/application/stores"
	"testing"
)

func send(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
func TestCallback(t *testing.T) {
	s := send("welcome%d", 12)
	fmt.Println(s)
}

func TestStores(t *testing.T) {
	s := stores.StoresHandler.Gds
	fmt.Printf("list is:%#v, length is:%d\n", s.Configs, len(s.Configs))
	configs := cache.CacheManager.GetGameGoldsType()
	fmt.Printf("the configs is:%#v\n", configs)
}
