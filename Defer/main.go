package main

import "fmt"

//defer should be used at very end of an function or LIFO(Last in first out)

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
