package main

import (
	"fmt"
	"net/http"
	"time"
)
func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2025-03-30 15:04:05")
	fmt.Fprintf(w, "Current Date & Time: %s", currentTime)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
