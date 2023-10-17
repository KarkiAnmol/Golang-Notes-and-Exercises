/*4. **Map of Slices:**
Create a map where keys are string names of cities and values are slices of the top
three tourist attractions in each city. Print the attractions for a specific city.*/

package main

import "fmt"

func main() {
	attractionsMap := map[string][]string{
		"New York": []string{"Statue of Liberty", "Central Park", "Empire State Building"},
		"Paris":    []string{"Eiffel Tower", "Louvre Museum", "Notre-Dame Cathedral"},
		"Tokyo":    []string{"Tokyo Disneyland", "Senso-ji Temple", "Shinjuku Gyoen National Garden"},
	}
	fmt.Println("The top 3 attractions in paris are ", attractionsMap["Paris"])

}
