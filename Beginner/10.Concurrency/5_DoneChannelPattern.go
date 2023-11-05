package main

import (
	"fmt"
	"time"
)

//Goroutine leak
// example of a goroutine leak would be a goroutine that's running in the background
// even though the main program is closed and it's actually consuming resources

//one way to avoid goroutine leak is to implement mechanism that will allow the parent goroutine to cancel
// its children and thats where our Done Channel pattern can help

func doWork(done <-chan bool) { //passing done channel as read only channel
	for {
		select {
		//if we have recieved from the done channel we can just return
		case <-done:
			return
		//this default case runs until the parent goroutine(i.e. main) shuts down done channel
		// using close(done)
		default:
			fmt.Println("Some Work....")
		}
	}
}

func main() {
	done := make(chan bool)
	//Infinitely running goroutine
	go doWork(done)

	time.Sleep(time.Second * 5) //letting the goroutine to continue to do work for 5 seconds
	close(done)                 // killing the goroutine doWork

}
