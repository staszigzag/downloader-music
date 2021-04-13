package server

import (
	"context"
	"net/http"

	"github.com/staszigzag/downloader-music/internal/config"
	"github.com/staszigzag/downloader-music/pkg/logger"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	logger.Info("Server is started...")
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	logger.Info("Server is stopped!")
	return nil
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(context.TODO())
}
