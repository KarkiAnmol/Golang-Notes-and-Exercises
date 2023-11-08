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










//----------------------------net/http-----------------------
// 
// 
// Go’s standard library includes something that other language distributions had
// considered the responsibility of a third party: a production quality HTTP/2 client and
// server.
// 
// The Client
// The net/http package defines a Client type to make HTTP requests and receive HTTP responses.
// A default client instance (cleverly named DefaultClient) in the net/http package, but you should avoid using it in production applications,
// because it defaults to having no timeout. Instead, instantiate your own.
// 
// You only need to create a single http.Client for your entire program, as it properly handles multiple
// simultaneous requests across goroutines:
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
// When you want to make a request, you create a new *http.Request instance with the
// http.NewRequestWithContext function, passing it a context, the method, and URL
// that you are connecting to. 
// If you are making a PUT, POST, or PATCH request, specify
// the body of the request with the last parameter as an io.Reader. If there is no body,use nil:
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}

// Once you have an *http.Request instance, you can set any headers via the Headers field of the instance. 
// Call the Do method on the http.Client with your http.Request
// and the result is returned in an http.Response:
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

// The response has several fields with information on the request. 
// -> The numeric code of the response status is in the StatusCode field, 
// -> the text of the response code is in the Status field, 
// -> the response headers are in the Header field, and 
// -> any returned content is in a Body field of type io.ReadCloser. 
// This allows us to use it with json.Decoder to process REST API responses:

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))

	var data struct {
		UserID int `json:"userId"`
		ID int `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)	

// There are functions in the net/http package to make GET, HEAD,
// and POST calls. Avoid using these functions because they use the
// default client, which means they don’t set a request timeout



//---------------The Server--------------
// 
// 
// The HTTP Server is built around the concept of an http.Server and the
// http.Handler interface. Just as the http.Client sends HTTP requests, the
// http.Server is responsible for listening for HTTP requests. It is a performant
// HTTP/2 server that supports TLS.

// A request to a server is handled by an implementation of the http.Handler interface
// that’s assigned to the Handler field. This interface defines a single method:
	type Handler interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}
// The *http.Request should look familiar, as it’s the exact same type that’s used to send
// a request to an HTTP server. The http.ResponseWriter is an interface with three
// methods:
	type ResponseWriter interface {
		Header() http.Header
		Write([]byte) (int, error)
		WriteHeader(statusCode int)
	}
// These methods must be called in a specific order. 
// -> (Header)First, call Header to get an instance of http.Header and set any response headers you need. 
//    If you don’t need to set any headers, you don’t need to call it. 
// -> (WriteHeader) Next, call WriteHeader with the HTTP status code for your response. (All the status codes are defined 
// 	  as constants in the net/http package. This would have been a good place to define a custom type, 
//	  but that was not done; all status code constants are untyped integers.)
//	  If you are sending a response that has a 200 status code, you can skip WriteHeader. 
// -> (Write) Finally, call the Write method to set the body for the response. 

// Here’s what a trivial handler looks like:

	type HelloHandler struct{}
	func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	}
// You instantiate a new http.Server just like any other struct:
	s := http.Server{
		Addr:		 ":8080", 
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
// The Addr field specifies the host and port the server listens on. If you don’t specify
// them, your server defaults to listening on all hosts on the standard HTTP port, 80

// You specify timeouts for the server’s reads, writes, and idles using time.Duration values
// Be sure to set these to properly handle malicious or broken HTTP clients, as the
// default behavior is to not time out at all. 
// Finally, you specify the http.Handler for your server with the Handler field.


// A server that only handles a single request isn’t terribly useful, so the Go standard
// library includes a request router, *http.ServeMux. You create an instance with the
// http.NewServeMux function. It meets the http.Handler interface, so it can be
// assigned to the Handler field in http.Server. It also includes two methods that allow
// it to dispatch requests. The first method is simply called Handle and takes in two
// parameters, a path and an http.Handler. If the path matches, the http.Handler is
// invoked.
// While you could create implementations of http.Handler, the more common pattern
// is to use the HandleFunc method on *http.ServeMux:
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	})

// This method takes in a function or closure and converts it to a http.HandlerFunc.

// Because an *http.ServeMux dispatches requests to http.Handler instances, and
// since the *http.ServeMux implements http.Handler, you can create an
// *http.ServeMux instance with multiple related requests and register it with a parent
// *http.ServeMux:
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good puppy!\n"))
	})
	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))
// In this example, a request for /person/greet is handled by handlers attached to
// person, while /dog/greet is handled by handlers attached to dog. When we register
// person and dog with mux, we use the http.StripPrefix helper function to remove
// the part of the path that’s already been processed by mux.

// //Middleware
// One of the most common requirements of an HTTP server is to perform a set of
// actions across multiple handlers, such as checking if a user is logged in, timing a
// request, or checking a request header
// Go handles this with a middleware pattern,the middleware pattern uses
// a function that takes in an http.Handler instance and returns an http.Handler.
// Usually, the returned http.Handler is a closure that is converted to an http.HandlerFunc.
// Here are two middleware generators, one that provides timing of
// requests and another that uses perhaps the worst access controls imaginable:
	func RequestTimer(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			end := time.Now()
			log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
		})
	}

	var securityMsg = []byte("You didn't give the secret password\n")

	func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("X-Secret-Password") != password {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(securityMsg)
					return
				}
				h.ServeHTTP(w, r)
			})
		}
	}
// These two middleware implementations demonstrate what middleware does. First,
// we do setup operations or checks. If the checks don’t pass, we write the output in the
// middleware (usually with an error code) and return. If all is well, we call the handler’s
// ServeHTTP method. When that returns, we run cleanup operations.
// The TerribleSecurityProvider shows how to create configurable middleware. You
// pass in the configuration information (in this case, the password), and the function
// returns middleware that uses that configuration information. It is a bit of a mind
// bender, as it returns a closure that returns a closure.

// We add middleware to our request handlers by chaining them:
	terribleSecurity := TerribleSecurityProvider("GOPHER")
	mux.Handle("/hello", terribleSecurity(RequestTimer(
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	}))))
// We get back our middleware from the TerribleSecurityProvider and then wrap
// our handler in a series of function calls. This calls the terribleSecurity closure first,
// then calls the RequestTimer, which then calls our actual request handler.
// Because the *http.ServeMux implements the http.Handler interface, you can apply
// a set of middleware to all of the handlers registered with a single request router:
	terribleSecurity := TerribleSecurityProvider("GOPHER")
	wrappedMux := terribleSecurity(RequestTimer(mux))
	s := http.Server{
		Addr:
		":8080",
		Handler: wrappedMux,
	}
// // Use idiomatic third-party modules to enhance the server
// Just because the server is production quality doesn’t mean that you shouldn’t use
// third-party modules to improve its functionality. If you don’t like the function chains
// for middleware, you can use a third-party module called alice, which allows you to
// use the following syntax:
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	}
	chain := alice.New(terribleSecurity, RequestTimer).ThenFunc(helloHandler)
	mux.Handle("/hello", chain)
// The biggest weakness in the HTTP support in the standard library is the built-in
// *http.ServeMux request router. It doesn’t allow you to specify handlers based on an
// HTTP verb or header, and it doesn’t provide support for variables in the URL path.
// Nesting *http.ServeMux instances is also a bit clunky. There are many, many projects
// to replace it, but two of the most popular ones are gorilla mux and chi. Both are con‐
// sidered idiomatic because they work with http.Handler and http.HandlerFunc
// instances, demonstrating the Go philosophy of using composable libraries that fit
// together with the standard library. They also work with idiomatic middleware, and
// both projects provide optional middleware implementations of common concerns.


// Just because the server is production quality doesn’t mean that you shouldn’t use
// third-party modules to improve its functionality. If you don’t like the function chains
// for middleware, you can use a third-party module called alice, which allows you to
// use the following syntax:
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	}
	chain := alice.New(terribleSecurity, RequestTimer).ThenFunc(helloHandler)
	mux.Handle("/hello", chain)
// The biggest weakness in the HTTP support in the standard library is the built-in
// *http.ServeMux request router. It doesn’t allow you to specify handlers based on an
// HTTP verb or header, and it doesn’t provide support for variables in the URL path.
// Nesting *http.ServeMux instances is also a bit clunky. There are many, many projects
// to replace it, but two of the most popular ones are gorilla mux and chi. Both are con‐
// sidered idiomatic because they work with http.Handler and http.HandlerFunc
// instances, demonstrating the Go philosophy of using composable libraries that fit
// together with the standard library. They also work with idiomatic middleware, and
// both projects provide optional middleware implementations of common concerns.