//similar to other programmes and switch is also smae

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welocme to dice game")

	rand.Seed(time.Now().UnixNano()) //we have seeded so we will the random number every single time
	diceNumber := rand.Intn(6) + 1   // rand.intn(6) will return 1,2,3,4,5 so we are adding 1 so we will get 6
	fmt.Println("value for dice number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot")
	}
}
