package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50

	myvar := "Walrus is great"
	fmt.Println(myvar)

	// I can specify the type, but it's better to only do it only when it would be different than the type Go would infer
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here ")

	var userName string
	var userTickets int

	// this is a comment

	userName = "John Doe"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets", userName, userTickets)

}
