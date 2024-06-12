package serverfx

import "net/http"

// NewHttp1Server
func NewHttp1Server(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}
