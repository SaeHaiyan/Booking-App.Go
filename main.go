package main

import (
	"fmt"
)

func main(){

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// fmt.Println("Welcome to", conferenceName, "Booking App")
	fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)

	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get Your tickets now!")

	for {
		// Ask for user details and number of tickets they want to book
		var userName string
		var email string
		var userTickets uint

		// This code lines is for Manual Input 
		// userName ="Haiyan"
		// userTickets = 4
		// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		fmt.Println("Enter Your First Name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter Your Email: ")
		fmt.Scan(&email)

		fmt.Println("Enter The Number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName)

		fmt.Printf("Thankyou %v for booking %v ticket with us! You will receive a confirmation email soon.\n", userName, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)


		fmt.Printf("These are all our Bookings: %v\n", bookings)

		 
	}
}




	// if userTickets <= remainingTickets {
	// 	remainingTickets -= userTickets
	// 	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email soon.\n", userName, userTickets)
	// 	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)
	// } else {
	// 	fmt.Printf("Sorry, we only have %v tickets remaining. Please try booking fewer tickets.\n", remainingTickets)
	// }

	// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)
 