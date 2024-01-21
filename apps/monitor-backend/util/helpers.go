package util

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error("Error encoding", err)
	}
}
