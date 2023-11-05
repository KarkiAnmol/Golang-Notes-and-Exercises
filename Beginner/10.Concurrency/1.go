package main

import (
	"fmt"
	"time"
)

func someFunc(num int) {
	fmt.Println("goroutine:", num)
}
func main() {
	//the three go functions are asynchronous
	go someFunc(1) // this function forks off of the main function and starts running and main continues to execute it's other task
	go someFunc(2)
	go someFunc(3)
	time.Sleep(time.Second * 2) // this makes main wait for 2 second before jumping onto next line
	fmt.Println("inside main function")
}
