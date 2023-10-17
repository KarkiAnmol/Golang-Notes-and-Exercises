/**
   **Problem 1: Slice Basics**
Write a Go program that declares a slice of integers, appends values to it,
and prints the slice's length, capacity, and elements.
**/
package main

import "fmt"

func main() {
	var iSlis []int // the value of this slice is nil, nil is an identifier that represents the lack of value for some types ('int' in this case)
	fmt.Println("Before appending values \n Slice length: ", len(iSlis), " capacity: ", cap(iSlis), " elements: ", iSlis)

	iSlis = append(iSlis, 10, 20, 30)

	fmt.Println("\n -------------------------------------- \n After appending values \n Slice length: ", len(iSlis), " capacity: ", cap(iSlis), " elements: ", iSlis)
	iSlis = append(iSlis, 40)
	fmt.Println("\n After Appending again \n Slice length: ", len(iSlis), " capacity: ", cap(iSlis), " elements: ", iSlis)

}
