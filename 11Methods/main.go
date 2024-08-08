package main

import "fmt"

//when function into the structs then we call them method

func main() {

	majid := User{"Majid", "majid@gmail.com", true, 80}
	fmt.Println(majid)
	fmt.Printf("Majid details are: %+v\n", majid)
	fmt.Printf("Majid details are: %+v\n", majid.Age)
	fmt.Printf("Majid details are: %+v\n", majid.Age)

	majid.GetStatus()
	majid.NewMail()
}

//defing structs

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user: ", u.Email)
}
