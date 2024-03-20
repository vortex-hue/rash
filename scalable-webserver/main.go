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

func userHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name string
		Age  int
		}

	var userOne User = User{"John", 25}

	time.Sleep(5 * time.Second)

	fmt.Fprintf(w, "User: %s, Age: %d", userOne.Name, userOne.Age)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", userHandler)
	fmt.Println("Server starting on port :8080...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
