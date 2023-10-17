map[keyType]valueType

//var declaration to create a map variable with zero value 
var nilMap map[string]int //string keys ,int values
//Attempting to write a nilMap variable causes panic,read gives 0

//:= declaration to create a map variable by assigning it a map literal:
totalWins := map[string]int{} //In this case, we are using an empty map literal. This is not the same as a nil map. It
// has a length of 0, but you can read and write to a map assigned an empty map literal.

 //non empty map literal
 teams := map[string][]string {
	"Orcas": []string{"Fred", "Ralph", "Bijou"}, //values are a slice of string
	"Lions": []string{"Sarah", "Peter", "Billie"},// 
	"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	//when we know how many keyvalue to put in a map use MAKE
	ages := make(map[int][]string, 10)

//Key points on Maps
//--> Maps grow automatically as you add key-value pairs
//--> use Make when you know how many keyvalue to add
//-->len gives number of key-value pairs
//-->zero value for a map is nil
//-->maps are not comparable,can be compared with nil only 
//--> key can be any comparable type which means slices and maps cannot be used as a key


//Reading and Writing a Map
// Using a map
totalWins := map[string]int{}
totalWins["Orcas"] = 1
totalWins["Lions"] = 2
fmt.Println(totalWins["Orcas"])
fmt.Println(totalWins["Kittens"])
totalWins["Kittens"]++
fmt.Println(totalWins["Kittens"])
totalWins["Lions"] = 3
fmt.Println(totalWins["Lions"])

// When you run this program, you’ll see the following output:
1
0
1
3

//The comma ok Idiom
//Go provides the comma ok idiom to tell the difference between a key that’s associated
// with a zero value and a key that’s not in the map:

m := map[string]int{
	"hello": 5,
	"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"] 
	fmt.Println(v, ok) // v=0 , ok true,/.... i.e. key exists but value is 0
	v, ok = m["goodbye"]
	fmt.Println(v, ok) // v=0, ok=false ... i.e. key doesn't exist


//Deleting from Maps
// Key-value pairs are removed from a map via the built-in delete function:
m := map[string]int{
"hello": 5,
"world": 10,
}
delete(m, "hello")

