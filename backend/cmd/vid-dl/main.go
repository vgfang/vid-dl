package main

import (
	"fmt"
	"net/http"
	"strconv"

	"vid-dl/pkg/handlers"
	"vid-dl/pkg/services"
)

func main() {
	port := 3000
	fmt.Printf("Server Starting on Port %d...\n", port)

	services.InitDB()

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/file-options", handlers.FileOptions)
	http.HandleFunc("/download-video", handlers.DownloadVideo)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
