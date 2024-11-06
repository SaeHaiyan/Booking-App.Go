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

	//Call greetUsers Function
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		//Call For User Input Function
		userName, email, userTickets := getUserInput()

		//Call Validation Function
		isValidName, isValidEmail, isValidTicketNumber := userValidation(userName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{

			//Call for ticketBooking Function
			ticketBooking(remainingTickets, userTickets, bookings, userName, email, conferenceName)

			// Call printUserName Function
			userNames := getUserName(bookings)
			fmt.Printf("The names of the bookings are: %v\n", userNames)

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

func getUserName(bookings []string) []string {
	userNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		userNames = append(userNames, names[0])
	}
	return userNames
}
	
func userValidation (userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(userName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, uint){
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

	return userName, email, userTickets
}

func ticketBooking(remainingTickets uint, userTickets uint, bookings []string, userName string, email string, conferenceName string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName)

	fmt.Printf("Thankyou %v for booking %v ticket with us! You will receive a confirmation email soon.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}