package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"go.uber.org/zap"

	"github.com/structx/lightnode/internal/core/domain"
)

// Transactions controller
type Transactions struct {
	log *zap.SugaredLogger
	ss  domain.SimpleService
}

// NewTransactions
func NewTransactions(logger *zap.Logger, simpleService domain.SimpleService) *Transactions {
	return &Transactions{
		log: logger.Sugar().Named("TransactionsController"),
		ss:  simpleService,
	}
}

// RegisterRoutesV1 build controller handler from exposed endpoints
func (tx *Transactions) RegisterRoutesV1(r chi.Router) {

	rr := chi.NewRouter()

	rr.Get("/{txHash}", tx.Fetch)
	rr.Get("/", tx.Paginate)

	r.Mount("/transactions", rr)
}

// TxPayload
type TxPayload struct {
	ID string `json:"id"`
}

// FetchTxResponse
type FetchTxResponse struct {
	Payload *TxPayload `json:"payload"`
	Elapsed int64      `json:"elapsed"`
}

// NewFetchTxResponse
func NewFetchTxResponse(tx *domain.Transaction) *FetchByHashResponse {
	return &FetchByHashResponse{
		Payload: &BlockPayload{},
	}
}

// Fetch
func (txc *Transactions) Fetch(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// implement handler
	w.WriteHeader(http.StatusInternalServerError)
}

// Paginate
func (tx *Transactions) Paginate(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// implement handler
	w.WriteHeader(http.StatusInternalServerError)
}
