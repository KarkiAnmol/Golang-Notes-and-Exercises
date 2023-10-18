 //---------------------------Blocks---------------------------
 
/** 
Each place where a declaration occurs is called a block. Variables, constants, types,
 and functions declared outside of any functions are placed in the package block.

 Within a function, every set of braces ({}) defines another block

   
*/

//Shadowing Variables
// Example : Shadowing variables
func main() {
x := 10
if x > 5 {
fmt.Println(x) //10
x := 5 //shadowing variable ,this shadowing only exists within this block
fmt.Println(x) //5
}
fmt.Println(x) //10
}

//Example : Shadowing with multiple assignment
func main() {
	x := 10
	if x > 5 {
	x, y := 5, 20 //:= only reuses variables that are declared in the current block
	fmt.Println(x, y) //When using :=, make sure that you don’t have any variables from
	//                 an outer scope on the lefthand side, unless you intend to shadow them.
	}
	fmt.Println(x)
	}
	Running this code gives you:
	5 20
	10
//You also need to be careful to ensure that you don’t shadow a package import. As in you declare
//  and assign fmt a new value which will make fmt package to be unusable for the rest of the main function


//Detecting Shadowed Variables
/*
Install shadow linter 

		If you are building with a Makefile, consider including shadow in the vet task:
		vet:
		go vet ./...
		shadow ./...
		.PHONY:vet


*/

//Universal block
//GO only has 25 keywords because it considers the built-in types (like int and string), 
// constants (like true and false), and functions (like make
 // or close) ,nil to be predeclared identifiers rather than a keyword. predeclared identifiers are placed in universal block
//  which is the block that contains all other blocks

//*****************Not even shadow detects shadowing of universe block identifiers***************


//---------if---------
// you don’t put parenthesis around the condition
// What Go adds is the ability to declare vari‐
// ables that are scoped to the condition and to both the if and else blocks

//Example  Scoping a variable to an if statement
if n := rand.Intn(10); n == 0 {
	fmt.Println("That's too low")
	} else if n > 5 {
	fmt.Println("That's too big:", n)
	} else {
	fmt.Println("That's a good number:", n)
	}


//for, Four Ways
//What makes
// Go different from other languages is that for is the only looping keyword in the lan‐
// guage. Go accomplishes this by using the for keyword in four different formats:

/*
• A complete, C-style for
• A condition-only for
• An infinite for
• for-range
*/

//• A complete, C-style for

	// Example 4-8. A complete for statement
	for i := 0; i < 10; i++ {
	fmt.Println(i)
	}
	//Rules to remember
	 // you must use := to initialize the variables
	 // just like variable declarations in if statements, you can shadow a variable here.

// • A condition-only for
// Example 4-9. A condition-only for statement  (like while statement)
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}


//• An infinite for
// Example 4
package main
import "fmt"
func main() {
	for {
	fmt.Println("Hello")
	}
}

//go version of do-while to run at least once 
	do {
	// things to do in the loop
	} while (CONDITION);
	// The Go version looks like this:
	for {
	// things to do in the loop
		if !CONDITION {
			break
			}
		}

//continue 
// Example  Using continue to make code clearer
for i := 1; i <= 100; i++ {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
		continue
	}
	if i%3 == 0 {
		fmt.Println("Fizz")
		continue
	}
	if i%5 == 0 {
		fmt.Println("Buzz")
		continue
	}
}



// • for-range
// to iterate over elements in some of Go’s built-in types. like strings,
// arrays, slices, and maps.

// Example The for-range loop
evenVals := []int{2, 4, 6, 8, 10, 12}

for i, v := range evenVals { // you get two loop variables, position 'i' and value 'v' 
		fmt.Println(i, v)
	}
//**If you don’t need to access the key, use an underscore (_) as
// the variable’s name. This tells Go to ignore the value.

// Running this code produces the following output:
0 2
1 4
2 6
3 8
4 10
5 12


// Example: only for getting keys from Map
uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
for k := range uniqueNames {
fmt.Println(k)
}

// Iterating over maps
// Example 4-15. Map iteration order varies
m := map[string]int{
"a": 1,
"c": 3,
"b": 2,
}
for i := 0; i < 3; i++ {
fmt.Println("Loop", i)
for k, v := range m {
fmt.Println(k, v)
}
}

// When you build and run this program, the output varies. Here is one possibility:
Loop 0 
c 3
b 2
a 1
Loop 1
a 1
c 3
b 2
Loop 2
b 2
a 1
c 3
/*this order of key-value pairs varying is a security feature. 
  In earlier Go versions, the iteration order for keys in a map was usu‐
ally (but not always) the same,this caused two problems:
 I) People would write code that assumed that the order was fixed, and this would
break at weird times.
 II) If maps always hash items to the exact same values, and you know that a server is
storing some user data in a map, you can actually slow down a server with an
attack called Hash DoS by sending it specially crafted data where all of the keys
hash to the same bucket.

To prevent these,Go implemented two changes:
 I) First, they modified the hash algorithm for maps to include a random
number that’s generated every time a map variable is created.
 II) Next, they made the
order of a for-range iteration over a map vary a bit each time the map is looped over.
These two changes make it far harder to implement a Hash DoS attack.

//EXCEPTION: To make it easier to debug and
log maps, the formatting functions (like fmt.Println) always out‐
put maps with their keys in ascending sorted order.
*/

//Iterating over strings
// Example Iterating over strings
samples := []string{"hello", "apple_π!"}
for _, sample := range samples {
for i, r := range sample {
fmt.Println(i, r, string(r))
}
fmt.Println()
// output
0 104 h
1 101 e
2 108 l
3 108 l
4 111 o

//Looking at the result for “apple_π!”
0 97 a
1 112 p
2 112 p
3 108 l
4 101 e
5 95 _
6 960 π //        the value at position 6 is 960. due to for range loop being iterated over runes not bytes
//First, notice that the first column skips the number 7.
8 33 !

/*
// The for-range value is a copy
	-->	You should be aware that each time the for-range loop iterates over your compound
			type, it copies the value from the compound type to the value variable. Modifying the
			value variable will not modify the value in the compound type.
*/

//Example : Modifying the value doesn’t modify the source

evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals { // value of 'evenVals' is copied into _,v
v *= 2
}
fmt.Println(evenVals) //[2 4 6 8 10 12]

// Labeling your for statements
// Example  Labels
outer:
for _, outerVal := range outerValues {
	for _, innerVal := range outerVal {
		// process innerVal
		if invalidSituation(innerVal) {
			continue outer
		}
	}
// here we have code that runs only when all of the
// innerVal values were sucessfully processed
}


//Choosing the Right for Statement
// A for-range loop is the best way to walk through a string, since it properly
// gives you back runes instead of bytes. Favor a for-range loop when iterating over all the contents of an
// instance of one of the built-in compound types.

/*
When should you use the complete for loop? The best place for it is when you aren’t
iterating from the first element to the last element in a compound type. While you
could use some combination of if, continue, and break within a for-range loop, a
standard for loop is a clearer way to indicate the start and end of your iteration.
*/

//switch
// Example  The switch statement
words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

for _, word := range words {
		switch size := len(word); size {
			case 1, 2, 3, 4:
				fmt.Println(word, "is a short word!")
			case 5:
				wordLen := len(word)
				fmt.Println(word, "is exactly the right length:", wordLen)
			case 6, 7, 8, 9: //an empty case means nothing happens. ,so no output for octopus(7) and gopher(6)
			default:
				fmt.Println(word, "is a long word!")
		}
}
//output
a is a short word!
cow is a short word!
smile is exactly the right length: 5
anthropologist is a long word!