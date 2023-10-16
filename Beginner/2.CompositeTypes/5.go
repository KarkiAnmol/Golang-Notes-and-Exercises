/*
 **Problem 5: Find the Smallest Element**
Write a function that takes a slice of integers and returns the smallest element in
the slice. If the slice is empty, return an error or a suitable message.
*/


package main

import (
	"errors"
	"fmt"
)

// smallest function takes a slice of integers and returns the smallest element in the slice.
// If the slice is empty, it returns an error.
func smallest(x []int) (int, error) {
	// Check if the slice is empty
	if len(x) == 0 {
		return 0, errors.New("ERROR: empty slice")
	}

	// Initialize the result with the first element
	result := x[0]

	// Iterate through the slice to find the smallest element
	for i := 0; i < len(x); i++ {
		if x[i] < result {
			result = x[i]
		}
	}

	// Return the smallest element and no error
	return result, nil
}

func main() {
	// Sample slice of integers
	x := []int{4, 8, 11, -22, 2, 3, 5}
	// To test with an empty slice, comment out the above line and uncomment the next line:
	// var x []int // empty slice

	// Call the smallest function to find the smallest element and check for errors
	smallestElement, err := smallest(x)

	// Check if there's an error
	if err != nil {
		fmt.Println(err)
	} else {
		// Print the smallest element if there are no errors
		fmt.Println("smallest element is", smallestElement)
	}
}
