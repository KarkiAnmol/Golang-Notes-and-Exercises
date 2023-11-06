package main

import (
	"io"
	"os"
	"strings"
)

//program that reads from an io.reader and writes to an io.writer
func main() {

	inputString := "Hello, this is a sample string for io.Reader and io.Writer.\n"

	// Creating an io.Reader from a string
	//strings.NewReader creates an io.Reader by returning a *strings.Reader type.
	//  The *strings.Reader type has a method named Read that matches the signature of the io.Reader interface.
	// This means *strings.Reader is an implementation of the io.Reader interface
	reader := strings.NewReader(inputString)

	// Using an io.Writer to write to the standard output (os.Stdout),os.Stdout is an instance of '*os.File'
	writer := os.Stdout

	// Reading from the io.Reader and writing to the io.Writer
	_, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}

}
