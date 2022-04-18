package services

import (
	"io/ioutil"
	"os"
)

func ReadFile(jsonFile *os.File) ([]byte, error) {
	jsonData, err := ioutil.ReadAll(jsonFile)
	return jsonData, err
}
