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
// channel. Remember that a read on an open channel pauses until there is data avail‐
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