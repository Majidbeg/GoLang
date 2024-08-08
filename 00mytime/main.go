package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welocme to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// why we are using 01-02-2006 monday
	// have to use this syntax if we need to use
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //this will return current date time

	createDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, presentTime.Location())
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
