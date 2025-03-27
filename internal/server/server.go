package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Cornpop456/go-6-sprint-final/internal/handlers"
)

type MorseServer struct {
	logger *log.Logger
	server http.Server
}

func (m *MorseServer) Start() error {
	return m.server.ListenAndServe()
}

func (m *MorseServer) Logger() *log.Logger {
	return m.logger
}

func New(logger *log.Logger) *MorseServer {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HandleIndex)
	mux.HandleFunc("POST /upload", handlers.HandleUpload)

	return &MorseServer{
		logger: logger,
		server: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
