package handler

import (
	"encoding/json"
	"fmt"
)

func DeepCopy(model interface{}, rpc interface{}) {

	b, _ := json.Marshal(model)

	fmt.Println(string(b))

	_ = json.Unmarshal(b, rpc)

}
