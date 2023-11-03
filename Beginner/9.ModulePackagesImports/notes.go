//Repositories, Modules, and Packages
/*
library manageent in go is based around these three concepts
- Repository is a place in a version control system where the source code for a project is stored
- Module is the root of a go library stored in a repository
- Module consists of one or more packages, which give the module organization and structure.
*/

//go.mod
// The command 
go mod init MODULE_PATH
// creates the go.mod file that makes the current directory the root of a module.


//Building Packages
//Imports and Exports
/*
Go’s import statement allows you to access
exported constants, variables, functions, and types in another package

A package’s exported identifiers (an identifier is the name of a variable, coonstant, type, function,
method, or a field in a struct) cannot be accessed from another current package
without an import statement.
*/

// how do you export an identifier in Go?
// Only An identifier whose name starts
// with an uppercase letter is exported



//Creating and Accessing a Package
	package math // package clause (The package clause is always the first nonblank, non-comment line in a Go source file)
	func Double(a int) int {
		return a * 2
	}


// You must specify an import path
// when importing from anywhere besides the standard library.
// prefer absolute path over relative path


// In formatter, there’s a file called formatter.go with the following contents:
package print
import "fmt"
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}

// Finally, the following contents are in the file main.go in the root directory:
package main

import (
	"fmt" // fmt package is from standard library
	"github.com/learning-go-book/package_example/formatter" //packages within our program,
	"github.com/learning-go-book/package_example/math"//so when importing packages anywhere else from standard library 
													  // we need to specify the path(ahsolute path is preferred for clarity)
)
func main() {
	num := math.Double(2) // accessing the double fuction from math package by prefexing the function name with the package name
	output := print.Format(num)
	fmt.Println(output)
}

//Every Go file in a directory must have an identical package clause. i.e. usually package main
// name of the package is determined by its package clause not the absolute path,We
// imported the print package with the import path github.com/learning-go-book/pack‐
// age_example/formatter. 


//make the name of the package match the name of the directory that contains the package

/*
However, there are a few situations where you use a different name for the package than for the directory.

1) The first is something we have been doing all along without realizing it. We declare a
package to be a staring point for a Go application by using the special package name
main. Since you cannot import the main package, this doesn’t produce confusing
import statements.
2) The other reasons for having a package name not match your directory name are less
common. If your directory name contains a character that’s not valid in a Go identi‐
fier, then you must choose a package name that’s different from your directory name.
It’s better to avoid this by never creating a directory with a name that’s not a valid
identifier.
3) The final reason for creating a directory whose name doesn’t match the package name
is to support versioning using directories.
*/



// Overriding a Package’s Name
	import (
		crand "crypto/rand"
		"encoding/binary"
		"fmt"
		"math/rand"
	)
	// We import crypto/rand with the name crand. This overrides the name rand that’s
	// declared within the package. We then import math/rand normally. When you look at
	// the seedRand function, you see that we access identifiers in math/rand with the rand
	// prefix, and use the crand prefix with the crypto/rand package:
	func seedRand() *rand.Rand {
		var b [8]byte
		_, err := crand.Read(b[:])
		if err != nil {
			panic("cannot seed with cryptographic random number generator")
		}
		r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
		return r
	}

//Package Comments and godoc

Here are the rules:

• Place the comment directly before the item being documented with no blank
lines between the comment and the declaration of the item.
• Start the comment with two forward slashes (//) followed by the name of the item.
• Use a blank comment to break your comment into multiple paragraphs.
• Insert preformatted comments by indenting the lines. 

Example 9-1. A package-level comment

// Package money provides various utilities to make it easy to manage money.
package money



Example 9-2. A struct comment

// Money represents the combination of an amount of money
// and the currency the money is in.
type Money struct {
	Value decimal.Decimal
	Currency string
}


Example 9-3. A well-commented function

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
//	USD - US Dollar
//	CAD - Canadian Dollar
//	EUR - Euro
//	INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
// ...
}
//Go linting tools such
// as golint and golangci-lint can report missing comments on
// exported identifiers.

//The command go
// doc PACKAGE_NAME displays the package godocs for the specified package and a list of
// the identifiers in the package



//dont use init functions
/*
When you declare a function named init that takes no
parameters and returns no values, it runs the first time the package is referenced by
another package. Since init functions do not have any inputs or outputs, they can
only work by side effect, interacting with package-level functions and variables.
The init function has another unique feature. Go allows you to declare multiple
init functions in a single package, or even in a single file in a package. There’s a
*/
// a blank import
// triggers the init function in a package but doesn’t give you access to any of the
// exported identifiers in the package:
import (
	"database/sql"
	_ "github.com/lib/pq" //blank import using '_'
)

