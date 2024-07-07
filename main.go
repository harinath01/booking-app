package main

import (
	"fmt"
	"strings"
	data "booking-app/data_classes"
	"booking-app/utils"
	"sync"
)

const conferenceName = "Go Conference"
const conferenceTicketsCount uint = 50

var remainingTicketsCount = conferenceTicketsCount
var bookings []data.Booking
var emailChannel = make(chan data.EmailRequest)
var wg sync.WaitGroup

func main() {
	greetUser()

	wg.Add(1)
	go utils.SendEmail(emailChannel, &wg)

	for {
		firstName, lastName, email, numOfTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, numOfTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(numOfTickets, firstName, lastName, email)
			printBookedUsersFirstName()
		} else {
			printInvalidInputErrorMsg(isValidEmail, isValidName, isValidTicketNumber, numOfTickets)
			continue
		}

		if remainingTicketsCount == 0 {
			fmt.Println("We're booked out, please come back next year.")
			close(emailChannel)
			wg.Wait()
			break
		}
	}
}



func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v tickets are still available\n", conferenceTicketsCount, remainingTicketsCount)
	fmt.Printf("Get your tickets here to attend\n")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numOfTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets, you want to book:")
	fmt.Scan(&numOfTickets)
	return firstName, lastName, email, numOfTickets
}

func validateUserInputs(firstName string, lastName string, email string, numOfTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := numOfTickets <= remainingTicketsCount
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(numOfTickets uint, firstName string, lastName string, email string) {
	remainingTicketsCount = remainingTicketsCount - numOfTickets
	user := data.User {
		FirstName: firstName,
		LastName: lastName,
		Email: email,
	}

	booking := data.Booking{
		BookedBy: user,
		NumberOfTickets: numOfTickets,
	}

	bookings = append(bookings, booking)
	fmt.Printf("Thank you, %v, for booking %v for %v, you will receive a confirmation at %v\n", user.GetFullName(), numOfTickets, conferenceName, email)
	emailChannel <- data.EmailRequest {
		To: user.Email,
		Subject: fmt.Sprintf("%v Ticket confirmation", conferenceName),
		Body: fmt.Sprintf("Thank you for booking %v tickets for %v, but please don't come"),
	}
}


func printBookedUsersFirstName() {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.BookedBy.GetFullName())
	}

	fmt.Printf("These guys also booked tickets for this conference: %v\n", firstNames)
}

func printInvalidInputErrorMsg(isValidEmail bool, isValidName bool, isValidTicketNumber bool, numOfTickets uint) {
	if !isValidEmail {
		fmt.Println("You've provided invalid email")
	}
	if !isValidName {
		fmt.Println("Please provide valid first name and last name, the one you provided was invalid")
	}
	if !isValidTicketNumber {
		fmt.Printf("We've only %v tickets left, so you can't book %v tickets\n", remainingTicketsCount, numOfTickets)
	}
}
