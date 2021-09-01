package common

import (
	"errors"
	"log"
)

var (
	RecordNotFound = errors.New("record not found")
)

// Use to recovery error from a go routine
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error: ", err)
	}
}
