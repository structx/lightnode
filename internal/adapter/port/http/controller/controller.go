package controller

import "net/http"

type V0 interface {
	RegisterRoutesV0() http.Handler
}

type V1 interface {
	RegisterRoutesV1() http.Handler
}

// V1P
type V1P interface {
	RegisterRoutesV1P() http.Handler
}
