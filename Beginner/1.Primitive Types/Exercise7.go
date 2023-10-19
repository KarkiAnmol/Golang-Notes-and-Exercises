/*
                     **Exercise 7: Age in Dog Years**

Write a Go program that takes a user's age as input and calculates and prints their age in dog years
 (multiply the age by 7).*/

package main

import "fmt"

func main() {
	var age int
	fmt.Print("Please enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Your age in dog years is ", age*7)
}
