package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data any, code int) {
	w.Header().Set("Content-Type", "application/jsone")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
