package leetcode

import (
	"encoding/json"
	"fmt"
)

func PrintTest(i interface{}) {
	bytes, _ := json.Marshal(i)
	fmt.Println(string(bytes))
}
