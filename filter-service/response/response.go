package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
