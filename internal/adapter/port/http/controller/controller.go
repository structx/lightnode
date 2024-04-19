package controller

import "net/http"

type V0 interface {
	RegisterRoutesV0() http.Handler
}

type V1 interface {
	RegisterRoutesV1() http.Handler
}
