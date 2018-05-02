package httptransport

import (
	"encoding/json"
	"io"
	"net/http"
)

const maxJSONBodySize = 10 << 20 // 10 MiB

// DecodeJSON decodes json
func DecodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(io.LimitReader(r, maxJSONBodySize)).Decode(v)
}

// EncodeJSON encodes json
func EncodeJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
