package main

import (
	"fmt"
	"net/http"
	"strconv"

	"vid-dl/pkg/handlers"
)

func main() {
	port := 3000
	fmt.Printf("Server Starting on Port %d...\n", port)

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/test/", handlers.TestHandler)
	http.HandleFunc("/file-options", handlers.FileOptions)
	http.HandleFunc("/download-video", handlers.DownloadVideo)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
