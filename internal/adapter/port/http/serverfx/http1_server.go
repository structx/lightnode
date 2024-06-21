package serverfx

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// NewHttp1Server
func InvokeHTTPServer(mux *http.ServeMux) {
	http.ListenAndServe(":8080", handlers.CompressHandler(handlers.CORS()(mux)))
}
