package services

import "os"

func OpenFile(filename string) (*os.File, error) {
	jsonFile, err := os.Open(filename)

	return jsonFile, err
}
