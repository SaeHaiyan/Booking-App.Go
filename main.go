package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName string
	email string
	ticketNumber uint
}

var wg = sync.WaitGroup{}

func main(){

	//Call greetUsers Function
	greetUsers()


	//Call For User Input Function
	userName, email, userTickets := getUserInput()

	//Call Validation Function
	isValidName, isValidEmail, isValidTicketNumber := UserValidation(userName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber{

		//Call for ticketBooking Function
		ticketBooking(userTickets, userName, email)

		wg.Add(1)
		go sendTicket(userTickets, userName, email)

		// Call getUserName Function
		userNames := getUserName()
		fmt.Printf("The names of the bookings are: %v\n", userNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our Conference is booked out. Come again Next Year.")
			//break
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
		//continue
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your tickets now!")
}

func getUserName() []string {
	userNames := []string{}
	for _, booking := range bookings {

		userNames = append(userNames, booking.userName)
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

	var userData = UserData {
		userName: userName,
		email: email,
		ticketNumber: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v for booking %v ticket with us! You will receive a confirmation email soon.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, userName string, email string) {

	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)

	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##############")

	wg.Done()
}