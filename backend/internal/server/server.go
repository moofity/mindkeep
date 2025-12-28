package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/internal/config"
	"backend/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer(cfg *config.Config) *http.Server {
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	NewServer := &Server{
		port: cfg.Server.Port,
		db:   database.New(cfg),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	return server
}
