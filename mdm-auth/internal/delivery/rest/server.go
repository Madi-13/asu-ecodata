package rest

import (
	"context"
	"mdm-auth/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Http.Port,
			Handler:        handler,
			ReadTimeout:    cfg.Http.ReadTimeout * time.Second,
			WriteTimeout:   cfg.Http.WriteTimeout * time.Second,
			MaxHeaderBytes: cfg.Http.MaxHeaderMegabytes << 20},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
