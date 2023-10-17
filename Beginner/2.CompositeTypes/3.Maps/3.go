/*3. **Nested Maps:**
Build a nested map structure that represents a library catalog. The outer map
should use categories as keys, and the inner maps should contain books and their availability.
Perform operations to add, remove, and display books.
*/

package main

import "fmt"

func main() {
	libraryCatalog := map[string]map[string]bool{
		"fiction": {
			"harry potter":       true,
			"lord of the rings":  false,
			"apple and the worm": true,
		},
		"history": {
			"ancient rome":        false,
			"the italian warrior": true,
		},
	}
	fmt.Println(libraryCatalog)
	for book, _ := range libraryCatalog["fiction"] {
		if book == "lord of the rings" {
			fmt.Println(book, " deleted")
			delete(libraryCatalog["fiction"], book)
		}
	}
	fmt.Println(libraryCatalog)

}
