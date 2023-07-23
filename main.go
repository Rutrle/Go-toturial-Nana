package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for { // same as if it would be "for true"

		// this is a comment
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			fmt.Printf("Bookings so far %v\n", bookings)

			firstNames := getFirstNames(bookings)

			fmt.Printf("First names of all bookings %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0
			// alternatively  noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conferece is sold out, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email appers to be wrong")
			}
			if !isValidTicketNumber {
				fmt.Println("Your number of tickets is invalid")
			}
		}
	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) { // types must be specified
	fmt.Printf("Welcome to our conference %v\n", confName)
	fmt.Printf("we have %v tickets remaining out of %v\n", remainingTickets, confTickets)
	fmt.Println("Get your tickets here ")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
		// _ is used to identfy unused variables
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	/*
			city := "Praha"

		isValidCity := city != "Singapore" && city != "London"
	*/
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Enter your name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets %v for %v\n", remainingTickets, conferenceName)
}

/*
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
*/
