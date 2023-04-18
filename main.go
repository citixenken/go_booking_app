package main

import (
	"fmt"

	"github.com/citixenken/go_booking_app/helper"
)

// package level variables
const conferenceTickets = 42
var conferenceName = "Go Conference"
// conferenceName := "Go Conference"
var remainingTickets uint = 42
// var bookings = make([]map[string]string, 0) //empty slice of maps
var bookings = make([]UserData, 0) //empty slice of struct

type UserData struct {
	firstName string
	userEmail string
	noOfTickets uint
}

func main() {
	// fmt.Println("Hello, gopher!")
	// fmt.Printf("Welcome to this year's %v\n", conferenceName)
	// fmt.Println("Get your tickets here...")
	// fmt.Printf("Total Tickets: %v\n", conferenceTickets )
	// fmt.Printf("Remaining Tickets: %v\n", remainingTickets )
	onboardUsers()

	// var bookings = [42]string{"James", "John", "Mary"}
	// var bookings [42]string //42 => array size
	// bookings[0] = "James"
	// bookings[1] = "John"
	// bookings[2] = "Mary"

	for {

		// user input: (&) pointer => variable that points to the memory address of another variable
		// fmt.Println(remainingTickets) //actual value
		// fmt.Println(&remainingTickets) //memory location

		//get user inputs
		firstName, lastName, userEmail, userTickets := getUserInputs()

		// validate user input
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)


		if isValidName && isValidEmail &&isValidTicket {

			// book ticket
			bookTicket(userTickets, firstName, userEmail)

			fmt.Printf("%v has booked %v slots for the %v event\n", firstName, userTickets, conferenceName)
			fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
			fmt.Printf("Tickets remaining: %v\n", remainingTickets)
			fmt.Printf("First booking value: %v\n", bookings[0])
			fmt.Printf("Tickets booked so far: %v\n", bookings)
			// fmt.Printf("Array Type: %T\n", bookings)
			// fmt.Printf("Array Length: %v\n", len(bookings))
			fmt.Printf("Slice Type: %T\n", bookings)
			fmt.Printf("Slice Length: %v\n", len(bookings))

			if remainingTickets == 0 {
				fmt.Println("Go Conference sold out! Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address missing @ sign!")
			}
			if !isValidTicket {
				fmt.Println("Check your ticket input value!")
			}

			// fmt.Println("Invalid data input!")
			fmt.Printf("We only have %v tickets remaining. Try again!\n", remainingTickets)
		}
	}
}

func onboardUsers() {
	fmt.Println("Hello, gopher!")
	fmt.Printf("Welcome to this year's %v\n", conferenceName)
	fmt.Println("Get your tickets here...")
	fmt.Printf("Total Tickets: %v\n", conferenceTickets )
	fmt.Printf("Remaining Tickets: %v\n", remainingTickets )
}

func getUserInputs() (string, string, string, uint) {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		fmt.Print("Your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Your email address: ")
		fmt.Scan(&userEmail)
		fmt.Print("Number of tickets bought: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, userEmail string)  {
	// next, reduce ticket number
	remainingTickets -= userTickets

	// bookings
	// create a user map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["userEmail"] = userEmail
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		userEmail: userEmail,
		noOfTickets: userTickets,
	}

	// bookings[0] = userName + " => " + userEmail //array
	// bookings = append(bookings, firstName + " => " + userEmail) //slice: flexible and dynamic in nature
	bookings = append(bookings, userData)
}