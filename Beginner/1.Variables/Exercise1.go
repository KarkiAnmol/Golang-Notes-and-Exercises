/* Write a Go program that declares and initializes variables of various data types
 */

package main

import "fmt"

func main() {

	//Declaring variables

	var num int
	var str string

	//Declaring multiple variables
	var (
		flag bool // no value assigned, set to false
		fnum float32
	)

	//Initializing Varibales
	num = 12_34 //These underscores have no effect on the value of the number. They improve the readability of literals
	str = "\n ..... Greetings ..... \n and \n  \"Salutations\"  \n"
	fnum = 99.31

	//Declaring and initializing variables with short declaration operator `:=`
	//:= is not legal outside of functions
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)

	const cnum int64 = 101

	rstr := rune('a')

	//Printing
	if !flag {
		fmt.Println(str)
	}
	fmt.Println(num, fnum, cnum)
	fmt.Println(`real value of x =`, real(x))
	fmt.Println(`imaginary value of y =`, imag(y))
	fmt.Println(x)
	fmt.Println(x * y)
	fmt.Println(rstr)

}
