// Concurrency Practices and Patterns

//Keep Your APIs Concurrency-Free
/*
Concurrency is an implementation detail, and good API design should hide imple‐
mentation details as much as possible.

Practically, this means that you should never expose channels or mutexes in your
API’s types, functions, and methods
If you expose a channel, you put the
responsibility of channel management on the users of your API. This means that the
users now have to worry about concerns like whether or not a channel is buffered or
closed or nil. They can also trigger deadlocks by accessing channels or mutexes in an
unexpected order.
*/

//Goroutines, for Loops, and Varying Variables
/*Most of the time, the closure that you use to launch a goroutine has no parameters.
Instead, it captures values from the environment where it was declared. There is one
common situation where this doesn’t work: when trying to capture the index or value
of a for loop*/
func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
// output:
20
20
20
20
20

// The last value assigned to v was 10. When the goroutines run, that’s
// the value that they see. This problem isn’t unique to for loops; any time a goroutine
// depends on a variable whose value might change, you must pass the value into the
// goroutine. There are two ways to do this. 

// The first is to shadow the value within the loop:
for _, v := range a {
	v := v
	go func() {
		ch <- v * 2
	}()
}
// If you want to avoid shadowing and make the data flow more obvious, you can also
// pass the value as a parameter to the goroutine:
for _, v := range a {
	go func(val int) {
		ch <- val * 2
	}(v)
}
//Any time your goroutine uses a variable whose value might change,
// pass the current value of the variable into the goroutine.


//Always Clean Up Your Goroutines

// Goroutine leak: 
// Unlike variables, the Go runtime can’t detect that a goroutine will never be
// used again. If a goroutine doesn’t exit, the scheduler will still periodically give it time
// to do nothing, which slows down your program. 


//The Done Channel Pattern
//This pattern provides a way to signal a goroutine that it’s time to stop processing.
//It uses a channel to signal that it’s time to exit
//an example where we pass the same data to multiple functions, 
// but only want the result from the fastest function:
func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
				case result <- searcher(s):
				case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}
// In our function, we declare a channel named done that contains data of type
// struct{}. We use an empty struct for the type because the value is unimportant; we
// never write to this channel, only close it. We launch a goroutine for each searcher
// passed in. The select statements in the worker goroutines wait for either a write on
// the result channel (when the searcher function returns) or a read on the done
// channel.
//  Remember that a read on an open channel pauses until there is data avail‐
// able and that a read on a closed channel always returns the zero value for the channel.
// This means that the case that reads from done will stay paused until done is closed. In
// searchData, we read the first value written to result, and then we close done. This
// signals to the goroutines that they should exit, preventing them from leaking.

//Using a Cancel Function to Terminate a Goroutine
The function must be called after
the for loop:
func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
				case <-done:
					return
				case ch<-i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}
func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
				break
		}
		fmt.Println(i)
	}
	cancel()
}
//The countTo function creates two channels, one that returns data and another for sig‐
// naling done. Rather than return the done channel directly, we create a closure that
// closes the done channel and return the closure instead. Cancelling with a closure
// allows us to perform additional clean-up work, if needed.


//When to Use Buffered and Unbuffered Channels
/*
By default, channels are unbuffered, and they are
easy to understand: one goroutine writes and waits for another goroutine to pick up
its work, like a baton in a relay race.

Buffered channels are much more complicated.
You have to pick a size, since buffered channels never have unlimited buffers.

Proper use of a buffered channel means that you must handle the case where the buffer is full
and your writing goroutine blocks waiting for a reading goroutine.

Buffered channels are useful when you know how many goroutines you have
launched, want to limit the number of goroutines you will launch, or want to limit the
amount of work that is queued up.

Buffered channels work great when you want to either gather data back from a set of
goroutines that you have launched or when you want to limit concurrent usage.
*/
// In our first example, we are processing the first 10 results on a channel. To do this, we
// launch 10 goroutines, each of which writes its results to a buffered channel:
func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <- ch // read from input channel into v and process the 'v' to populate the channel 10 times
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results) // read from buffered channel of limit 10 and append into a slice
	}
	return out
}
//When all of the values have been read, we return
// the results, knowing that we aren’t leaking any goroutines.


