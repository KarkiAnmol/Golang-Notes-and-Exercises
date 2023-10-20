// & is address operator, it precedes a value type and returns the address of the memory location
// where the value is stored
x := "hello"     // variable holding value in some memory location say 0600H
pointerToX := &x // pointerToX = 0600H

//Dereferencing
// * is indirection operator, it precedes a variable of pointer type and returns the pointed to value

x := 10
pointerToX := &x
fmt.Println(pointerToX)  // prints a memory address
fmt.Println(*pointerToX) // prints 10 -- Deferencing the value pointed by pointerToX
z := 5 + *pointerToX
fmt.Println(z) // prints 15

// Before dereferencing a pointer, you must make sure that the pointer is non-nil. else the program will panic
var x *int
fmt.Println(x == nil) // prints true
fmt.Println(*x)       // panics

// A pointer type is a type that represents a pointer.
//  It is written with a * before a type name.
//  A pointer type can be based on any type:
x := 10
var pointerToX *int
pointerToX = &x

// The built-in function 'new' creates a pointer variable. It returns a pointer to a zero
// value instance of the provided type:
var x = new(int)      // creates a pointer variable 'x' , same as var x *int
fmt.Println(x == nil) // prints false
fmt.Println(*x)       // prints 0

// For structs, use an & before a struct literal to create a pointer instance.
//  You can’t use an & before a primitive literal (numbers, booleans, and
// strings) or a constant because they don’t have memory addresses; they exist only at
// compile time. When you need a pointer to a primitive type, declare a variable and
// point to it:

x := &Foo{} //creating a pointer instance for a struct literal Foo
var y string
z := &y // declaring a variable to point to a primitive type because string literals hold no memory addresses and only exist at compile time

// Not being able to take the address of a constant is sometimes inconvenient. If you
// have a struct with a field of a pointer to a primitive type, you can’t assign a literal
// directly to the field:
type person struct {
	FirstName  string
	MiddleName *string // field of a pointer to a primitive type 'string'
	LastName   string
}
p := person{
	FirstName:  "Pat",
	MiddleName: "Perry", // This line won't compile because you can't assign a literal directly to the field since
	LastName:   "Peterson",
}

// compile and you get the error message:
// cannot take the address of "Perry"

//Two ways to resolve this problem
// 1) introduce a variable to hold the constant value.
// 2) write a helper function that takes in a boolean, numeric, or string type and returns a pointer to that type:

//Helper functions
	func stringp(s string) *string {
		return &s //the constant is copied to a parameter called varibale 's',since it's a variable,it has address in memory
	}
// With that function, you can now write:
	p := person{
		FirstName: "Pat",
		MiddleName: stringp("Perry"), // This works because pointer variable stores the address of the variable that contains the value of the variable
		LastName:"Peterson",
	}
