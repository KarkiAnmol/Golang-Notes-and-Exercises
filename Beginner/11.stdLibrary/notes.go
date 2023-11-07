//The Standard Library
//io and Friends
//Both io.Reader and io.Writer define a single method:

type Reader interface {
	//takes in a slice of bytes and this slice is passed into implementation of method and modified
	//the slice will be written upto len(p) bytes and the method returns the number of bytes written
	Read(p []byte) (n int, err error)
}
type Writer interface {
	//takes in a slice of bytes and returns number of bytes written 'n'
	Write(p []byte) (n int, err error)
}

// You might expect this:
type NotHowReaderIsDefined interface {
	Read() (p []byte, err error)
}

// but There’s a very good reason why io.Reader is defined the way it is
// Let’s write a func‐
// tion that’s representative of how to work with an io.Reader to understand:
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048) //slice of bytes, len(buf) is 2048 ,a buffer of size 2048 bytes
	out := map[string]int{}   // map string: int
	for {
		n, err := r.Read(buf)       //invoking read method of io.Reader package that returns the number of bytes written to n
		for _, b := range buf[:n] { // ranging loop over slice 'buf'
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++ //incrementing integer value in map for every alphabet
			}
		}
		if err == io.EOF { // it isn’t really an error. It indicates that
			// there’s nothing left to read from the io.Reader. When io.EOF is returned, we are fin‐
			// ished processing and return our result.
			return out, nil
		}
		if err != nil { // if error exists,return nil map and error value
			return nil, err
		}
	}
}

//If the Read method were written to return a []byte, it would
// require a new allocation on every single call. Each allocation would end up on the
// heap, which would make quite a lot of work for the garbage collector.
// Since, we create our buffer once and reuse it on every
// call to r.Read. This allows us to use a single memory allocation to read from a poten‐
// tially large data source.
//
//If we want to reduce the allocations further, we could create a pool of buffers when
// the program launches. We would then take a buffer out of the pool when the function
//starts, and return it when it ends. By passing in a slice to io.Reader, memory alloca‐
// tion is under the control of the developer.

//In most cases when
// a function or method has an error return value, we check the error before we try to
// process the nonerror return values. We do the opposite for Read because there might
// have been bytes returned before an error was triggered by the end of the data stream
// or by an unexpected condition.



//TIME
// There are two main types used to represent time,
// time.Duration and time.Time.
// A period of time is represented with a time.Duration, a type based on an int64. 
//  For example, you represent a duration of 2 hours and 30 minutes with:
	d := 2 * time.Hour + 30 * time.Minute // d is of type time.Duration
// These constants make the use of a time.Duration both readable and type-safe. 

// Monotonic Time
// Most operating systems keep track of two different sorts of time: 
// the wall clock : corresponds to the current time
// monotonic clock: which simply counts up from the time the computer was booted

// The reason for tracking two different clocks
// is that the wall clock doesn’t uniformly increase. Daylight Saving Time, leap seconds,
// and NTP (Network Time Protocol) updates can make the wall clock move unexpect‐
// edly forward or backward. This can cause problems when setting a timer or finding
// the amount of time that’s elapsed.
// To address this potential problem, Go uses monotonic time to track elapsed time
// whenever a timer is set or a time.Time instance is created with time.Now.




// ---------------encoding/json--------------------
// marshaling: converting from a Go data type to an encoding
// unmarshaling: converting to a Go data type
// 
// Use Struct Tags to Add Metadata
// Let’s say that we are building an order management system and have to read and
// write the following JSON:
	{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
	}
// We define types to map this data:
	type Order struct {
		ID string `json:"id"`
		DateOrdered time.Time `json:"date_ordered"` 
		CustomerID string `json:"customer_id"`
		Items []Item `json:"items"`
	}
	type Item struct {
		ID string `json:"id"`
		Name string `json:"name"`
	}

// For JSON processing, we use the tag name json to specify the name of the JSON field
// that should be associated with the struct field. If no json tag is provided, the default
// behavior is to assume that the name of the JSON object field matches the name of the
// Go struct field.

//Struct tags are never evaluated auto‐
// matically; they are only processed when a struct instance is passed into a function.



//Unmarshaling and Marshaling
// 
// The Unmarshal function in the encoding/json package is used to convert a slice of
// bytes into a struct. If we have a string named data, this is the code to convert data to
// a struct of type Order:
	var o Order // instance of struct Order assigned to variable 'o'
	data := "apple"
	// The json.Unmarshal function populates data into an input parameter, just like the
	// implementations of the io.Reader interface.
	err := json.Unmarshal([]byte(data), &o) // convert data to struct
	if err != nil {
		return err
	}

// We use the Marshal function in the encoding/json package to write an Order
// instance back as JSON, stored in a slice of bytes:
	out, err := json.Marshal(o) // convert struct instance to data in bytes


