/*
**Exercise 5: Temperature Conversion**

Develop a program that converts temperatures from Celsius to Fahrenheit. Declare a variable
 for the Celsius temperature, convert it to Fahrenheit, and print the result.*/

package main

import "fmt"

func main() {
	var tempC float64
	fmt.Print("enter the temperature in celcius")
	fmt.Scanln(&tempC)

	fmt.Println("The temperature in fahrenheit is ", (tempC*(9/5))+32)

}
