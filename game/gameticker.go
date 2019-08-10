package game

import (
	"fmt"
	"github.com/jmesyan/nano/serialize"
	"github.com/jmesyan/nano/serialize/protobuf"
	"github.com/jmesyan/nano/utils"
	"reflect"
)

var (
	TickHandler *GameTicker
)

type GameTicker struct {
	ticker     int32
	serializer serialize.Serializer
	callbacks  map[int32]reflect.Value
}

func NewGameTicker() *GameTicker {
	return &GameTicker{
		ticker:     0,
		serializer: protobuf.NewSerializer(),
		callbacks:  make(map[int32]reflect.Value),
	}
}

func (this *GameTicker) GetTick(callback reflect.Value) int32 {
	if callback.Kind() == reflect.Func {
		if this.ticker > 65535 {
			this.ticker = 0
		}
		this.ticker++
		this.callbacks[this.ticker] = callback
		return this.ticker
	}
	return 0
}

func (this *GameTicker) ExecTick(tick int32, data reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			logger.Println(fmt.Sprintf("gameticker err: %v", err))
			println(utils.Stack())
		}
	}()

	if callback, ok := this.callbacks[tick]; ok {
		callback.Call([]reflect.Value{data})
	}
}

func init() {
	TickHandler = NewGameTicker()
}
