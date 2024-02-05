package utils

import (
	"encoding/json"
)

func Merge(des interface{}, src interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, des)
}
