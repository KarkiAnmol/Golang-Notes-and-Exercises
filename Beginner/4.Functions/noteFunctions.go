//*************************NOTE*****************
//***********************Functions**************

//Declaring and Calling Functions
//the keyword func, the name of the function, the input parameters, and the return type
func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

//Example 5-1. Using a struct to simulate named parameters
type MyFuncOpts struct {
		FirstName string
		LastName string
		Age int
	}
func MyFunc(opts MyFuncOpts) error {
		// do something here
}
func main() {
	MyFunc(MyFuncOpts {
	LastName: "Patel",
	Age: 50,
})
My Func(MyFuncOpts {
	FirstName: "Joe",
	LastName: "Smith",
})
}

//Variadic Input Parameters and Slices
//Go supports Variadic input parameters i.e it allows any number of input parameters .that's why fmt.Println 
// allows any number of input parameter

//The variadic parameter must be
// the last (or only) parameter in the input parameter list. You indicate it with three dots
// (…) before the type.

//The variable that’s created within the function is a slice of the
// specified type. You use it just like any other slice.

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals)) // []
	for _, v := range vals {
		out = append(out, base+v) 
	}
	return out
}

func main() {
	fmt.Println(addTo(3)) // []
	fmt.Println(addTo(3, 2)) // [3]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5,7,9,11]
	a := []int{4, 3} 
	fmt.Println(addTo(3, a...)) //[7,6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4,5,6,7,8]
}
//you can supply however many values you want for the variadic
// parameter, or no values at all. Since the variadic parameter is converted to a slice, you
// can supply a slice as the input. However, you must put three dots (…) after the variable
// or slice literal. If you do not, it is a compile-time error.

//Multiple Return Values
func divAndRemainder(numerator int, denominator int) (int, int, error) { //When a Go function
	// returns multiple values, the types of the return values are listed in parentheses, separated by commas.
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil //putting parentheses around multiple returns 
	// causes compile time error
	//By convention, the error is always the last (or only) value returned from a function.
}	



//Multiple Return Values Are Multiple Values
//You must assign each value returned from a function. If
// you try to assign multiple return values to one variable, you get a compile-time error


//Ignoring Returned Values
//assign the unused values to the name _


