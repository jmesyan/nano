package game

import (
	"fmt"
	"github.com/jmesyan/nano/utils"
	"reflect"
	"testing"
)

type Person struct {
	Name string
}

func As(a *Person) {
	fmt.Println(a.Name)
}
func pcall(method reflect.Method, args []reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			logger.Println(fmt.Sprintf("github.com/jmesyan/nano/dispatch: %v", err))
			println(utils.Stack())
		}
	}()

	if r := method.Func.Call(args); len(r) > 0 {
		if err := r[0].Interface(); err != nil {
			logger.Println(err.(error).Error())
		}
	}
}
func TestReflect(t *testing.T) {
	value := reflect.ValueOf(As)
	fmt.Printf("value is:%v\n", value.Kind() == reflect.Func)
	field := value.Type().In(0)
	fmt.Printf("field is:%#v\n", field)
	data := reflect.New(field.Elem()).Interface()
	svg := data.(*Person)
	svg.Name = "wecome"
	args := []reflect.Value{reflect.ValueOf(svg)}
	value.Call(args)
}

type He struct {
	Name string
}

func TestTicker(t *testing.T) {
	ab := func(p *Person) {
		fmt.Println(p.Name)
	}
	tick := TickHandler.GetTick(reflect.ValueOf(ab))
	p := &Person{Name: "hellworld"}
	TickHandler.ExecTick(tick, reflect.ValueOf(p))
}
