package apiserver

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	src := "{\"cmd\":\"getGameListState\",\"timestamp\":1565678143}"
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(src), &obj)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(obj)
	}
}
