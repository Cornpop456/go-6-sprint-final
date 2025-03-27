package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Cornpop456/go-6-sprint-final/internal/handlers"
)

type MorseServer struct {
	Logger *log.Logger
	Server http.Server
}

func Create(logger *log.Logger) *MorseServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleIndex)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	return &MorseServer{
		Logger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
