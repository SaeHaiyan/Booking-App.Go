package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your tickets now!")

	for {
		// Ask for user details and number of tickets they want to book
		var userName string
		var email string
		var userTickets uint

		fmt.Println("Enter Your First Name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter Your Email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		// Check if enough tickets are available
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you cannot book %v tickets.\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for booking %v tickets! You will receive a confirmation email at %v.\n", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// Display first names of all bookings
		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			if len(names) > 0 {
				firstNames = append(firstNames, names[0])
			}
		}
		fmt.Printf("The names of the bookings are: %v\n", firstNames)

		// Exit the loop if all tickets are booked
		if remainingTickets == 0 {
			fmt.Println("Our conference is fully booked. Come back next year!")
			break
		}
	}
}
