package handlers

import (
	"encoding/json"
	"fmt"
)

const (
	SUCCESS_STAUTS = 1
	ERROR_STAUTS = 0
	SUCCESS_MSG = "success"
)


func DeepCopy(model interface{}, rpc interface{}) {

	b, _ := json.Marshal(model)

	fmt.Println(string(b))

	_ = json.Unmarshal(b, rpc)

}
