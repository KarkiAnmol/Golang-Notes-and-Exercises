package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)
	go func() {
		myChannel <- "plane"
	}()
	go func() {
		anotherChannel <- "dog"
	}()
	// select statement makes the go routine wait on multiple channel operations

	select { // this select statement is going to block until one of the two cases below run
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	}
}
