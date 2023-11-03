/*

Concurrency is the computer science term for breaking up a single process into inde‐
pendent components and specifying how these components safely share data.

The concurrency model of GO is focued on CSP(Communicating Sequential Processes)

Features that are the backbone of concurrency in Go:
	-> goroutines, channels, and the select keyword


*/

// When to Use Concurrency
// concurrency is not parallelism
// Concurrency is a tool to better
// structure the problem you are trying to solve. Whether or not concurrent code runs
// in parallel (at the same time) depends on the hardware and if the algorithm allows it.

/*
Goroutines

The goroutine is the core concept in Go’s concurrency model.
A process is an instance of a program that’s being run by a computer’s operating system
The operating system associates some resources, such as memory, with the process and makes sure that
other processes can’t access them. A process is composed of one or more threads. A
thread is a unit of execution that is given some time to run by the operating system.
Threads within a process share access to resources.

A CPU can execute instructions from one or more threads at the same time, depending on the number of cores.

Goroutines are lightweight processes managed by the Go runtime.
When a GO program starts, the Go runtime creates a number of threads and launches a single goroutine to run your program.
all of the goroutines created by your program,including the initial one, are assigned to these threads automatically by the
Go runtime scheduler,just as OS schedules threads across CPU cores.

This might seem like extra work, since the underlying operating system already includes a
scheduler that manages threads and processes, but it has several benefits:

	• Goroutine creation is faster than thread creation, because you aren’t creating an
	operating system–level resource.
	• Goroutine initial stack sizes are smaller than thread stack sizes and can grow as
	needed. This makes goroutines more memory efficient.
	• Switching between goroutines is faster than switching between threads because it
	happens entirely within the process, avoiding operating system calls that are (rel‐
	atively) slow.
	• The scheduler is able to optimize its decisions because it is part of the Go process.
	The scheduler works with the network poller, detecting when a goroutine can be
	unscheduled because it is blocking on I/O. It also integrates with the garbage col‐
	lector, making sure that work is properly balanced across all of the operating sys‐
	tem threads assigned to your Go process.

A goroutine is launched by placing the go keyword before a function invocation. Just
like any other function, you can pass it parameters to initialize its state. However, any
values returned by the function are ignored.

 goroutines in Go are ultimately executed by underlying operating system threads, but they're
 managed in a different way by the Go runtime. The Go runtime abstracts the complexity of managing
 threads and allows you to work with goroutines more easily.
 It multiplexes multiple goroutines onto a smaller number of OS threads.
*/
func process(val int) int {
	// do something with val
}
func runThingConcurrently(in <-chan int, out chan<- int) {
	go func() { //goroutine
		for val := range in {
			result := process(val)
			out <- result
		}
	}()
}

//Channels
/*
Goroutines communicate using channels. Like slices and maps, channels are a built-in
type created using the make function: */
ch := make(chan int)
/*
Like maps, channels are reference types. When you pass a channel to a function, you
are really passing a pointer to the channel. Also like maps and slices, the zero value
for a channel is nil.
*/

//Reading, Writing, and Buffering
a := <-ch // reads a value from ch and assigns it to a
ch <- b   // write the value in b to ch

/*
It is rare for a goroutine to read and write to the same channel. When assigning a
channel to a variable or field, or passing it to a function, use an arrow before the chan
keyword (ch <-chan int) to indicate that the goroutine only reads from the channel.
Use an arrow after the chan keyword (ch chan<- int) to indicate that the goroutine
only writes to the channel. Doing so allows the Go compiler to ensure that a channel
is only read from or written by a function.
*/
