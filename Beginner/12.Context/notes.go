// 	-------------------THE CONTEXT----------------------------
// Servers need a way to handle metadata on individual requests. This metadata falls
// into two general categories:
// -> metadata that is required to correctly process the request,
// -> metadata on when to stop processing the request
//
// For example, an HTTP server might want to use a tracking ID to identify a chain of requests through a set of microservices.
// It also might want to set a timer that ends requests to other microservices if they take too long.
// Many languages use threadlocal variables to store this kind of
// information, associating data to a specific operating system thread of execution. This
// does’t work in Go because goroutines don’t have unique identities that can be used to
// look up values. More importantly, threadlocals feel like magic; values go in one place
// and pop up somewhere else.
// Go solves the request metadata problem with a construct called the context. Let’s see
// how to use it correctly.

//What Is the Context?
// a context is simply an instance that meets the Context interface defined in the context package.
//
// there is another Go convention that the context is explicitly passed through your program as the first parameter of a function.
// The usual name for the context parameter is ctx:
func logic(ctx context.Context, info string) (string, error) {
	// do some interesting stuff here
	return "", nil
}

// the context package also contains several factory functions for creating and wrapping contexts.
// When you don’t have an existing context, such as at the entry point to a command-line program,
// create an empty initial context with the function context.Background. This returns a variable
// of type context.Context. (Yes, this is an exception to the usual pattern of returning a concrete type from a function call.)

// An empty context is a starting point; each time you add metadata to the context, you
// do so by wrapping the existing context using one of the factory functions in the
// context package:

// context.TODO, that also creates an
// empty context.Context. It is intended for temporary use during
// development.Production code shouldn’t include
// context.TODO.

// There are two context-related methods on
// http.Request:
// • Context returns the context.Context associated with the request.
// • WithContext takes in a context.Context and returns a new http.Request with
//   the old request’s state combined with the supplied context.Context.
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// firstly, we extract the existing context from the request using the Context method
		ctx := req.Context()
		// wrap the context with stuff -- we'll see how soon!

		// After we put values into the context(ctx), we create a new request(req) based on the old request
		//  and the now-populated context using the WithContext method
		req = req.WithContext(ctx)
		// Finally, we call the handler and pass it our new request(req) and
		// the existing http.ResponseWriter(rw).
		handler.ServeHTTP(rw, req)
	})
}

// When you get to the handler, you extract the context from the request using the
// Context method and call your business logic with the context as the first parameter,
// just like we saw previously:
func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

// There’s one more situation where you use the WithContext method: when making an
// HTTP call from your application to another HTTP service. Just like we did when
// passing a context through middleware, you set the context on the outgoing request
// using WithContext:
type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}
	// do the rest of the stuff to process the response
	id, err := processResponse(resp.Body)
	return id, err
}

// Cancellation
// Imagine that you have a request that spawns several goroutines, each one calling a
// different HTTP service. If one service returns an error that prevents you from returning a valid result,
// there is no point in continuing to process the other goroutines. In
// Go, this is called cancellation and the context provides the mechanism for
// implementation.

// To create a cancellable context, use the context.WithCancel function.
// It takes in a context.Context as a parameter and returns a context.Context and a context.CancelFunc.
// The returned context.Context is not the same context that
// was passed into the function. Instead, it is a child context that wraps the passed-in
// parent context.Context. A context.CancelFunc is a function that cancels the con‐
// text, telling all of the code that’s listening for potential cancellation that it’s time to
// stop processing.

// A context is treated as an immutable instance. Whenever we add information to a context, we do so
// by wrapping an existing parent context with a child context. This allows us to use contexts
// to pass information into deeper layers of the code. The context is never used to pass information
// out of deeper layers to higher layers.
// servers.go
// slowServer simulates a slow HTTP server response.
func slowServer() *httptest.Server {
    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(2 * time.Second)
        w.Write([]byte("Slow response"))
    }))
    return s
}

// fastServer simulates a fast HTTP server response with optional error handling.
func fastServer() *httptest.Server {
    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Query().Get("error") == "true" {
            w.Write([]byte("error"))
            return
        }
        w.Write([]byte("ok"))
    }))
    return s
}
// The functions set up two HTTP servers. One sleeps for two seconds, returning "Slow response," while 
// the other checks for a query parameter "error." If true, it responds with "error"; otherwise,
//  it returns "ok." The code uses httptest.Server for testing interactions with remote servers
//   within the same program. The next step involves writing the client part in client.go.

// client.go
var client = http.Client{}

// callBoth initiates calls to both slow and fast servers with cancellation support.
func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    var wg sync.WaitGroup
    wg.Add(2)

    // Concurrently call slow and fast servers
    go func() {
        defer wg.Done()
        err := callServer(ctx, "slow", slowURL)
        if err != nil {
            cancel()
        }
    }()

    go func() {
        defer wg.Done()
        err := callServer(ctx, "fast", fastURL+"?error="+errVal)
        if err != nil {
            cancel()
        }
    }()

    // Wait for both calls to complete
    wg.Wait()
    fmt.Println("done with both")
}

// callServer performs an HTTP call and handles errors and cancellations.
func callServer(ctx context.Context, label string, url string) error {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        fmt.Println(label, "request err:", err)
        return err
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(label, "response err:", err)
        return err
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(label, "read err:", err)
        return err
    }

    result := string(data)
    if result != "" {
        fmt.Println(label, "result:", result)
    }

    if result == "error" {
        fmt.Println("cancelling from", label)
        return errors.New("error happened")
    }

    return nil
}
// In this code segment, the callBoth function is central. It establishes a 
// cancellable context and a corresponding cancel function from the provided
//  context. Two concurrent processes are then launched, each utilizing the 
//  cancellable context, a label, and a URL when calling the callServer function. 
//  The program waits for both processes to complete, invoking the cancel function upon any reported error.
// main.go

// main.go
// main initializes servers, creates a context, and calls the clients with the context and server URLs.
func main() {
    ss := slowServer()
    defer ss.Close()

    fs := fastServer()
    defer fs.Close()

    ctx := context.Background()
    callBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
	
	
// Here’s what happens if you run without an error:
	$ make run-ok
	go build
	./context_cancel false
	fast result: ok
	slow result: Slow response
	done with both
// And here’s what happens if an error is triggered:
	$ make run-cancel
	go build
	./context_cancel true
	fast result: error
	cancelling from fast
	slow response err: Get "http://127.0.0.1:38804": context canceled
	done with both


//Context Example

// First runs a set of Queries and returns the result from the first to respond,cancelling the others
func First(ctx context.Context,urls []string) (*Result,error){
    c:= make(chanResult,len(urls))
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    search := func (url string){
        c <- runQuery(ctx,url)
    }
}