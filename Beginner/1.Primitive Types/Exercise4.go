/***Exercise 4: Constants**

Declare a constant in Go to represent the value of pi (3.14159265), and use
it to calculate the circumference of a circle with a given radius.*/

package main

import "fmt"

func main() {

	const pi float64 = 3.14159265
	var radius float64
	fmt.Print("Enter the radius: ")
	fmt.Scanln(&radius)

	fmt.Println("The circumference of a circle for the given radius is ", 2*pi*radius)
}
