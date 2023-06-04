package main

import (
	"go_bank/app"
	"go_bank/logger"
)

func main() {
	// log.Println("Starting web service...")
	logger.Info("Starting web service...")
	app.Start()
}
