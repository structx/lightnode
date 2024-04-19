package server

import (
	"net/http"
	"time"

	"github.com/trevatk/k2/internal/adapter/setup"
)

// New
func New(cfg *setup.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Server.ListenAddr,
		Handler:      handler,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
