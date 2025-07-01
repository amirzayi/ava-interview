package jsonutil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return fmt.Errorf("failed to encode json: %w", err)
	}
	return nil
}
