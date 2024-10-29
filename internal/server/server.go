package server

import (
	"context"
	"github.com/RicliZz/app_invest/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(cfg config.Config, router *gin.Engine) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Addr,
			Handler: router,
		},
	}
}

func (s *APIServer) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal("Server not started")
	}
}

func (s *APIServer) Shutdown() {
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Fatal("Server not shutdown")
	}
}
