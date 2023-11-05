package main

//Program to implement fork-join method
import "fmt"

func main() {
	myNameChannel := make(chan string) //unbuffered open channel
	//goroutine to write a string in the channel
	go func() {
		myNameChannel <- "anmol"
	}()
	name := <-myNameChannel // the main function waits to reads the value from the channel until the
	// goroutine writtes in the empty channel
	fmt.Println(name)
}
