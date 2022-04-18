package services

import "encoding/json"

func DeserializeFile(data []byte, i interface{}) error {
	return json.Unmarshal(data, &i)
}