//Backpressure
// Another technique that can be implemented with a buffered channel is backpressure.
// It is counterintuitive, but systems perform better overall when their components limit
// the amount of work they are willing to perform. We can use a buffered channel and a
// select statement to limit the number of simultaneous requests in a system:

	//a struct that contains a buffered channel with a number of “tokens” 
	type PressureGauge struct {
		ch chan struct{}
	}

	func New(limit int) *PressureGauge {
		ch := make(chan struct{}, limit)
		for i := 0; i < limit; i++ {
			ch <- struct{}{}
		}
		return &PressureGauge{
			ch: ch,
		}
	}
	//Every time a goroutine wants to use the function, it calls Process

	func (pg *PressureGauge) Process(f func()) error {
		//The select tries to read a token from the channel. If it can, the function runs, and the token is returned to the buffered channel
		select { 
			case <-pg.ch:
				f()
				pg.ch <- struct{}{}
				return nil
			default:
				return errors.New("no more capacity")
		}
	}

// Here’s a quick example that
// uses this code with the built-in HTTP server
	func doThingThatShouldBeLimited() string {
			time.Sleep(2 * time.Second)
			return "done"
	}
	func main() {
		pg := New(10)
		http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
			err := pg.Process(func() {
				w.Write([]byte(doThingThatShouldBeLimited()))
			})
			if err != nil {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Too many requests"))
			}
		})
		http.ListenAndServe(":8080", nil)
	}


//Turning Off a case in a select
/*
If one of the
cases in a select is reading a closed channel, it will always be successful, returning
the zero value. Every time that case is selected, you need to check to make sure that
the value is valid and skip the case. If reads are spaced out, your program is going to
waste a lot of time reading junk values.*/

// When you detect that a channel has been closed, set the
// channel’s variable to nil. The associated case will no longer run, because the read
// from the nil channel never returns a value:

// in and in2 are channels, done is a done channel.
for {
	select {
		case v, ok := <-in:
			if !ok {
				in = nil // the case will never succeed again!
				continue
			}

		// process the v that was read from in
		case v, ok := <-in2:
			if !ok {
				in2 = nil // the case will never succeed again!
				continue
		}
		// process the v that was read from in2
		case <-done:
			return
	}
}

//How to Time Out Code
/*
One of the things that we can do with concurrency in Go is manage how much time a
request (or a part of a request) has to run
*/
func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
		case <-done:
			return result, err
		case <-time.After(2 * time.Second):
			return 0, errors.New("work timed out")
	}
}
/*Any time you need to limit how long an operation takes in Go, you’ll see a variation
on this pattern. We have a select choosing between two cases. The first case takes
advantage of the done channel pattern we saw earlier. We use the goroutine closure to
assign values to result and err and to close the done channel. If the done channel
closes first, the read from done succeeds and the values are returned.
The second channel is returned by the After function in the time package. It has a
value written to it after the specified time.Duration has passed*/


//Using WaitGroups
// Sometimes one goroutine needs to wait for multiple goroutines to complete their
// work. If you are waiting for a single goroutine, you can use the done channel pattern
// that we saw earlier. But if you are waiting on several goroutines, you need to use a
// WaitGroup, which is found in the sync package in the standard library.
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
		}()
	wg.Wait()
}
/*
A sync.WaitGroup doesn’t need to be initialized, just declared, as its zero value is use‐
ful. There are three methods on sync.WaitGroup: 
1) Add, which increments the counter
of goroutines to wait for; 
2) Done, which decrements the counter and is called by a
goroutine when it is finished; and 
3) Wait, which pauses its goroutine until the counter
hits zero. Add is usually called once, with the number of goroutines that will be
launched. Done is called within the goroutine. To ensure that it is called, even if the
goroutine panics, we use a defer
*/

