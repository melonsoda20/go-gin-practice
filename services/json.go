package services

import "encoding/json"

func DeserializeFile(data []byte, i interface{}) {
	json.Unmarshal(data, &i)
}
