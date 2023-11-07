package main

import (
	"encoding/json"
	"log"
)

// Person struct represents an individual with specific attributes.
type Person struct {
	Name       string `json:"name"`          // Name of the person
	Age        int    `json:"age,omitempty"` // Age of the person; omitempty avoids marshalling if left empty
	Sex        string `json:"sex"`           // Sex of the person
	Occupation string `json:"-"`             // Occupation of the person; '-' won't marshal this field
}

func main() {
	// Initializing a Person struct instance
	p := Person{
		Name:       "Anmol Karki",
		Age:        22,
		Sex:        "M",
		Occupation: "Student",
	}

	// Marshalling struct into JSON
	pBytes, err := json.Marshal(p) // Converts the struct into JSON bytes
	log.Print(err)                 // Log any error during marshalling
	log.Print(string(pBytes))      // Log the marshalled JSON string

	// Unmarshalling JSON into struct
	// RawMessage is a raw encoded JSON value.
	// It can be used to delay JSON decoding or precompute a JSON encoding.
	pAsRawJSON := json.RawMessage(`{"name":"Anmol Karki","age":22,"sex":"M"}`)
	var p2 Person
	err2 := json.Unmarshal(pAsRawJSON, &p2) // Converts JSON into the provided struct
	log.Print(err2)                         // Log any error during unmarshalling
	log.Printf("%+v", p2)                   // Log the unmarshalled struct values
}
