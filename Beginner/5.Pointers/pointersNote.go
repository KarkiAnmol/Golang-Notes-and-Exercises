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




//Pointers Indicate Mutable Parameters
// There is no mechanism in the language to declare that
// other kinds of values are immutable

/*
Rather than declare that some variables and parameters are
immutable, Go developers use pointers to indicate that a parameter is mutable.

Since Go is a call by value language, the values passed to functions are copies.
 For nonpointer types like primitives, structs, and arrays, this means that the called func‐
tion cannot modify the original. Since the called function has a copy of the original
data, the immutability of the original data is guaranteed.

However, if a pointer is passed to a function, the function gets a copy of the pointer.
This still points to the original data, which means that the original data can be modi‐
fied by the called function.
*/

// impliactions
// 1) if you pass  a nil poitner to a function , you cannot make the value non-nil. You can 
// only reasssign a value if there was already a value assigned to the pointer

// 2)if you want the value assigned to a
// pointer parameter to still be there when you exit the function, you must dereference
// the pointer and set the value. If you change the pointer, you have changed the copy,
// not the original. Dereferencing puts the new value in the memory location pointed to
// by both the original and the copy. Here’s a short program that shows how this works:
func failedUpdate(px *int) {
	x2 := 20
	px = &x2 // updates the value in the copy pointer not the original
}
func update(px *int) {
	*px = 20 // Deferencing updates the value in the memory location pointed by the original pointer
}
func main() {
	x := 10
	failedUpdate(&x) //copy the address of x into the parameter px
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}

//Pointers Are a Last Resort



