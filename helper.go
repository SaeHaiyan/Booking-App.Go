package main

import "strings"

func UserValidation (userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(userName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}