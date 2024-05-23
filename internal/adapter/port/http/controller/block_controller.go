// Package controller endpoint http endpoints
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"go.uber.org/zap"

	pkgcontroller "github.com/structx/go-dpkg/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/core/domain"
)

// Blocks ...
type Blocks struct {
	log     *zap.SugaredLogger
	service domain.SimpleService
}

// interface compliance
var _ pkgcontroller.V1 = (*Blocks)(nil)

// NewBlocks
func NewBlocks(logger *zap.Logger, simple domain.SimpleService) *Blocks {
	return &Blocks{
		log:     logger.Sugar().Named("BlockController"),
		service: simple,
	}
}

// RegisterRoutesV1
func (bc *Blocks) RegisterRoutesV1(r chi.Router) {

	rr := chi.NewRouter()

	rr.Get("/{blockHash}", bc.FetchByHash)
	rr.Get("/", bc.PaginatePartials)

	r.Mount("/blocks", rr)
}

// BlockPayload
type BlockPayload struct{}

// FetchByHashResponse
type FetchByHashResponse struct {
	Payload *BlockPayload `json:"payload"`
	Elapsed int64         `json:"elapsed"`
}

// NewFetchByHashResponse
func NewFetchByHashResponse(block *domain.Block, start time.Time) *FetchByHashResponse {
	return &FetchByHashResponse{}
}

// Render
func (fr *FetchByHashResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusAccepted)
	err := json.NewEncoder(w).Encode(fr)
	if err != nil {
		return fmt.Errorf("failed to encode response %v", err)
	}
	return nil
}

// FetchByHash
func (bc *Blocks) FetchByHash(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	ctx := r.Context()
	hash := chi.URLParamFromCtx(ctx, "blockHash")

	block, err := bc.service.QueryBlockByHash(ctx, []byte(hash))
	if err != nil {

	}

	response := NewFetchByHashResponse(block, start)
	render.Render(w, r, response)
}

// BlockPartial
type BlockPartial struct {
	Hash      string `json:"hash"`
	Timestamp string `json:"timestamp"`
}

// PaginateParialsResponse
type PaginatePartialsResponse struct {
	Payload []*BlockPartial `json:"payload"`
	Elapsed int64           `json:"elapsed"`
}

// NewPaginatePartialsResponse
func NewPaginatePartialsResponse(s []*domain.Block, start time.Time) *PaginatePartialsResponse {

	bs := make([]*BlockPartial, 0, len(s))

	for _, b := range s {
		bs = append(bs, &BlockPartial{
			Hash:      b.Hash,
			Timestamp: b.Timestamp,
		})
	}

	return &PaginatePartialsResponse{
		Payload: bs,
		Elapsed: time.Since(start).Milliseconds(),
	}
}

// Render
func (ppr *PaginatePartialsResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(http.StatusAccepted)
	return nil
}

// PaginatePartials
func (bc *Blocks) PaginatePartials(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	start := time.Now()

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.ParseInt(limitStr, base, bitSize)
	if err != nil {
		return
	}

	offset, err := strconv.ParseInt(offsetStr, base, bitSize)
	if err != nil {
		return
	}

	bc.log.Debugf("PaginatePartials", "limit", limit, "offset", offset)

	blockSlice, err := bc.service.PaginateBlocks(ctx, limit, offset)
	if err != nil {
		return
	}

	response := NewPaginatePartialsResponse(blockSlice, start)
	_ = render.Render(w, r, response)
}
