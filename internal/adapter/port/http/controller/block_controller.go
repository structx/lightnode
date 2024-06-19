// Package controller endpoint http endpoints
package controller

import (
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"go.uber.org/zap"

	pkgcontroller "github.com/structx/go-dpkg/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/service"
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

	rr.Get(blockHashPath, bc.FetchByHash)
	rr.Get("/", bc.PaginatePartials)

	r.Mount(blockPath, rr)
}

// BlockPayload
type BlockPayload struct {
	Hash          string `json:"hash"`
	PrevHash      string `json:"prev_hash"`
	Timestamp     string `json:"timestamp"`
	Height        int    `json:"height"`
	AccessCtrlRef string `json:"access_ctrl_ref"`
	AccessHash    string `json:"access_hash"`
}

// FetchByHashResponse
type FetchByHashResponse struct {
	Payload *BlockPayload `json:"payload"`
	Elapsed int64         `json:"elapsed"`
}

// NewFetchByHashResponse
func NewFetchByHashResponse(block *domain.Block, start time.Time) *FetchByHashResponse {
	return &FetchByHashResponse{
		Payload: &BlockPayload{
			Hash:          hex.EncodeToString(block.Hash),
			PrevHash:      hex.EncodeToString(block.PrevHash),
			Timestamp:     block.Timestamp,
			Height:        block.Height,
			AccessCtrlRef: block.AccessCtrlRef,
			AccessHash:    block.AccessHash,
		},
		Elapsed: time.Since(start).Milliseconds(),
	}
}

// Render
func (fr *FetchByHashResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusAccepted)
	return nil
}

// FetchByHash
func (bc *Blocks) FetchByHash(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	ctx := r.Context()
	hashStr := chi.URLParamFromCtx(ctx, "blockHash")

	bc.log.Debugw("FetchByHash", "hash", hashStr)

	block, err := bc.service.ReadBlockByHash(ctx, hashStr)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			render.Render(w, r, pkgcontroller.ErrNotFound)
			return
		}

		bc.log.Errorf("failed to query block by hash %v", err)
		render.Render(w, r, pkgcontroller.ErrInternalServerError)
		return
	}

	response := NewFetchByHashResponse(block, start)
	_ = render.Render(w, r, response)
}

// BlockPartial
type BlockPartial struct {
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
	Height    int    `json:"height"`
	Timestamp string `json:"timestamp"`
}

// PaginateParialsResponse
type PaginatePartialsResponse struct {
	Payload []*BlockPartial `json:"payload"`
	Elapsed int64           `json:"elapsed"`
}

// NewPaginatePartialsResponse
func NewPaginatePartialsResponse(s []*domain.PartialBlock, start time.Time) *PaginatePartialsResponse {

	bs := make([]*BlockPartial, 0, len(s))

	for _, b := range s {
		bs = append(bs, &BlockPartial{
			Hash:      hex.EncodeToString(b.Hash),
			PrevHash:  hex.EncodeToString(b.PrevHash),
			Height:    b.Height,
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
	w.Header().Add("Content-Type", "application/json")
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
		bc.log.Errorf("failed to paginate blocks %v", err)
		_ = render.Render(w, r, pkgcontroller.ErrInternalServerError)
		return
	}

	_ = render.Render(w, r, NewPaginatePartialsResponse(blockSlice, start))
}
