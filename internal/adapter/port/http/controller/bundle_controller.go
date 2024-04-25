package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

// RegisterRoutesV0
func (b *Bundle) RegisterRoutesV0() http.Handler {

	r := chi.NewRouter()

	r.Get("/health", b.Health)

	return r
}

// Health
func (b *Bundle) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		b.log.Errorf("unable to write response %v", err)
		http.Error(w, "bad health check", http.StatusInternalServerError)
		return
	}
}
