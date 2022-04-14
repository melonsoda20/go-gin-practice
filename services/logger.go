package services

import (
	"errors"
	"log"
)

func LogError(err error) {
	log.Println(err)
}

func LogErrorMessage(msg string) {
	log.Println(errors.New(msg))
}
