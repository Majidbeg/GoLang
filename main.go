package main

//this is an entry point for the execution
//first line of code

//fmt is packge which will have different inbuilt functions
import "fmt"

func main() {
	// var conferenceName string = "Go conference"
	conferenceName := "Go conference" //do not work with const only with var
	const conferenceTickets uint = 50 //value can't be changed
	var remainingTickets uint = 50
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets) //returns memory addess

	fmt.Printf("conferenceName is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask the user for their nme

	// fmt = fromate package
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //passing the memory address
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("thank you %v %v for booking %v tickets. you wil recieve a confirmation mail at %v", firstName, lastName, userTickets, email)

}
