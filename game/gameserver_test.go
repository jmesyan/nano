package game

import (
	"fmt"
	"testing"
)

func send(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
func TestCallback(t *testing.T) {
	s := send("welcome%d", 12)
	fmt.Println(s)
}
