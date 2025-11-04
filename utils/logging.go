package utils

import (
	"log"
	"time"
)

func Logging(p string) {
	log.Printf("%s -> p", time.Now())
}