//Named Return Values
//specify names for your return values.
//When you supply names to your return values, what you are doing is pre-declaring
// variables that you use within the function to hold the return values
//Named return values are initialized to their zero values when created. This means that we can return them
// before any explicit use or assignment
func divAndRemainder(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

// While named return values can sometimes help clarify your code, they do have some
// potential corner cases. 
// Problems:
// --> shadowing: Just like any other variable,
// you can shadow a named return value. 
// --> with named return values ,you don’t have to return them.

// Go compiler inserts code
// that assigns whatever is returned to the return parameters. The named return param‐
// eters give a way to declare an intent to use variables to hold the return values, but
// don’t require you to use them

//Blank Returns—Never Use These!With named return values you can just write return without specifying  	
// the values that are returned.his returns the last
// values assigned to the named return values
//If your function returns values, never use a blank return. It can
// make it very confusing to figure out what value is actually returned.


//Functions Are Values
// functions in Go are values. The type of a function
// is built out of the keyword func and the types of the parameters and return values.
// This combination is called the signature of the function. Any function that has the
// exact same number and types of parameters and return values meets the type
// signature.

//Having functions as values allows us to do some clever things, such as build a primi‐
// tive calculator using functions as values in a map.

//set of functions that all have the same signature
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

//map to associate operators with functions
var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
func main() {
	expressions := [][]string{
	[]string{"2", "+", "3"},
	[]string{"2", "-", "3"},
	[]string{"2", "*", "3"},
	[]string{"2", "/", "3"},
	[]string{"2", "%", "3"},
	[]string{"two", "+", "three"},
	[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0]) //To convert a string to an int.
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
			if err != nil {
				fmt.Println(err)
				continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
		}
		}

//Output

5
-1
6
0
unsupported operator: %
strconv.Atoi: parsing "two": invalid syntax
invalid expression: [5]



// Function Type Declarations
//Just like you can use the type keyword to define a struct, you can use it to define a
// function type,

	type opFuncType func(int,int) int
// We can then rewrite the opMap declaration to look like this:
	var opMap = map[string]opFuncType {
	// same as before
	}
//We don’t have to modify the functions at all. Any function that has two input parame‐
// ters of type int and a single return value of type int automatically meets the type and
// can be assigned as a value in the map.


//Anonymous Functions
//-->Not only can you assign functions to variables, you can also define new functions
// within a function and assign them to variables.
// These inner functions are anonymous functions;
// Anonymous Functions are called anonymous because they don’t have a name. Neither do you
// have to assign them to a variable. You can just write them inline and call them
// immediately.
func main() {
	for i := 0; i < 5; i++ {
		func(j int) { //Decalring anonymous function with the keyword func immediately followed by
						// the input parameters, the return values, and the opening brace. 
			fmt.Println("printing", j, "from inside of an anonymous function")}(i)
			//It is a compile-time error to try to put a function name between func and the input parameters.
		}
}
// In this example, we are passing the i variable from the for loop in here. It is assigned to the 
// j input parameter of our anonymous function.

//Situations where anonymous functions would be useful
//However, there are two situations where
// 1) defer statements 
// 2) launching goroutines.



//Closures

//Functions declared inside of functions are special; they are closures
//means that functions declared inside of functions are able to
// access and modify variables declared in the outer function

//benefits of closures:
/*
 1) closures allow you to do is limit a function’s scope (If a function is only
going to be called from one other function, but it’s called multiple times, you can use
an inner function to “hide” the called function. This reduces the number of declara‐
tions at the package level, which can make it easier to find an unused name.)

 2)  
*/


//Passing functions as Parameters is very useful
/*
One example is sorting slices. There’s a function in the sort package in the standard
library called sort.Slice. It takes in any slice and a function that is used to sort the
slice that’s passed in. Let’s see how it works by sorting a slice of a struct using two
different fields.
*/

type Person struct {
		FirstName string
		LastName string
		Age int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people) // [{Pat Patterson 37} {Tracy Bobbert 23} {Fred Fredson 18}]

	// Next, we’ll sort our slice by last name and print out the results:

	// sort by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people) // [{Tracy Bobbert 23} {Fred Fredson 18} {Pat Patterson 37}]
/*
The closure that’s passed to sort.Slice has two parameters, i and j, but within the
closure, we can refer to people so we can sort it by the LastName field.
*/

// sort by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // [{Fred Fredson 18} {Tracy Bobbert 23} {Pat Patterson 37}]

//Returning Functions from Functions
// Not only can you use a closure to pass some function state to another function, you
// can also return a closure from a function.

// Here is our function that returns a closure:
func makeMult(base int) func(int) int {
	return func(factor int) int {
	return base * factor
}
}
// And here is how we use it:
func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)

    for i := 0; i < 3; i++ {
			fmt.Println(twoBase(i), threeBase(i))
	}
}

// Running this program gives the following output:
0 0
2 3
4 6


//defer

/*
Programs often create temporary resources, like files or network connections, that
need to be cleaned up. This cleanup has to happen, no matter how many exit points a
function has, or whether a function completed successfully or not. In Go, the cleanup
code is attached to the function with the defer keyword.
*/

//Let’s take a look at how to use defer to release resources.
//We’ll do this by writing a
// simple version of cat, the Unix utility for printing the contents of a file.

func main() {
	if len(os.Args) < 2 { // making sure the file name is specified on the command line by checking the length of os a slice in the os package that contains the name of the program
						  // launched and the arguments passed to it.
		log.Fatal("no file specified")
	}
	
	f, err := os.Open(os.Args[1]) //read-only file handle with the Open function in the os package
	if err != nil { //if there is a problem opening the file
		log.Fatal(err)
	}
	defer f.Close() //Once we know we have a valid file handle, we need to close it after we use it
	                // To ensure the cleanup code runs, we use the defer keyword, followed by a function or method call
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
			if e4rr != nil {
				if err != io.EOF {
				log.Fatal(err)
				}
	break
}
}
}

//Few things to know about Defer keyword

/*
1) you can defer multiple closures in a Go function. They run in LIFO; the last defer
registered runs first.

2) The code within defer closures runs after the return statement. As I mentioned, you
can supply a function with input parameters to a defer. Just as defer doesn’t run
immediately, any variables passed into a deferred closure aren’t evaluated until the
closure runs.*/


//Go Is Call By Value
//You might hear people say that Go is a call by value language and wonder what that
// means. It means that when you supply a variable for a parameter to a function, Go
// always makes a copy of the value of the variable.

//Example