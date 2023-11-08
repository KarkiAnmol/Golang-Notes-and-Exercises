package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler handles requests made to the root URL ("/").
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HOME PAGE")
}

// formHandler processes form data sent via POST request to the "/form" endpoint.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	if err := r.ParseForm(); err != nil {
		// If there's an error parsing the form, it writes an error message to the response.
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Printf("POST request was successful")

	// Retrieving values from the form data.
	name := r.FormValue("name")
	age := r.FormValue("age")

	// Printing the received form values.
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Age: %s\n", age)
}

// helloHandler handles requests to the "/hello" endpoint with specific checks.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// If the URL path is not "/hello", respond with a 404 error.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// If the request method is not GET, respond with a 405 error.
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Respond with a simple "HELLO WORLD" message.
	fmt.Fprintf(w, "HELLO WORLD")
}

func main() {
	// Setting up handlers for different routes.
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)            // Handler for the root URL
	http.HandleFunc("/form", formHandler)   // Handler for the "/form" endpoint
	http.HandleFunc("/hello", helloHandler) // Handler for the "/hello" endpoint
	fmt.Println("Starting Server at port 8084")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
	}
}
