package main

//this is an entry point for the execution
//first line of code

//fmt is packge which will have different inbuilt functions
import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go conference"
	conferenceName := "Go conference" //do not work with const only with var
	const conferenceTickets uint = 50 //value can't be changed
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets) //returns memory addess

	fmt.Printf("conferenceName is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		//or
		//var bookings = []string{}

		var firstName string
		var lastName string
		var email string
		var userTickets uint
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

		if userTickets > remainingTickets {
			fmt.Printf("We have only %v remaiing, so you can't book %v ticktes \n", remainingTickets, userTickets)
			continue //we use for next iteration istead of breaking an loop using break
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("thank you %v %v for booking %v tickets. you wil recieve a confirmation mail at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

		// _ is called blank identifier
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("the first names of bookings are %v\n", firstNames)

		// noTikcetsRemaining := remainingTickets == 0
		if remainingTickets == 0 {
			//end the programme
			fmt.Printf("Our conference is booked out. Come back next year")
			break
		}

	}

}
