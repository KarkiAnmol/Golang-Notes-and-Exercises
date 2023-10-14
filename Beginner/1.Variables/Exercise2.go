/***Exercise 2: Swapping Variables**
Create a Go program that swaps the values of two integer variables without using a temporary variable.
*/
package main

import "fmt"

func main() {
	var first int
	var second int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&first) //10
	fmt.Print("Enter second numebr: ")
	fmt.Scanln(&second) // 14

	first = first - second  // -4
	second = first + second // 10
	first = second - first

	fmt.Println("First Number: ", first)
	fmt.Println("Second Numeber: ", second)
}
