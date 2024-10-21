package output

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func WriteBinary(w http.ResponseWriter, statusCode int, payload []byte) {
	w.Header().Add("Content-Type", "application/octet-stream")
	w.WriteHeader(statusCode)
	w.Write(payload)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, err)
}
