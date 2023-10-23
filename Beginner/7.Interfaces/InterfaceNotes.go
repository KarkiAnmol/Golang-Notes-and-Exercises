// A Quick Lesson on Interfaces

//Here’s the definition of the Stringer interface in the fmt package:
type Stringer interface {
	String() string //method set of the interface 'Stringer'
}

// In an interface declaration, an interface literal appears after the name of the interface
// type. It lists the methods that must be implemented by a concrete type to meet the
// interface.

//Interfaces are usually named with “er” endings. We’ve already seen fmt.Stringer, but
// there are many more, including io.Reader, io.Closer, io.ReadCloser, json.Mar
// shaler, and http.Handler.

//Interfaces Are Type-Safe Duck Typing
//What makes Go’s interfaces special is that they are implemented implicitly

/*
A concrete type does not declare that it implements an interface.
If the method set for a concrete type contains all of the methods in the method set for an interface,
the concrete type implements the interface.

This means that the concrete type can be assigned to a vari‐
able or field declared to be of the type of the interface.

This implicit behavior makes interfaces the most interesting thing about types in Go,
because they enable both type-safety and decoupling, bridging the functionality in
both static and dynamic languages
*/

/*
	Why interfaces?
--> because “Program to an interface, not an implementation.”

Doing so allows you to depend on behavior, not on implementation, allowing
you to swap implementations as needed. This allows your code to evolve over time, as
requirements inevitably change
*/


/*
Dynamically typed languages like Python, Ruby, and JavaScript don’t have interfaces.
Instead, those developers use “duck typing,” which is based on the expression “If it
walks like a duck and quacks like a duck, it’s a duck.”


*/

	type LogicProvider struct {}

	func (lp LogicProvider) Process(data string) string {
	// business logic
	}

	type Logic interface {
		Process(data string) string
	}

	type Client struct{
		L Logic
	}

	func(c Client) Program() {
		// get data from somewhere
		c.L.Process(data)
	}

	main() {
		c := Client{
			L: LogicProvider{},
		}
		c.Program()
	}
	// In the Go code, there is an interface, but only the caller (Client) knows about it;
// there is nothing declared on LogicProvider to indicate that it meets the interface.

//If there’s an interface in the standard library that describes what
// your code needs, use it!

//It’s perfectly fine for a type that meets an interface to specify additional methods that
// aren’t part of the interface.



//Embedding and Interfaces

// Just like you can embed a type in a struct, you can also embed an interface in an inter‐
// face. For example, the io.ReadCloser interface is built out of an io.Reader and an
// io.Closer:
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	type ReadCloser interface {
		Reader
		Closer
	}
// Just like you can embed a concrete type in a struct, you can also
// embed an interface in a struct.


//Accept Interfaces, Return Structs
/*
What the pharse means is that the business logic invoked by your functions should be invoked via interfaces,
but the output of your functions should be a concrete type.

functions accepting interfaces makes the code flexible and explicitly declare exactly what functionality is being used

REASONS TO AVOID RETURNING INTERFACES

1) 	you lose the main advvantage of implicit interfaces : decoupling 

2)  Another reason to avoid returning interfaces is versioning. If a concrete type is
	returned, new methods and fields can be added without breaking existing code. 
	The same is not true for an interface.Adding a new method to an interface means that
	you need to update all existing implementations of the interface, or your code breaks.
*/



// Interfaces and nil
/*
We also use nil to represent the zero value for an interface instance
but it’s not as simple as it is for concrete types.

In order for an interface to be considered nil both the type and the value must be
nil
*/

	var s *string
	fmt.Println(s == nil) // prints true
	
	var i interface{}
	fmt.Println(i == nil) // prints true
	
	i = s
	fmt.Println(i == nil) // prints false

// In the Go runtime, interfaces are implemented as a pair of pointers, one to the underlying type and one to the underlying value. 
// As long as the type is non-nil, the interface is non-nil. (Since you cannot have a variable without a type, if the value pointer
// is non-nil, the type pointer is always non-nil.)

//What nil indicates for an interface is whether or not you can invoke methods on it.
// If an interface is nil, invoking any methods on it triggers a panic
//If an interface is non-nil,
// you can invoke methods on it. (But note that if the value is nil anhd the methods of
	// the assigned type don’t properly handle nil, you could still trigger a panic.)



// The Empty Interface Says Nothing
// Sometimes in a statically typed language, you need a way to say that a variable could
// store a value of any type. Go uses interface{} to represent this:

	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName string
	} {"Fred", "Fredson"}
	//Because an empty inter‐
// face doesn’t tell you anything about the value it represents,this just matches every type in GO

// One common use of the empty interface is as a placeholder for data of uncer‐
// tain schema that’s read from an external source, like a JSON file:

	// one set of braces for the interface{} type,
	// the other to instantiate an instance of the map
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("testdata/sample.json")

	if err != nil {
		return err
	}
	defer contents.Close()
	json.Unmarshal(contents, &data) // the contents are now in the data map


//If you see a function that takes in an empty interface, it’s likely that it is using reflec‐
// tion (which we’ll talk about in Chapter 14) to either populate or read the value



//If you find yourself in a situation where you had to store a value into an empty inter‐
// face, you might be wondering how to read the value back again. To do that, we need
// to look at type assertions and type switches.


