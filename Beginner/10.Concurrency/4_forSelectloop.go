package main

import "fmt"

//for select loop is just using a select inside the iterations of a for loop
func main() {
	buffChannel := make(chan string, 3) //Buffered channel with limit 3
	chars := []string{"a", "b", "c"}
	for _, v := range chars {
		v := v
		select {
		case buffChannel <- v:
		}
	}
	close(buffChannel) //closing the channel since writing is done on it
	for result := range buffChannel {
		fmt.Println(result) // even though the channel is closed,read on an non-empty channel is valid
	}
}
