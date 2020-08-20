package main

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status bool `json:"status"`
}

func (h Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Health{true})
}
