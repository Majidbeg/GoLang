package main

import "fmt"

//can't call function inside the function in go lang
// anonymus function also existing in go
func main() {
	fmt.Println("Welocme to Functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("My message is: ", myMessage)

}

// Need to define param type
func adder(valOne int, valTwo int) int { //int after bracket is function signature what function is returning
	return valOne + valTwo
}

//all the vuales are type of int
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "My name is majid"
}

func greeter() {
	fmt.Println("Namstey from go lang")
}
