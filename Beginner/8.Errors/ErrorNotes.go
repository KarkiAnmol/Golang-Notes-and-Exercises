//How to Handle Errors: The Basics
//Go handles errors by returning a value of type
// error as the last return value for a function.

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

/*A new error is created from a string by calling the New function in the errors pack‐
age. Error messages should not be capitalized nor should they end with punctuation
or a newline. In most cases, you should set the other return values to their zero values
when a non-nil error is returned.*/

/*
Go doesn’t have special constructs to detect if an
error was returned. Whenever a function returns, use an if statement to check the
error variable to see if it is non-nil:
*/

// error is a built-in interface that defines a single method:
type error interface {
	Error() string
}

/*
Anything that implements this interface is considered an error. The reason why we
return nil from a function to indicate that no error occurred is that nil is the zero
value for any interface type.
*/

//reasons why Go uses a returned error instead of thrown exceptions.
/*
1)First, exceptions add at least one new code path through the code. These
paths are sometimes unclear, especially in languages whose functions don’t include a
declaration that an exception is possible. This produces code that crashes in surpris‐
ing ways when exceptions aren’t properly handled, or, even worse, code that doesn’t
crash but whose data is not properly initialized, modified, or stored.

2) The Go compiler requires that all variables must be read. Making errors return val‐
ues forces developers to either check and handle error conditions or make it explicit
that they are ignoring errors by using an underscore (_) for the returned error value
*/

/*
!!!!!!!!!!!you can ignore all of the return values
from a function. If you ignore all the return values, you would be
able to ignore the error, too. In most cases, it is very bad form to
ignore the values returned from a function. Please avoid this,
except for cases like fmt.Println.
*/

//The error handling is indented inside
// an if statement. The business logic is not

//----Use Strings for Simple Errors----
// two ways to create an error from a string
// 1) first is errors.New function. It takes in a string and returns an error
func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

// 2) The second way is to use the fmt.Errorf function.
// This function allows you to use all
// of the formatting verbs for fmt.Printf to create an error. Like errors.New, this string
// is returned when you call the Error method on the returned error instance:
func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i
}

//Sentinel Errors
/*
Sentinel errors are usually used to indicate that you cannot start or continue process‐
ing.

Sentinel errors are one of the few variables that are declared at the package level. By
convention, their names start with Err

They should be treated as read-only because it is a programming error to change their value
*/

func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat { //ErrFormat is a sentinel error. which is returned when the data that was passed in isn't a ZIP file
		fmt.Println("Told you so")
	}
}
//Another sentinel error in the package is rsa.ErrMessageTooLong in the crypto/rsa package

//Using Constants for Sentinel Errors
package consterr
	type Sentinel string
	func(s Sentinel) Error() string {
	return string(s)
}
package mypkg
const (
	ErrFoo = consterr.Sentinel("foo error") //casting a error literal to a string that implements a type   	 		
	ErrBar = consterr.Sentinel("bar error")
)
