package runtimehelper

import (
	"log"
)

func Recover(msg string) {
	if err := recover(); err != nil {
		log.Println(msg, err)
	}
}
