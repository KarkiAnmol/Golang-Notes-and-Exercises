/**Problem 4: Reverse a Slice**
Create a function that takes a slice of integers as an argument and returns a new
slice with the elements reversed. For example, if the input is `[1, 2, 3, 4]`, the output should be `[4, 3, 2, 1]`.*/

package main

import "fmt"

func revSlice(x []int, j int) []int {
	result := make([]int, len(x))
	z := len(x)
	for i := 0; i < j; i++ {
		z = z - 1
		result[i] = x[z]

	}
	return result

}

func main() {
	x := []int{1, 2, 3, 4} //input slice

	y := revSlice(x, len(x)) //output slice

	fmt.Println("reversed slice is ", y)

}
