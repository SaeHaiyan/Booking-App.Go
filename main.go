package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main(){

	//Call greetUsers Function
	greetUsers()

	for {

		//Call For User Input Function
		userName, email, userTickets := getUserInput()

		//Call Validation Function
		isValidName, isValidEmail, isValidTicketNumber := helper.UserValidation(userName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{

			//Call for ticketBooking Function
			ticketBooking(userTickets, userName, email)

			// Call getUserName Function
			userNames := getUserName()
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

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your tickets now!")
}

func getUserName() []string {
	userNames := []string{}
	for _, booking := range bookings {

		userNames = append(userNames, booking["username"])
	}
	return userNames
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

func ticketBooking(userTickets uint, userName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create a map for user 
	var userData = make(map[string]string) 
	userData["username"] = userName
	userData["email"] = email
	userData["ticketNumber"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v for booking %v ticket with us! You will receive a confirmation email soon.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}