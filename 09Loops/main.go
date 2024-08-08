package main

import "fmt"

func main() {

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// i returns the index not the actual value
	// for i := range days {
	// 	fmt.Println((days[i]))
	// }

	//for each kind of loop
	// for _, value := range days {
	for index, value := range days {
		fmt.Printf("Index is %v and Value is %v \n", index, value)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue //contnue just skip the one phase and continue other value
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

lco: //defineig the goto var
	fmt.Println("Hey i am your goto")

}
