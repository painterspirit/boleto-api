package util

import (
	"encoding/json"
)

//FromJSON converts string json to object
func FromJSON(jsonStr string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), obj)
}

//ToJSON converts object to json string
func ToJSON(obj interface{}) string {
	s, _ := json.Marshal(obj)
	return string(s)
}
