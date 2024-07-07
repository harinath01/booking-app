package main

import (
	"fmt"
	"strings"
)

const conferenceName = "Go Conference"
const conferenceTicketsCount uint = 50

var remainingTicketsCount = conferenceTicketsCount
var bookings []Booking

type Booking struct {
	bookedBy User
	numberOfTickets uint
}

type User struct {
	firstName string
	lastName string
	email string
}

func (user User) getFullName() string {
	return user.firstName + " " + user.lastName
}

func main() {
	greetUser()

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
	user := User {
		firstName: firstName,
		lastName: lastName,
		email: email,
	}

	booking := Booking{
		bookedBy: user,
		numberOfTickets: numOfTickets,
	}

	bookings = append(bookings, booking)
	fmt.Printf("Thank you, %v, for booking %v for %v, you will receive a confirmation at %v\n", user.getFullName(), numOfTickets, conferenceName, email)
}


func printBookedUsersFirstName() {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.bookedBy.getFullName())
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
