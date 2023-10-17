/*2. **Iterating Over a Map:**
Write a program that iterates over a map of cities and their populations, and prints the
city names with populations over one million.
*/
package main

import "fmt"

func main() {
	city := map[string]int64{
		"new york":  10023452,
		"kathmandu": 986731,
		"beijing":   765834211,
		"delhi":     372659819,
		"dhaka":     246781,
	}
	for cityName, popln := range city {
		if popln > 1000000 {
			fmt.Println(cityName, " : ", popln)
		}
	}

}
