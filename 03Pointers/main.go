package main

import (
	"fmt"
)

//varibales got stored in the memory location and sometime copy got created
// Pointers gives us the surety whatever resource your are passing on the pointers
// We will pass the memo address of the varibale.
// We use pointers to avoid irregularities

func main() {
	fmt.Println("Welocme to pointers")

	// myPointers is an variblaes which is an holding pointer as datatype which is type of int
	// Default value of an pointer in nill
	// & reference the pointer

	// var myPointers *int
	// fmt.Println("value of pointer is", myPointers)

	var number int = 23
	var ptr = &number
	fmt.Println("value of pointer is", ptr)  // it will return memory adderress
	fmt.Println("value of pointer is", *ptr) // it will return 23

	*ptr = *ptr + 2
	fmt.Println("the value is: ", number)
	fmt.Println("the value is: ", &number)

}
