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

	// array
	var bookings = [50]string{"Nana", "Nicole", "Peter"}
	var myarray [30]string
	myarray[0] = "Nana"

	fmt.Println(myarray)

	var userName string
	var lastName string
	var userTickets uint
	var email string

	// this is a comment
	fmt.Println("Enter your name")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = userName + " " + lastName

	var bookings_slice []string
	bookings_slice = append(bookings_slice, userName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation on your email %v\n", userName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets %v for %v\n", remainingTickets, conferenceName)

}
