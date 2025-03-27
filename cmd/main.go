package main

import (
	"log"
	"os"

	"github.com/Cornpop456/go-6-sprint-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http", log.LstdFlags)

	morseServer := server.New(logger)

	err := morseServer.Start()

	if err != nil {
		morseServer.Logger().Fatal(err)
	}
}
