// Important topic slices

package main

import (
	"fmt"
	"sort"
)

//when definignt he value inside the array [2] then it is more of array
// when it is empty then it is called as slices

func main() {
	var fruitList = []string{"Apple", "peach"}
	fmt.Printf("type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "tomato")
	fmt.Println(fruitList) //[Apple peach Mango Banana]

	//going to use Alot
	fruitList = append(fruitList[1:]) // 1: means slice of slices at 1st element
	fmt.Println(fruitList)            //[peach Mango Banana]

	fruitList = append(fruitList[1:3]) // 1:3 means slice of slices at 1st element
	fmt.Println(fruitList)             //[Mango Banana]

	fruitList = append(fruitList[:3]) // 3: means slice of slices at 1st element
	fmt.Println(fruitList)            //[Mango Banana]

	highScore := make([]int, 4) //what kind of datatype and size we are going to have
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore) //sorts and other kind of functions are vailable for slices not for array in go
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted((highScore)))

	//How to remove vlaue from a slice on basis of index

	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2 //need to remove second index form the slice
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
