//------------------------Types, Methods, and Interfaces----------------------------

//Types and Methods in Go

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

//This should be read as declaring a user-defined type with the name Person to have the underlying type of struct literal that follows
// Not only struct literal you can use any primitive type or compound type literal to define a concrete type

// few examples
type Score int
type Converter func(string) Score
type TeamScores map[string]Score

//An abstract type is one that specifies what a type should
// do, but not how it is done. A concrete type specifies what and how.

//Methods
// The methods for a type are defined at the package block level:
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

//Method names can't be overloaded

//Pointer Receivers and Value Receivers
//Go uses parameters of pointer type to indicate that a
// parameter might be modified by the function.

/*
The same rules apply for method receivers, too.
They can be pointer receivers (the type is a pointer) or value receivers
(the type is a value type).

The following rules help you determine when to use each
kind of receiver:

	1) • If your method modifies the receiver, you must use a pointer receiver.
	2) If your method needs to handle nil instances (see “Code Your Methods for nil
Instances” on page 133), then it must use a pointer receiver.
    3) • If your method doesn’t modify the receiver, you can use a value receiver.

*/

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() { //using pointer reciever as parameter because the method modifies the reciever
	c.total++
	c.lastUpdated = time.Now()
}
func (c Counter) String() string { //using value reciever because the method doesn't modify the reciever
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

//Code Your Methods for nil Instances
// if a method with value reciver gets nil instance it will panic
// a method with pointer reciver works if it's written to handle the possibility of a nil instance

// Here’s an implementation of a binary tree that takes advantage of nil values for the receiver:


package main

import (
	"fmt"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}


//Go allows you to call a method on a nil receiver, and there are
// situations where it is useful, like our tree node example. However, most of the time it’s
// not very useful



//Methods Are Functions Too

	type Adder struct {
		start int
	}
	func (a Adder) AddTo(val int) int {
		return a.start + val
	}

// we create an instance of the type in the usual way and invoke its method:
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15
// We can also assign the method to a variable or pass it to a parameter of type func(int)int. 
	//This is called a method value:
	f1 := myAdder.AddTo
	fmt.Println(f1(10))
	// prints 20


	// A method value is a bit like a closure, since it can access the values in the fields of the
	// instance from which it was created.
	// You can also create a function from the type itself. This is called a method expression:
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // prints 25

	// In the case of a method expression, the first parameter is the receiver for the method;
// our function signature is func(Adder, int) int.



//Functions Versus Methods

/*
package-level state should be effectively immutable.

  	Any time your logic depends on values that are configured at startup or changed while your
program is running, those values should be stored in a struct and that logic should be
implemented as a method. 

	If your logic only depends on the input parameters, then it
should be a function.
*/



//Type Declarations Aren’t Inheritance
// In addition to declaring types based on built-in Go types and struct literals, you can
// also declare a user-defined type based on another user-defined type:
type HighScore Score
type Employee Person


// assigning untyped constants is valid
	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	hs = s // compilation error! You can’t assign an instance of type SCORE to a variable of type HIGHSCORE
	s = i // compilation error!
	s = Score(i)	// ok
	hs = HighScore(s) // ok


//when should you create a user defined type based on builtin types ?
//types are documentation. They make code clearer by providing a name
// for a concept and describing the kind of data that is expected. It’s clearer for someone
// reading your code when a method has a parameter of type Percentage than of type
// int, and it’s harder for it to be invoked with an invalid value


//iota Is for Enumerations—Sometimes
/*
Many programming languages (like APL, i.e. 'A Programming Language') have the concept of enumerations, where you can
specify that a type can only have a limited set of values. Go doesn’t have an enumera‐
tion type. Instead, it has iota, which lets you assign an increasing value to a set of
constants.*/

// When using iota, the best practice is to first define a type based on int that will represent all of the valid values:
		type MailCategory int
// Next, use a const block to define a set of values for your type:
		const (
			Uncategorized MailCategory = iota //type specified and value is set to iota
			Personal //neither the type nor a value assigned
			Spam//neither the type nor a value assigned
			Social //neither the type nor a value assigned
			Advertisements //neither the type nor a value assigned
		)
/*
When the Go compiler sees this, it repeats the type and the assignment to all of the subsequent con‐
stants in the block, and increments the value of iota on each line. 
This means that it
assigns 0 to the first constant (Uncategorized), 
1 to the second constant (Personal),
and so on.
 When a new const block is created, iota is set back to 0.
*/


//Use Embedding for Composition
/*
“Favor object composition over class inheritance”

While Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion:
*/
	type Employee struct {
			Name string
			ID string
		}
	func (e Employee) Description() string {
		return fmt.Sprintf("%s (%s)", e.Name, e.ID)
	}
	type Manager struct {
			Employee //field of type Employee has no name assigned to it. This makes 'Employee' an embedded field.
			Reports []Employee
		}
	func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	}

// 	Any fields or methods declared on an  embedded field are promoted to the containing struct and can be invoked directly on it.
//  That makes the following code valid:
	m := Manager{
			Employee: Employee{
			Name: "Bob Bobson",
			ID: "12345",
		},
			Reports: []Employee{},
	}
	fmt.Println(m.ID) // prints 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)

	// If the containing struct has fields or methods with the same name as an embedded
	// field, you need to use the embedded field’s type to refer to the obscured fields or
	// methods. If you have types defined like this:
		type Inner struct {
				X int
			}
		type Outer struct {
				Inner
				X int
			}
		// You can only access the X on Inner by specifying Inner explicitly:
		o := Outer{	
				Inner: Inner{
					X: 10,
				},
				X: 20,
		}
		fmt.Println(o.X) // prints 20
		fmt.Println(o.Inner.X) // prints 10

//Embedding Is Not Inheritance
// You cannot assign a variable of type
// Manager to a variable of type Employee. If you want to access the Employee field
// in Manager, you must do so explicitly.

	var eFail Employee = m // compilation error!
	var eOK Employee = m.Employee // ok!

// You’ll get the error:	cannot use m (type Manager) as type Employee in assignment
//The methods on the embedded field have no idea they are embedded.

// /If you have a method on an
// embedded field that calls another method on the embedded field, and the containing
// struct has a method of the same name, the method on the embedded field will not
// invoke the method on the containing struct. This behaviour is demonstrated in this code:

	type Inner struct {
		A int
	}
	
	func (i Inner) IntPrinter(val int) string {
		return fmt.Sprintf("Inner: %d", val)
	}
	func (i Inner) Double() string {
		return i.IntPrinter(i.A * 2)
	}
	type Outer struct {
		Inner
		S string
	}
	func (o Outer) IntPrinter(val int) string {
		return fmt.Sprintf("Outer: %d", val)
	}
	func main() {
		o := Outer{
			Inner: Inner{
				A: 10,
			},
			S: "Hello",
		}
		fmt.Println(o.Double())
	}
//Running this code produces the output:
// 	Inner: 20

/*
While embedding one concrete type inside another won’t allow you to treat the outer
type as the inner type, the methods on an embedded field do count toward the
method set of the containing struct. This means they can make the containing struct
implement an interface.
*/



