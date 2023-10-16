/* **Problem 3: Copy Slices**
Write a Go program that copies the elements from one slice to another. Ensure that
the original slice remains unchanged.
*/

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5} // creating a slice literal
	y := make([]int, 3)      //creating slice using make function and specifying the length to 3
	copy(y, x)               //copy x to y ,the values at position 0,1 and 2
	fmt.Println("x : ", x, len(x), cap(x))
	fmt.Println("y : ", y, len(y), cap(y))

	y = append(y, 9)
	fmt.Println("after appending in y ")
	fmt.Println("x : ", x, len(x), cap(x))
	fmt.Println("y : ", y, len(y), cap(y))

}
