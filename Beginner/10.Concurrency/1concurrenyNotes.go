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
other processes can’t access them. A process is composed of one or more threads. 
A jthread is a unit of execution that is given some time to run by the operating system.
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
channel to a variable or field, or passing it to a function,
*/
(ch <-chan int) // use an arrow before the chankeyword to indicate that the goroutine only reads from the channel.
(ch chan<- int) // Use an arrow after the chan keyword to indicate that the goroutine only writes to the channel.
//Doing so allows the Go compiler to ensure that a channel is only read from or written by a function


//you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines.
/*
By default channels are unbuffered. 
Every write to an open, unbuffered channel causes the writing goroutine to pause until another goroutine reads 
from the same channel. and vice-versa 
*/


/*
Go also has buffered channels. These channels buffer a limited number of writes
without blocking. If the buffer fills before there are any reads from the channel, a sub‐
sequent write to the channel pauses the writing goroutine until the channel is read.
Just as writing to a channel with a full buffer blocks, reading from a channel with an
empty buffer also blocks.
*/

//A buffered channel is created by specifying the capacity of the buffer when creating the channel:
ch := make(chan int, 10)

// len to find out how many values are currently in the buffer 
// cap to find out the maximum buffer size.

//For-range and channels 
	for v := range ch {
		fmt.Println(v)
	}
	// Unlike other for-range loops, there is only a single variable declared for the channel,
	// which is the value. The loop continues until the channel is closed, or until a break or
	// return statement is reached.


//Closing a Channel
//When you are done writing to a channel, you close it using the built-in close function:
close(ch)
//Once a channel is closed, any attempts to write to the channel or close the channel again will panic.
//Interestingly, attempting to read from a closed channel always suc‐
// ceeds. If the channel is buffered and there are values that haven’t been read yet, they
// will be returned in order. If the channel is unbuffered or the buffered channel has no
// more values, the zero value for the channel’s type is returned.


// Be aware that closing a channel is only required if there is a goroutine wait‐
// ing for the channel to close (such as one using a for-range loop to read from the
// channel). Since a channel is just another variable, Go’s runtime can detect channels
// that are no longer used and garbage collect them.


//we use the comma ok
// idiom to detect whether a channel has been closed or not:
v, ok := <-ch
// If ok is set to true, then the channel is open. If it is set to false, the channel is closed.


//How Channels Behave
// Channels have many different states, each with a different behavior when reading, writing, or closing.
// check the image(HowChannelsBehave.png) out to see the table of different behaviors


//select
/*
if you can perform two concurrent operations, which one do you do first? You
can’t favor one operation over others, or you’ll never process some cases. This is
called starvation.
*/

//The select keyword allows a goroutine to read from or write to one of a set of multiple channels.
select {
	case v := <-ch:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case ch3 <- x:
		fmt.Println("wrote", x)
	case <-ch4:
		fmt.Println("got value on ch4, but ignored it")
}
//Like a switch, each case in a select creates its own block
//Each case in a select is a read or a write to a channel.
//it picks randomly from any of its cases that can go forward; order is unimportant.
// It also cleanly resolves the starva‐
// tion problem, as no case is favored over another and all are checked at the same time.

/*
If you have
two goroutines that both access the same two channels, they must be accessed in the
same order in both goroutines, or they will deadlock. This means that neither one can
proceed because they are waiting on each other. If every goroutine in your Go appli‐
cation is deadlocked, the Go runtime kills your program

so random choosing from select solves deadlock problem too*/

Example 10-1. Deadlocking goroutines
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v // ch1 will be written with value 1
		v2 := <-ch2 // attempt of read from unbuffered open channel will pause until something is written
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v // write attempt on ch2
	v2 := <-ch1 // read attempt from ch1 will be paused until something is written on ch1
	fmt.Println(v, v2)
}
// output
fatal error: all goroutines are asleep - deadlock!


Example 10-2. Using select to avoid deadlocks
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
		case ch2 <- v:
		case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}
// If you run this program on The Go Playground you’ll get the output:
2 1
// The goroutine that we launched wrote the value 1 into ch1, so the read from ch1 into v2 in
// the main goroutine is able to succeed.
// Since select is responsible for communicating over a number of channels, it is often
// embedded within a for loop:
for {
	select {
		case <-done:
			return
		case v := <-ch:
			fmt.Println(v)
	}
}
// If you want to implement a nonblocking read or write on a channel,
// use a select with a default. The following code does not wait if there’s no value to
// read in ch; it immediately executes the body of the default:
select {
	case v := <-ch:
		fmt.Println("read from ch:", v)
	default:
		fmt.Println("no value written to ch")
}





