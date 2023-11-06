package main

import (
	"fmt"
	"time"
)

func main() {
	//parse function converts from a string to a time.Time
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	//format converts time.Time to a string
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
}
