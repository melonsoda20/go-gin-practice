package services

import "encoding/json"

func DeserializeJSON(data []byte, i interface{}) error {
	return json.Unmarshal(data, &i)
}

func SerializeJson(i interface{}) ([]byte, error) {
	return json.Marshal(&i)
}
