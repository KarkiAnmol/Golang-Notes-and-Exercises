package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// result is a structure to store the result of an HTTP request, including URL, error, and latency.
type result struct {
	url     string
	err     error
	latency time.Duration
}

// get does an HTTP request and sends the result to a channel.
func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	// Make the HTTP request and handle errors.
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		// Calculate the latency and send the result to the channel.
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result) // Channel to receive results.
	list := []string{            // List of URLs to make HTTP requests.
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8080",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Iterate over the list of URLs and start a goroutine for each HTTP request.
	for _, url := range list {
		go get(ctx, url, results)
	}

	// After starting goroutines, wait for results.
	for range list {
		r := <-results
		if r.err != nil {
			// Log an error if there was an issue with the HTTP request.
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			// Log the URL and the latency if the request was successful.
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
