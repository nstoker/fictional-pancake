package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nstoker/fictional-pancake/internal/version"
)

type Health struct {
	OK      bool   `json:"ok"`
	Version string `json:"version"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	healthy := Health{}
	healthy.OK = true
	healthy.Version = version.Version()

	json.NewEncoder(w).Encode(healthy)
}
