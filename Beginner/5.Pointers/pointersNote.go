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
//Pointers make it harder to understand data flow and can create extra work for garbage collector 

//Rather than populating a struct by passing a pointer to it into a function, have
// the function instantiate and return the struct

// Example 6-3. Don’t do this
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Example 6-4. Do this
func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

//The only time you should use pointer parameters to modify a variable is when the
// function expects an interface. You see this pattern when working with JSON

//When returning values from a function, you should favor value types. Only use a
// pointer type as a return type if there is state within the data type that needs to be  modified.


//Pointer Passing Performance
/*
If a struct is large enough, there are performance improvements from using a pointer
to the struct as either an input parameter or a return value.
the size of a pointer is the same for all data types. Passing a value into a func‐
tion takes longer as the data gets larger. so it makes sense to use pointer to make the program faster*/

//For data structures that are smaller than a megabyte, it is actually slower to return a
// pointer type than a value type

//For the vast majority of cases, the
// difference between using a pointer and a value won’t affect your program’s perfor‐
// mance. But if you are passing megabytes of data between functions, consider using a
// pointer even if the data is meant to be immutable.



//The Zero Value Versus No Value
/*
Nil pointers can be used to distinguish zero value in your program when the distinction matters,but we don't
prefer to use it because using pointer means making the value mutable, so we instead prefer the comma ok idiom
*/

//within the Go runtime, a map is implemented as a pointer to a struct. 
// Passing a map to a function means that you are copy‐
// ing a pointer.this is why  any modifications made to a map that's passed to a function are reflected in the original value
// that was passed in
//Because of this, you should avoid using maps for input parameters or return values,
// especially on public APIs

// Go is a strongly typed language; rather than passing a map around, use a struct


// Passing slice to a function has more complicated behaviour because any modifications made to the contents of the slice is reflected
// in the original value but using append to change the length isn't reflected in the original value, even if the slice has capacity 
// greater than its length.
// That’s because a slice is implemented as a struct with three fields: 
		// 1. an int field for length
		// 2. an int field for capacity, and 
		// 3. a pointer to a block of memory.
// Changing the values in the slice changes the memory that the pointer points to, so the
// changes are seen in both the copy and the original

// Changes to the length and capacity are not reflected back in the original, because they are only in the copy.
//  Changing the capacity means that the pointer is now pointing to a new, bigger block of memory.


//By default, you should assume that a slice is not modified by a function.
// Your function’s documentation should specify if it modifies the slice’s contents.


/*
********IMPORTANT **

The reason you can pass a slice of any size to a function is that the
data that’s passed to the function is the same for any size slice: two
int values and a pointer. The reason that you can’t write a function
that takes an array of any size is because the entire array is passed
to the function, not just a pointer to the data.

*/

//Slices as Buffers

/*
Even though Go is a garbage-collected language, writing idiomatic Go means avoid‐
ing unneeded allocations. Rather than returning a new allocation each time we read
from a data source, we create a slice of bytes once and use it as a buffer to read data
from the data source:
*/

file, err := os.Open(fileName)
if err != nil {
		return err
}
 defer file.Close()
 data := make([]byte, 100)
 for {
		count, err := file.Read(data)
		if err != nil {
		return err
		}		
if count == 0 {
	return nil
}
process(data[:count])
}
// Remember that we can’t change the length or capacity of a slice when we pass it to a
// function, but we can change the contents up to the current length. In this code, we
// create a buffer of 100 bytes and each time through the loop, we copy the next block of
// bytes (up to 100) into the slice. We then pass the populated portion of the buffer to
// process.



//Reducing the Garbage Collector’s Workload
/*
Using buffers is just one example of how we reduce the work done by the garbage
collector. When programmers talk about “garbage” what they mean is “data that has
no more pointers pointing to it.”

The job of a garbage collector is to automatically detect unused memory and
recover it so it can be reused.

Go is unusual in that it can actually increase the size of a stack
while the program is running. This is possible because each gorout‐
ine has its own stack and goroutines are managed by the Go run‐
time, not by the underlying operating system

*/






