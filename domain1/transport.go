package domain1

import (
	"net/http"

	"github.com/acoshift/go-ddd-service-boilerplate/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

// NewHTTPTransport creates new HTTP transport for domain1
func NewHTTPTransport(ep Endpoint) http.Handler {
	mux := http.NewServeMux()

	errorEncoder := func(w http.ResponseWriter, err error) {
		status := http.StatusInternalServerError
		switch err {
		case ErrEntity1NotFound:
			status = http.StatusNotFound
		}
		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
	}

	mux.Handle("/create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		resp, err := ep.Create(r.Context(), &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, &resp)
	}))

	// or use https://github.com/acoshift/hrpc for RPC-HTTP style API
	// mux.Handle("/create", m.Handler(ep.Create))

	return mux
}
