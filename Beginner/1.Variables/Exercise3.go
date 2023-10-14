/*
 **Exercise 3: User Input and Arithmetic**

Write a Go program that takes two integer inputs from the user, performs
 arithmetic operations (addition, subtraction, multiplication, division), and prints the results.
*/

package main

import "fmt"

func main() {
	var first int
	var second int
	fmt.Print("enter first number: ")
	fmt.Scanln(&first)
	fmt.Print("enter second number: ")
	fmt.Scanln(&second)

	fmt.Println("Addition: ", first+second)
	fmt.Println("Subtraction: ", first-second)
	fmt.Println("multiplication: ", first*second)
	fmt.Println("division: ", first/second)
}
