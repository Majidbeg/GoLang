// structs are alternate of classes in golang
// there is no inheritance in golang and no super or parent
package main

import "fmt"

func main() {

	majid := User{"Majid", "majid@gmail.com", true, 80}
	fmt.Println(majid)
	fmt.Printf("Majid details are: %+v\n", majid)
	fmt.Printf("Majid details are: %+v\n", majid.Age)
	fmt.Printf("Majid details are: %+v\n", majid.Age)
}

//defing structs

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
