// Package controller endpoint http endpoints
package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

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

// NewBlocks
func NewBlocks(logger *zap.Logger, simple domain.SimpleService) *Blocks {
	return &Blocks{
		log:     logger.Sugar().Named("BlockController"),
		service: simple,
	}
}

// RegisterRoutesV1
func (bc *Blocks) RegisterRoutesV1(mux *http.ServeMux) {
	mux.HandleFunc(blockHashPath, bc.FetchByHash)
	mux.HandleFunc(blockPath, bc.PaginatePartials)
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

func newFetchByHashResponse(block *domain.Block, start time.Time) *FetchByHashResponse {
	return &FetchByHashResponse{
		Payload: &BlockPayload{
			Hash:          block.Hash,
			PrevHash:      block.PrevHash,
			Timestamp:     block.Timestamp,
			Height:        block.Height,
			AccessCtrlRef: block.AccessCtrlRef,
			AccessHash:    block.AccessHash,
		},
		Elapsed: time.Since(start).Milliseconds(),
	}
}

// FetchByHash
func (bc *Blocks) FetchByHash(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	ctx := r.Context()
	hashStr := r.PathValue("blockHash")

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

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(newFetchByHashResponse(block, start))
	if err != nil {
		bc.log.Errorf("unable to code response %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
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

func newPaginatePartialsResponse(s []*domain.PartialBlock, start time.Time) *PaginatePartialsResponse {

	bs := make([]*BlockPartial, 0, len(s))

	for _, b := range s {
		bs = append(bs, &BlockPartial{
			Hash:      b.Hash,
			PrevHash:  b.PrevHash,
			Height:    b.Height,
			Timestamp: b.Timestamp,
		})
	}

	return &PaginatePartialsResponse{
		Payload: bs,
		Elapsed: time.Since(start).Milliseconds(),
	}
}

// PaginatePartials paginate partial blocks with limit and offset
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

	blockSlice, err := bc.service.PaginateBlocks(ctx, limit, offset)
	if err != nil {
		bc.log.Errorf("failed to paginate blocks %v", err)
		_ = render.Render(w, r, pkgcontroller.ErrInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(newPaginatePartialsResponse(blockSlice, start))
	if err != nil {
		bc.log.Errorf("unable to encode response %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
