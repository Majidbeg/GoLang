package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//taking input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax or err err
	// input, err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Println("Types of this rating is %T", input)
}
