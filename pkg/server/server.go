package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Config struct {
	Addr         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type HTTPListener interface {
	ListenAndServe() error
}

type Server struct {
	server *http.Server
}

func (s *Server) ListenAndServe() error {
	log.Infof("Server ready and listen %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func NewHttpServer(handler http.Handler, cfg Config) *Server {
	return &Server{
		server: &http.Server{
			Handler:      handler,
			Addr:         cfg.Addr,
			ReadTimeout:  cfg.ReadTimeOut,
			WriteTimeout: cfg.WriteTimeOut,
		},
	}
}
