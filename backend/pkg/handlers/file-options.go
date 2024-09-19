package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vid-dl/pkg/services"
)

func FileOptions(w http.ResponseWriter, r *http.Request) {
	// TODO: Make sure params are safe
	type requestBody struct {
		Url string `json:"url"`
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var body requestBody
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

	jsonData, err := json.Marshal(output)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON at file options: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "appplication/json")
	w.Write(jsonData)
}
