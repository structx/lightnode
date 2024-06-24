package controller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// Bundle
type Bundle struct {
	log *zap.SugaredLogger
}

// NewBundle
func NewBundle(logger *zap.Logger) *Bundle {
	return &Bundle{
		log: logger.Sugar().Named("BundleController"),
	}
}

// RegisterRootRoutes
func (bc *Bundle) RegisterRootRoutes(mux *http.ServeMux) {
	mux.HandleFunc(health, bc.Health)
	mux.Handle(metrics, promhttp.Handler())
}

// Health
func (bc *Bundle) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		bc.log.Errorf("unable to write response %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
