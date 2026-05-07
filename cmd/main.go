package main

import (
	"kafka_http/internal/app"
	"kafka_http/internal/logger"
	"log"
)

func main() {
	logg, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(logg)
}
