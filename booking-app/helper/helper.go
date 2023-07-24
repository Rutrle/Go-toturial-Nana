package helper

import "strings"

// velké pismeno - na export jinam
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	/*
			city := "Praha"

		isValidCity := city != "Singapore" && city != "London"
	*/
	return isValidName, isValidEmail, isValidTicketNumber
}
