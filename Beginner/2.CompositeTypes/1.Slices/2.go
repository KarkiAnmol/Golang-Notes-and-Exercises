/*
**Problem 2: Slicing**
Create a program that takes a slice of integers, slices it to create a new slice,
and then prints the new slice. Experiment with different slices, including both simple slices and slices of slices.
*/

package main

import "fmt"

func slicer(x []int, i int) []int {
	y := x[:i]
	return y
}
func main() {

	var x = []int{1, 2, 3, 4, 5}

	z := slicer(x, 2)
	fmt.Println("x: ", x, "\n z: ", z)

}
