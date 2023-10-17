/*3. **Struct Composition:**
Create a struct "Rectangle" with fields for width and height. Embed a "Location" struct within "Rectangle" to represent
the location of the rectangle's top-left corner. Display the full details of a rectangle, including its location.*/

package main

import "fmt"

func main() {
	type Location struct {
		X, Y int
	}
	type Rectangle struct {
		width, height int
		Location
	}

	r := Rectangle{
		width:  10,
		height: 4,
		Location: Location{
			X: 199,
			Y: 29,
		},
	}
	fmt.Printf("Rectangle Details:\n")
	fmt.Printf("Width: %d\n", r.width)
	fmt.Printf("Height: %d\n", r.height)
	fmt.Printf("Location (X, Y): (%d, %d)\n", r.X, r.Y)
}