//JSON, Readers, and Writers
// 
// The json.Marshal and json.Unmarshal functions work on slices of bytes.As we just
// saw, most data sources and sinks in Go implement the io.Reader and io.Writer
// interfaces. While you could use ioutil.ReadAll to copy the entire contents of an
// io.Reader into a byte slice so it can be read by json.Unmarshal, this is inefficient.
// Similarly, we could write to an in-memory byte slice buffer using json.Marshal and
// then write that byte slice to the network or disk, but it’d be better if we could write to
// an io.Writer directly.
// 
// The encoding/json package in Go provides the json.Decoder and json.Encoder types,
// allowing interaction with anything that satisfies the io.Reader and io.Writer interfaces, respectively.

// Let's start by defining a simple struct, 'Person', representing our data:
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// We create an instance of the 'Person' struct:
toFile := Person{
	Name: "Fred",
	Age:  40,
}

// As the os.File type implements both io.Reader and io.Writer interfaces, we use it to 
// illustrate json.Decoder and json.Encoder capabilities.
// First, we write 'toFile'(instance of the struct) to a temporary file using json.NewEncoder, which returns a 
// json.Encoder for the temporary file. The 'Encode' method then writes 'toFile' to the file:

// Creates a new temporary file using ioutil.TempFile function.
// Parameters:
// - os.TempDir(): Retrieves the system's temporary directory path.
// - "sample-": Prefix for the generated temporary file's name.
// The function returns a file descriptor 'tmpFile' and an error 'err'.
// 'tmpFile' is a handle to the newly created temporary file, while 'err' contains
// an error, if any, encountered during the file creation process.
tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
if err != nil {
    panic(err)
}
defer os.Remove(tmpFile.Name()) // Ensure file cleanup
err = json.NewEncoder(tmpFile).Encode(toFile)
if err != nil {
    panic(err)
}
err = tmpFile.Close() // Close the file after writing
if err != nil {
    panic(err)
}

// Once 'toFile' is written, we can read the JSON data back in.
//  By opening the temporary file and passing a reference to json.NewDecoder, 
// we decode the JSON data using the 'Decode' method with a variable of type Person:


// Opens the temporary file for reading.
// 'tmpFile.Name()' retrieves the complete name of the previously created temporary file.
// 'tmpFile2' is a file descriptor or a reference for the opened file, and 'err' captures any errors
//  encountered during the opening process.
tmpFile2, err := os.Open(tmpFile.Name())

if err != nil {
    panic(err)
}
var fromFile Person
// Uses a JSON decoder to read and decode the contents of the opened file 'tmpFile2'.
// The 'Decode' method reads JSON data from 'tmpFile2' and populates the 'fromFile' variable
// of type 'Person' (defined earlier) with the decoded data. 'err' captures any errors
// encountered during the decoding process.
err = json.NewDecoder(tmpFile2).Decode(&fromFile)
if err != nil {
    panic(err)
}
err = tmpFile2.Close() // Close the file after reading
if err != nil {
    panic(err)
}

// Finally, we print the retrieved data in a structured way:
fmt.Printf("%+v\n", fromFile)




//Encoding and Decoding JSON Streams
// when you have multiple JSON structs to read or write at once json.Decoder and json.Encoder can be used
func main() {
	// JSON data with multiple objects
	data := `{"name": "Fred", "age": 40}
			{"name": "Mary", "age": 21}
			{"name": "Pat", "age": 30}`

	// Initialize a struct to store the JSON data
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Reading multiple JSON objects using json.Decoder
	// 
	//initializes a new json.Decoder (i.e.'dec') by associating it with a string
	// strings.NewReader creates a new io.Reader from the provided string data. 
	// This function returns a *strings.Reader, which implements the io.Reader interface, 
	// allowing reading from the provided string.
	dec := json.NewDecoder(strings.NewReader(data))
	var t Person //instance of the struct
	// 
	// The json.Decoder can then be used to decode and process JSON data from this string source 
	// using its methods, such as More() to check for additional content and Decode() to decode JSON objects 
	// one by one.
    // dec.More() is a method of the json.Decoder that checks if there is more content to be decoded 
	// from the input. It returns true if there is more content to read.
	for dec.More() {
		err := dec.Decode(&t) //decoding the next JSON object from the input source into the variable t, which should be a struct
		if err != nil {
			panic(err)
		}
		// Process t - Here you might perform actions with each decoded object
		fmt.Printf("Name: %s, Age: %d\n", t.Name, t.Age)
	}

	// Writing multiple values using json.Encoder
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	allInputs := []Person{
		{Name: "Fred", Age: 40},
		{Name: "Mary", Age: 21},
		{Name: "Pat", Age: 30},
	}

	for _, input := range allInputs {
		err := enc.Encode(input)
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
	fmt.Println("Encoded JSON:", out)

	// Additional note: Reading a single object from an array using json.Decoder
	// This can be done to process JSON objects within an array without loading the entire array into memory at once.
	// Refer to the Go documentation for an example demonstrating this approach.
}
