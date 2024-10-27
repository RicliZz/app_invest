package server

import (
	"context"
	"github.com/RicliZz/app_invest/config"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler http.Handler) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Addr,
			Handler: handler,
		},
	}
}

func (s *APIServer) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *APIServer) Stop() error {
	return s.httpServer.Shutdown(context.Background())
}
