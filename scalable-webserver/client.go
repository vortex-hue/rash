package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8000")
			if err != nil {
				fmt.Printf("Error on request %d: %s\n", i, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Request %d completed\n", i)
		}(i)
	}
	wg.Wait()
}