// You’ll notice that we don’t explicitly pass the sync.WaitGroup. There are two reasons.
// The first is that you must ensure that every place that uses a sync.WaitGroup is using
// the same instance. If you pass the sync.WaitGroup to the goroutine function and
// don’t use a pointer, then the function has a copy and the call to Done won’t decrement
// the original sync.WaitGroup. By using a closure to capture the sync.WaitGroup, we
// are assured that every goroutine is referring to the same instance.

// The second reason is design. Remember, you should keep concurrency out of your
// API. As we saw with channels earlier, the usual pattern is to launch a goroutine with a
// closure that wraps the business logic. The closure manages issues around concur‐
// rency and the function provides the algorithm.

// when you have
// multiple goroutines writing to the same channel, you need to make sure that the
// channel being written to is only closed once. A sync.WaitGroup is perfect for this.
// Let’s see how it works in a function that processes the values in a channel concur‐
// rently, gathers the results into a slice, and returns the slice:
func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
		defer wg.Done()
		for v := range in {
			out <- processor(v)
		}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}
//In our example, we launch a monitoring goroutine that waits until all of the process‐
// ing goroutines exit. When they do, the monitoring goroutine calls close on the out‐
// put channel. The for-range channel loop exits when out is closed and the buffer is
// empty. Finally, the function returns the processed values.

// Use WaitGroups only when you have something to clean up (like closing a chan‐
	// nel they all write to) after all of your worker goroutines exit.


//Running Code Exactly Once
// init should be
// reserved for initialization of effectively immutable package-level state. However,
// sometimes you want to lazy load, or call some initialization code exactly once after
// program launch time. This is usually because the initialization is relatively slow and
// may not even be needed every time your program runs. The sync package includes a
// handy type called Once that enables this functionality.
	type SlowComplicatedParser interface {
		Parse(string) string
	}
	var parser SlowComplicatedParser
	var once sync.Once
	func Parse(dataToParse string) string {
		once.Do(func() {
			parser = initParser()
		})
		return parser.Parse(dataToParse)
	}
	func initParser() SlowComplicatedParser {
	// do all sorts of setup and loading here
	}
/*
We have declared two package-level variables, parser, which is of type Slow
ComplicatedParser, and once, which is of type sync.Once. Like sync.WaitGroup, we
do not have to configure an instance of sync.Once (this is called making the zero
value useful). Also like sync.WaitGroup, we must make sure not to make a copy of an
instance of sync.Once, because each copy has its own state to indicate whether or not
it has already been used. Declaring a sync.Once instance inside a function is usually
the wrong thing to do, as a new instance will be created on every function call and
there will be no memory of previous invocations.
In our example, we want to make sure that parser is only initialized once, so we set
the value of parser from within a closure that’s passed to the Do method on once. If
Parse is called more than once, once.Do will not execute the closure again.*/


// Putting Our Concurrent Tools Together
// We have a function
// that calls three web services. We send data to two of those services, and then take the
// results of those two calls and send them to the third, returning the result. The entire
// process must take less than 50 milliseconds, or an error is returned.
// We’ll start with the function we invoke:
func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC: make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2),
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

