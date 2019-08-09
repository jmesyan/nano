package utils

import (
	"encoding/binary"
	"fmt"
	"sync"
	"testing"
	"time"
)

func null(key, value int) int {
	key += value
	return key
}

func test1(m map[int]int) {
	if len(m) < 1 {
		return
	}
	for k, v := range m {
		null(k, v)
	}
}

func BenchmarkEmptyMap1(b *testing.B) {
	b.ReportAllocs()

	m := map[int]int{}
	for i := 0; i < b.N; i++ {
		test1(m)
	}
}

func test2(m map[int]int) {
	for k, v := range m {
		null(k, v)
	}
}

func BenchmarkEmptyMap2(b *testing.B) {
	b.ReportAllocs()

	m := map[int]int{}
	for i := 0; i < b.N; i++ {
		test2(m)
	}
}

func TestUint32ToByte(t *testing.T) {
	b := make([]byte, 4)
	var v uint32
	v = 1688272828
	fmt.Println(b)
	binary.LittleEndian.PutUint32(b, v)
	fmt.Println(b)
}

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("welcome")
		wg.Done()
	})
	wg.Wait()
}

func TestFillZero(t *testing.T) {
	s := "00000000001234"
	fmt.Println(s[len(s)-10:])
}
