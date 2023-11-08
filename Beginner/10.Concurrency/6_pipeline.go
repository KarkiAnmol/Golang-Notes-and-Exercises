package main

import (
	"fmt"
)

// sliceToChannel function converts the slice into channel
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int) //unbuffered channel
	go func() {
		for _, v := range nums {
			out <- v //after write this go routine waits until the value from out
			// channel is read because it is unbuffered channel (another function 'sq'
			//  reads from this channel so the both goroutines are running concurrently)
		}
		close(out)
	}()
	return out
}

// sq function squares each input value received from the channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	// Input
	nums := []int{1, 2, 3, 4, 5}

	// Stage 1: Convert slice to channel
	dataChannel := sliceToChannel(nums)

	// Stage 2: Square each value from the channel
	finalChannel := sq(dataChannel)

	// Stage 3: Retrieve and print squared values
	for n := range finalChannel {
		fmt.Println(n)
	}
}
