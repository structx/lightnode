package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Transactions controller
type Transactions struct{}

func (tx *Transactions) RegisterRoutesV1() http.Handler {

	r := chi.NewRouter()

	return r
}
