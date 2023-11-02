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
!!!!!!!!!!!  you can ignore all of the return values
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


// Errors Are Values

// First, define your own enumeration
// to represent the status codes:
	type Status int

	const (
	InvalidLogin Status = iota + 1
	NotFound
	)

// Second, define a StatusErr to hold this value and implement the Error interface
	type StatusErr struct {
		Status Status
		Message string
	}
	func (se StatusErr) Error() string {
		return se.Message
	}
// Now we can use StatusErr to provide more details about what went wrong:
	func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
		err := login(uid, pwd)
		if err != nil {
			return nil, StatusErr{
				Status:
				InvalidLogin,
				Message: fmt.Sprintf("invalid credentials for user %s", uid),
			}
		}
		data, err := getData(file)
		if err != nil {
			return nil, StatusErr{
				Status:
				NotFound,
				Message: fmt.Sprintf("file %s not found", file),
			}
		}e
	return data, nil
	}

/*
Even when you define your own custom error types, always use error as the return type for the error result. 
This allows you to return different types of errors from your
function and allows callers of your function to choose not to depend on the specific error type.		
*/

// If you are using your own error type, be sure you don’t return an uninitialized instance.

//Wrapping errors
/*
When you preserve an error while adding
additional information, it is called wrapping the error. 

When you have a series of wrapped errors, it is called an error chain.

The convention is to write : %w at the end of the error for‐
mat string and make the error to be wrapped the last parameter passed to
fmt.Errorf.

The standard library also provides a function for unwrapping errors, the Unwrap
function in the errors package. You pass it an error and it returns the wrapped error,
if there is one. If there isn’t, it returns nil.

*/

	func fileChecker(name string) error {
		f, err := os.Open(name)
		if err != nil {
			return fmt.Errorf("in fileChecker: %w", err) //Wrapping the Error using fmt.Errorf functions
		}
		f.Close()
		return nil
	}

	func main() {
		err := fileChecker("not_here.txt")
		if err != nil {
			fmt.Println(err)
			if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
				fmt.Println(wrappedErr)
			}
		}
	}

/**If you want to create a new error that contains the message from
another error, but don’t want to wrap it, use fmt.Errorf to create
an error, but use the %v verb instead of %w:

}*/
err := internalFunction()
if err != nil {
	return fmt.Errorf("internal failure: %v", err)
}

//Is and As
/*
Wrapping errors is a useful way to get additional information about an error, but it
introduces problems.

If a sentinel error is wrapped, you cannot use == to check for it,
nor can you use a type assertion or type switch to match a wrapped custom error. 

Go solves it with Is and As.

*/
/*
To check if the returned error or any errors that it wraps match a specific sentinel
error instance, use errors.Is. It takes in two parameters, the error that is being
checked and the instance you are comparing against. The errors.Is function returns
true if there is an error in the error chain that matches the provided sentinel error.*/
		func fileChecker(name string) error {
			f, err := os.Open(name)
			if err != nil {
				return fmt.Errorf("in fileChecker: %w", err) //Wrapping Error BY CONVENTION 
			}
			f.Close()
			return nil
		}
		func main() {
			err := fileChecker("not_here.txt")
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					fmt.Println("That file doesn't exist")
				}
			}
		}
// Running this program produces the output:
	// That file doesn't exist

// 	By default, errors.Is uses == to compare each wrapped error with the specified
// error. If this does not work for an error type that you define (for example, if your
// error is a noncomparable type), implement the Is method on your error:
	type MyErr struct {
		Codes []int
	}
	func (me MyErr) Error() string {
		return fmt.Sprintf("codes: %v", me.Codes)
	}
	func (me MyErr) Is(target error) bool {
		if me2, ok := target.(MyErr); ok {
			return reflect.DeepEqual(me, me2) //reflect.DeepEqual can compare anything including slices
		}
		return false
	}
/*The errors.As function allows you to check if a returned error (or any error it
wraps) matches a specific type. It takes in two parameters. The first is the error being
examined and the second is a pointer to a variable of the type that you are looking for.
If the function returns true, an error in the error chain was found that matched, and
that matching error is assigned to the second parameter. If the function returns
false, no match was found in the error chain. Let’s try it out with MyErr:

*/
	err := AFunctionThatReturnsAnError()
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Code)
	}

//Use errors.Is when you are looking for a specific instance or spe‐
// cific values. Use errors.As when you are looking for a specific
// type.

//Wrapping Errors with defer
// Sometimes you find yourself wrapping multiple errors with the same message:
	func DoSomeThings(val1 int, val2 string) (string, error) {
		
		val3, err := doThing1(val1)
		if err != nil {
			return "", fmt.Errorf("in DoSomeThings: %w", err)
		}

		val4, err := doThing2(val2)
		if err != nil {
			return "", fmt.Errorf("in DoSomeThings: %w", err)
		}
		
		result, err := doThing3(val3, val4)
		if err != nil {
			return "", fmt.Errorf("in DoSomeThings: %w", err)
		}
		return result, nil
	}
// We can simplify this code by using defer:
	func DoSomeThings(val1 int, val2 string) (_ string, err error) {
		
		defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
		}()

		val3, err := doThing1(val1)
		if err != nil {
			return "", err
		}

		val4, err := doThing2(val2)
		if err != nil {
			return "", err
		}
		return doThing3(val3, val4)
	}

//panic and recover

/*
Go generates a panic whenever there is a situation where the Go runtime is unable to figure out what should happen next.
 This could be due to a programming error (like an attempt to read past the end of a slice) or 
environmental problem (like running out of memory).

As soon as a panic happens, the current function exits immediately and 
any defers attached to the current function start running.
When those defers complete, the defers attached to the calling function run, and so
on, until main is reached. The program then exits with a message and a stack trace.
*/

/*
The built-in recover function is called from within a defer to
check if a panic happened. If there was a panic, the value assigned to the panic is
returned. Once a recover happens, execution continues normally. Let’s take a look
with another sample program.
*/

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

//The 'recover' funciton doesn't make clear what could fail. It just ensures that if a panic happpens then we can
// print out a message and continue 

/*
While building libraries for third parties, do not let panics escape the boundaries of the public API.
If panic is possible, a public function should use a recover to convert the panic into an
error, return it, and let the calling code decide what to do with them
*/
