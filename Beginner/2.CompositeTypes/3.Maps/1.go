/*1. **Basic Map Operation:**
Create a map that stores the names of people as keys and their ages as values. Add a few entries, delete one, and print the map.*/

package main

import "fmt"

func main() {
	people := map[string]int{}
	people["jorge"] = 22
	people["john"] = 23
	people["mary"] = 19
	fmt.Println(people)

	v, ok := people["john"]
	if ok {
		fmt.Println("\nJohn of age ", v, "is being deleted...\n")
		delete(people, "john")
	}

	fmt.Println(people)
}
