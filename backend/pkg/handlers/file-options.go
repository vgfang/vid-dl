package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vid-dl/pkg/services"
	"vid-dl/pkg/utils"
)

func FileOptions(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "appplication/json")

	// check what video website they are trying to get to
	urltype := services.GetUrlType(url)
	if urltype == "invalid" {
		utils.WriteJSONMessage(w, 400, "invalid url")
		return
	} else if urltype == "unknown" {
		utils.WriteJSONMessage(w, 400, "unknown url")
		return
	}

	output, err := services.GetFileOptions(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(output)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON at file options: %v", err), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
