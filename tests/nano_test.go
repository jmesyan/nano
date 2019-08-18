package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ExportStruct struct {
	Good   string `json:"good"`
	friend string `json:"friend"`
}

func TestGood(t *testing.T) {
	s := &ExportStruct{
		Good:   "1111",
		friend: "2222",
	}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", string(data))
}
