package utils

import "encoding/json"

func ToJSONStr(i interface{}) string {
	j, _ := json.Marshal(i)
	return string(j)
}