/*The first thing we do is set up a context that times out in 50 milliseconds. When
there’s a context available, use its timer support rather than calling time.After. One
of the advantages of using the context’s timer is that it allows us to respect timeouts
that are set by the functions that called this function. We talk about the context in
Chapter 12 and cover using timeouts in detail in “Timers” on page 261. For now, all
you need to know is that reaching the timeout cancels the context. The Done method
on the context returns a channel that returns a value when the context is canceled,
either by timing out or by calling the context’s cancel method explicitly.
After we create the context, we use a defer to make sure the context’s cancel func‐
tion is called. As we’ll discuss in “Cancellation” on page 258, you must call this func‐
tion or resources leak.
We then populate a processor instance with a series of channels that we’ll use to
communicate with our goroutines. Every channel is buffered, so that the goroutines
that write to them can exit after writing without waiting for a read to happen. (The
errs channel has a buffer size of two, because it could potentially have two errors
written to it.)
The processor struct looks like this:*/
type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC chan CIn
	errs chan error
}
	// Next, we call the launch method on processor to start three goroutines: one to call
	// getResultA, one to call getResultB, and one to call getResultC:
	func (p *processor) launch(ctx context.Context, data Input) {
		go func() {
			aOut, err := getResultA(ctx, data.A)
			if err != nil {
				p.errs <- err
				return
			}
			p.outA <- aOut
		}()
		go func() {
			bOut, err := getResultB(ctx, data.B)
			if err != nil {
				p.errs <- err
				return
			}
			p.outB <- bOut
		}()
		go func() {
			select {
				case <-ctx.Done():
					return
				case inputC := <-p.inC:
					cOut, err := getResultC(ctx, inputC)
					if err != nil {
						p.errs <- err
						return
					}
					p.outC <- cOut
			}
		}()
	}
		// The goroutines for getResultA and getResultB are very similar. They call their
		// respective methods. If an error is returned, they write the error to the p.errs chan‐
		// nel. If a valid value is returned, they write the value to their channels (p.outA for
		// getResultA and p.outB for getResultB).
		// Since the call to getResultC only happens if the calls to getResultA and getResultB
		// succeed and happen within 50 milliseconds, the third goroutine is slightly more com‐
		// plicated. It contains a select with two cases. The first is triggered if the context is
		// canceled. The second is triggered if the data for the call to getResultC is available. If
		// the data is available, the function is called, and the logic is similar to the logic for our
		// first two goroutines.
		// After the goroutines are launched, we call the waitForAB method on processor:
		func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
			var inputC CIn
			count := 0
			for count < 2 {
				select {
					case a := <-p.outA:
							inputC.A = a
							count++
					case b := <-p.outB:
						inputC.B = b
						count++
					case err := <-p.errs:
						return CIn{}, err
					case <-ctx.Done():
						return CIn{}, ctx.Err()
				}
			}
			return inputC, nil
		}
// 		This uses a for-select loop to populate inputC, an instance of CIn, the input param‐
// 		eter for getResultC. There are four cases. The first two read from the channels
// 		written to by our first two goroutines and populate the fields in inputC. If both of
// these cases execute, we exit the for-select loop and return the value of inputC and a
// nil error.
// The next two cases handle error conditions. If an error was written to the p.errs
// channel, we return the error. If the context has been canceled, we return an error to
// indicate the request is canceled.
// Back in GatherAndProcess, we perform a standard nil check on the error. If all is
// well, we write the inputC value to the p.inC channel and then call the waitForC
// method on processor:
func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
		case out := <-p.outC:
			return out, nil
		case err := <-p.errs:
			return COut{}, err
		case <-ctx.Done():
			return COut{}, ctx.Err()
	}
}
// This method consists of a single select. If getResultC completed successfully, we
// read its output from the p.outC channel and return it. If getResultC returned an
// error, we read the error from the p.errs channel and return it. Finally, if the context
// has been canceled, we return an error to indicate that. After waitForC completes,
// GatherAndProcess returns the result to its caller.
// If you trust the author of getResultC to do the right thing, this code can be simpli‐
// fied. Since the context is passed into getResultC, the function can be written to
// respect the timeout and return an error if it was triggered. In that case, we can call
// getResultC directly in GatherAndProcess. This eliminates the inC and outC channels
// from process, a goroutine from launch, and the entire waitForC method. The gen‐
// eral principle is to use as little concurrency as your program needs to be correct.
// By structuring our code with goroutines, channels, and select statements, we sepa‐
// rate the individual steps, allow independent parts to run and complete in any order,
// and cleanly exchange data between the dependent parts. In addition, we make sure
// that no part of the program hangs, and we properly handle timeouts set both within
// this function and from earlier functions in the call history. If you are not convinced
// that this is a better method for implementing concurrency, try to implement this in
// another language. You might be surprised at how difficult it is.

//When to use Mutexes instead of channels