/*any package-level variables configured via init
should be effectively immutable. While Go doesn’t provide a way to enforce that their
value does not change, you should make sure that your code does not change them. If
you have package-level variables that need to be modified while your program is run‐
ning, see if you can refactor your code to put that state into a struct that’s initialized
and returned by a function in the package.
*/

//Only declare a single init function per package and document init function that loads files or accesses network 
// so that security conscious users are not surprised by unexpected I/O


//Circular Dependencies
/*
-No circular dependency
Go does not allow you to have a circular dependency between packages.
This means that if package A imports package B, directly or indirectly, package B can‐
not import package A, directly or indirectly.

-
If you find yourself with a circular dependency, you have a few options. In some
cases, this is caused by splitting packages up too finely. If two packages depend on
each other, there’s a good chance they should be merged into a single package
*/


//Gracefully renaming and Reorganizing your API

/*You might want to rename some of the exported identifiers or move them to another package
within your module. To avoid a backward-breaking change, don’t remove the original
identifiers; provide an alternate name instead.

With a function or method, this is easy. You declare a function or method that calls
the original. For a constant, simply declare a new constant with the same type and
value, but a different name.

When you want to rename or move an exported type, you have to use an alias. Quite
simply, an alias is a new name for a type
*/
// Let’s say we have a type called Foo:
	type Foo struct {
		x int
		S string
	}
	func (f Foo) Hello() string {
		return "hello"
	}
	func (f Foo) goodbye() string {
		return "goodbye"
	}
	// If we want to allow users to access Foo by the name Bar, all we need to do is:
	type Bar = Foo
	// To create an alias, we use the type keyword, the name of the alias, an equals sign, and
	// the name of the original type.



//Working with Modules
// Importing third party code
/*
Unlike many other compiled languages, Go compiles all code for your application
into a single binary, whether it was code you wrote or code from third parties. Just as
we saw when we imported a package from within our own project, when you import
a third-party package, you specify the location in the source code repository where
the package is located
*/


// //Working with Versions
// we can see what versions of the module
// are available with the go list command:
$ go list -m -versions github.com/learning-go-book/simpletax
github.com/learning-go-book/simpletax v1.0.0 v1.1.0
//By default, the go list command lists the packages that are used in your project. The
// -m flag changes the output to list the modules instead, and the -versions flag changes
//go list to report on the available versions for the specified module

// In this case, we
// see that there are two versions, v1.0.0 and v1.1.0. Let’s downgrade to version v1.0.0
// and see if that fixes our problem. We do that with the go get command:
$ go get github.com/learning-go-book/simpletax@v1.0.0
// The go get command lets us work with modules, updating the versions of our
// dependencies.

// Now if we look at go.mod, we’ll see the version has been changed:
// module region_tax
go 1.15
require (
	github.com/learning-go-book/simpletax v1.0.0
	github.com/shopspring/decimal v1.2.0
)
// We also see in go.sum that it contains both versions of simpletax:
	github.com/learning-go-book/simpletax v1.0.0 h1:iH+7ADkdyrSqrMR2GzuWS...
	github.com/learning-go-book/simpletax v1.0.0/go.mod h1:/YqHwHy95m0M4Q...
	github.com/learning-go-book/simpletax v1.1.0 h1:Z/6s1ydS/vjblI6PFuDEn...
	github.com/learning-go-book/simpletax v1.1.0/go.mod h1:/YqHwHy95m0M4Q...
// This is fine; if you change a module’s version, or even remove a module from your
// project, there still might be an entry for it in go.sum. This doesn’t cause any problems.

//For all major versions besides 0 and 1, the path to the module must end in vN,
// where N is the major version.

// First, we are going to change our import of
// simpletax to:
"github.com/learning-go-book/simpletax/v2"

// This changes our import to refer to the v2 module.
// Next, we’re going to change the code in main to the following:
func main() {
	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	zip := os.Args[2]
	country := os.Args[3]
	percent, err := simpletax.ForCountryPostalCode(country, zip)
	if err != nil {
		log.Fatal(err)
	}
	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(total)
}

//Vendoring
/*
To ensure that a module always builds with identical dependencies, some organiza‐
tions like to keep copies of their dependencies inside their module. This is known as
vendoring.
It’s enabled by running the command go mod vendor. This creates a direc‐
tory called vendor at the top level of your module that contains all of your module’s
dependencies.
If new dependencies are added to go.mod or versions of existing dependencies are
upgraded with go get, you need to run go mod vendor again to update the vendor
directory. If you forget to do this, go build, go run, and go test will refuse to run
and display an error message.
The advantage of vendoring is that you know exactly
what third-party code is going to be used by your project. The downside is that it dra‐
matically increases the size of your project in version control
*/
