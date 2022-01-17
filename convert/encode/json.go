package encode

import (
	"encoding/json"
	"log"
)

// JSON encode
func JSON(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(_json)
}
