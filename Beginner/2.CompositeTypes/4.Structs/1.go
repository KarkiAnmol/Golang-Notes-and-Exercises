/*. **Create a Struct:**
Define a struct named "Person" with fields for a person's name, age, and address.
 Create instances of the struct and print their details. */
package main

import "fmt"

func main() {
	type Person struct {
		name    string
		age     int
		address string
	}
	bob := Person{
		name:    "Bob Marley",
		age:     49,
		address: "mars",
	}
	fmt.Printf("%s of age %d lives in %s \n", bob.name, bob.age, bob.address)
}
