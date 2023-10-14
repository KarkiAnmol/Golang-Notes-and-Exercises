/* **Exercise 6: Area of a Rectangle**

Create a Go program that calculates the area of a rectangle. Declare variables for
 the length and width, and then compute and print the area. */

package main

import "fmt"

func main() {

	var length, width int
	fmt.Print("enter the length: ")
	fmt.Scanln(&length)
	fmt.Print("enter the width: ")
	fmt.Scanln(&width)

	fmt.Println("The area is ", length*width)
}
