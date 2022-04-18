package services

import "encoding/json"

func DeserializeFile(data []byte, i interface{}) (interface{}, error) {
	err := json.Unmarshal(data, &i)

	return i, err
}
