package main

import (
	"fmt"
	"strings"
)

func main(){

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// fmt.Println("Welcome to", conferenceName, "Booking App")
	// fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)

	// // fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	// fmt.Println("Get Your tickets now!")

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

		isValidName := len(userName) >= 3
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName)

			fmt.Printf("Thankyou %v for booking %v ticket with us! You will receive a confirmation email soon.\n", userName, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			userNames := []string{}
			for _, booking := range bookings {

				var names = strings.Fields(booking)
				userNames = append(userNames, names[0])
			}
			fmt.Printf("The names of the bookings are: %v\n", userName)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come again Next Year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("Name entered is too short, need to be 3 or more characters")
			}
			if !isValidEmail {
				fmt.Println("Email address not valid, must contain '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of ticket entered is invalid.")
			}
			fmt.Printf("Your input data is invalid, Try Again")
			continue
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v Booking Application!!\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get Your tickets now!")
}
	




	// if userTickets <= remainingTickets {
	// 	remainingTickets -= userTickets
	// 	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email soon.\n", userName, userTickets)
	// 	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)
	// } else {
	// 	fmt.Printf("Sorry, we only have %v tickets remaining. Please try booking fewer tickets.\n", remainingTickets)
	// }

	// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)
 