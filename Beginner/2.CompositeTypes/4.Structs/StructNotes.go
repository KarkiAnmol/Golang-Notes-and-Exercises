//NOTE
//Why not maps and Structs?
//->Limitaitons:
//--->There is no way to constrain a map to only allow certain keys
//--->all of the values in a map must be of the same type
//--->maps are not an ideal way to pass data from function to function	


//When you have related data that you want to group together, you should define a struct.
package main 
import "fmt"
type person struct {
	name string
	age int
	pet string
	}

// Once a struct type is declared, we can define variables of that type:
	var fred person

// A struct literal can be assigned to a variable as well:
	bob := person{} //initializes all of the fields in struct to their zero values

//declare and initialize
	julia := person{
		"Julia",
		40,
		"cat",
		}

// The second struct literal style looks like the map literal style:
		beth := person{
		age: 30,
		name: "Beth",
		}

//A field in a struct is accessed with dotted notation:
		bob.name = "Bob"
		fmt.Println(beth.name)

// You can also declare that a variable implements a struct type without first giving the
// struct type a name. This is called an anonymous struct
var person struct {
		name string
		age int
		pet string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	pet := struct {
			name string
			kind string
	}{
	name: "Fido",
	kind: "dog",
	}
//pet is anonymous struct
//Situations when anonymous struct is useful:
// 1) Translate external data into a struct (unmarshaling data)
// 2) Translate struct into external data like JSON or protocol buffers(marshaling data)


//Comparable struct
//-->Structs that are entirely composed of comparable types are comparable; those with slice or map fields are not


//Go does allow you to perform a type conversion from one struct type to
// another if the fields of both structs have the same names, order, and types.
type firstPerson struct {
			name string
			age int
		}

type secondPerson struct { //Type conversion of an instance of firstPerson to secondPerson allowed but cant use == to compare 
			name string
			age int
		}

type thirdPerson struct { // type conversion of an instance of first to third not allowed due to differnece in order of fields 
			age int
			name string
		}

type fourthPerson struct {
			firstName string // type conversion from first to fourth not allowed due to differnece in field names 
			age int
		}
//Finally, we can’t convert an instance of firstPerson to fifthPerson because there’s an additional field:
		type fifthPerson struct {
		name string
		age int
		favoriteColor string
		}

//if two struct variables are being compared and at least one of them has a type that’s an anonymous struct, you can com‐
// pare them without a type conversion if the fields of both structs have the same names,
// order, and types. You can also assign between named and anonymous struct types if
// the fields of both structs have the same names, order, and types:
type firstPerson struct {
name string
age int
}
f := firstPerson{
name: "Bob",
age: 50,
}
var g struct {
name string
age int
}
// compiles -- can use = and == between identical named and anonymous structs
g = f
fmt.Println(f == g)
