package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("echo", "Hello, World!")
	output, err := cmd.Output()

	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
