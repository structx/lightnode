package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Transactions controller
type Transactions struct{}

// RegisterRoutesV1 build controller handler from exposed endpoints
func (tx *Transactions) RegisterRoutesV1() http.Handler {

	r := chi.NewRouter()

	return r
}
