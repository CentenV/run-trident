package main

import (
	"github.com/centenv/run-trident/internal/logger"
)


func main() {
	var log = logger.GetInstance()
	
	log.Error("hello world")
	log.Info("hello world")
}
