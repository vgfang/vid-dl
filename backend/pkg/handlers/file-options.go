package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vid-dl/pkg/services"
)

type RequestBody struct {
	Url string `json:"url"`
}

func FileOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
		return
	}

	url := body.Url
	defer r.Body.Close()

	output, err := services.GetFileOptions(url)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
