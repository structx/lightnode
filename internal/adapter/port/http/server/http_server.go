// Package server http server provider
package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/trevatk/go-pkg/domain"
)

// New constructor
func New(cfg domain.Config, handler http.Handler) *http.Server {
	scfg := cfg.GetServer()
	return &http.Server{
		Addr:         net.JoinHostPort(scfg.BindAddr, fmt.Sprintf("%d", scfg.Ports.HTTP)),
		Handler:      handler,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
