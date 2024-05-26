package controller

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"go.uber.org/zap"

	pkgcontroller "github.com/structx/go-dpkg/adapter/port/http/controller"
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

	r.Mount("/blocks/{blockHash}/transactions", rr)
}

// TxPayload
type TxPayload struct {
	ID            string   `json:"id"`
	Type          string   `json:"type"`
	Sender        string   `json:"sender"`
	Receiver      string   `json:"receiver"`
	Data          []byte   `json:"data"`
	Timestamp     string   `json:"timestamp"`
	Signatures    []string `json:"signatures"`
	AccessCtrlRef string   `json:"access_ctrl_ref"`
}

// FetchTxResponse
type FetchTxResponse struct {
	Payload *TxPayload `json:"payload"`
	Elapsed int64      `json:"elapsed"`
}

// NewFetchTxResponse
func NewFetchTxResponse(tx *domain.Transaction, start time.Time) *FetchTxResponse {
	return &FetchTxResponse{
		Payload: &TxPayload{
			ID:            hex.EncodeToString(tx.ID),
			Type:          tx.Type,
			Sender:        tx.Sender,
			Receiver:      tx.Receiver,
			Data:          tx.Data,
			Timestamp:     tx.Timestamp,
			Signatures:    tx.Signatures,
			AccessCtrlRef: tx.AccessCtrlRef,
		},
		Elapsed: time.Since(start).Milliseconds(),
	}
}

// Render
func (ftr *FetchTxResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusAccepted)
	return json.NewEncoder(w).Encode(ftr)
}

// Fetch
func (txc *Transactions) Fetch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	start := time.Now()

	bs := chi.URLParamFromCtx(ctx, "blockHash")
	ts := chi.URLParamFromCtx(ctx, "txHash")

	bh, err := hex.DecodeString(bs)
	if err != nil {
		_ = render.Render(w, r, pkgcontroller.ErrInvalidRequest(err))
		return
	}

	th, err := hex.DecodeString(ts)
	if err != nil {
		_ = render.Render(w, r, pkgcontroller.ErrInvalidRequest(err))
		return
	}

	tx, err := txc.ss.ReadTxByHash(ctx, bh, th)
	if err != nil {
		txc.log.Errorf("failed to read tx by hash %v", err)
		_ = render.Render(w, r, pkgcontroller.ErrInternalServerError)
		return
	}

	_ = render.Render(w, r, NewFetchTxResponse(tx, start))
}

// PartialTransaction
type PartialTransaction struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Timestamp string `json:"timestamp"`
}

// PaginateTransactionsResponse
type PaginateTransactionsResponse struct {
	Partials []*PartialTransaction `json:"partials"`
	Elapsed  int64                 `json:"elapsed"`
}

// NewPaginateTransactionsResponse
func NewPaginateTransactionsResponse(s []*domain.PartialTransaction, start time.Time) *PaginateTransactionsResponse {

	txs := make([]*PartialTransaction, 0, len(s))

	for _, t := range s {
		txs = append(txs, &PartialTransaction{
			ID:        hex.EncodeToString(t.ID),
			Type:      t.Type,
			Sender:    t.Sender,
			Receiver:  t.Receiver,
			Timestamp: t.Timestamp,
		})
	}

	return &PaginateTransactionsResponse{
		Partials: txs,
		Elapsed:  time.Since(start).Milliseconds(),
	}
}

// Render
func (ptr *PaginateTransactionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusAccepted)
	return json.NewEncoder(w).Encode(ptr)
}

// Paginate
func (tx *Transactions) Paginate(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	start := time.Now()

	blockHash := chi.URLParamFromCtx(ctx, "blockHash")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.ParseInt(limitStr, base, bitSize)
	if err != nil {
		_ = render.Render(w, r, pkgcontroller.ErrInvalidRequest(err))
		return
	}

	offset, err := strconv.ParseInt(offsetStr, base, bitSize)
	if err != nil {
		_ = render.Render(w, r, pkgcontroller.ErrInvalidRequest(err))
		return
	}

	txs, err := tx.ss.PaginateTransactions(ctx, []byte(blockHash), limit, offset)
	if err != nil {
		tx.log.Errorf("failed to paginate transactions %v", err)
		_ = render.Render(w, r, pkgcontroller.ErrInternalServerError)
		return
	}

	_ = render.Render(w, r, NewPaginateTransactionsResponse(txs, start))
}
