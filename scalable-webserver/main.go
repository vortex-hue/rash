package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCount int32

func handler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&requestCount, 1)
	currentCount := atomic.LoadInt32(&requestCount)

	// Simulate a compute-heavy operation
	time.Sleep(2 * time.Second)

	fmt.Fprintf(w, "Hello, you've hit the endpoint. Current request count: %d", currentCount)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
