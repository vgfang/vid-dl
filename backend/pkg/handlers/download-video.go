package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vid-dl/pkg/services"
)

func DownloadVideo(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Url     string `json:"url"`
		VideoID string `json:"videoID"`
		AudioID string `json:"audioID"`
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
	videoID := body.VideoID
	audioID := body.AudioID
	defer r.Body.Close()

	output, err := services.DownloadVideo(url, videoID, audioID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
