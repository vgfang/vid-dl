package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
)

func main() {
	port := 3000
	fmt.Printf("Server Starting on Port %d...\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
