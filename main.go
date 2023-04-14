package main

import "fmt"

func main() {
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"

	const conferenceTickets = 42
	var remainingTickets uint = 42

	fmt.Println("Hello, gopher!")
	fmt.Printf("Welcome to this year's %v\n", conferenceName)
	fmt.Println("Get your tickets here...")
	fmt.Printf("Total Tickets: %v\n", conferenceTickets )
	fmt.Printf("Remaining Tickets: %v\n", remainingTickets )

	// var bookings = [42]string{"James", "John", "Mary"}
	// var bookings [42]string //42 => array size
	var bookings []string //slice
	// bookings[0] = "James"
	// bookings[1] = "John"
	// bookings[2] = "Mary"


	var userName string
	var userEmail string
	var userTickets uint

	// user input: (&) pointer => variable that points to the memory address of another variable
	// fmt.Println(remainingTickets) //actual value
	// fmt.Println(&remainingTickets) //memory location
	fmt.Print("Your name: ")
	fmt.Scan(&userName)
	fmt.Print("Your email address: ")
	fmt.Scan(&userEmail)
	fmt.Print("Number of tickets bought: ")
	fmt.Scan(&userTickets)

	// next, reduce ticket number
	remainingTickets -= userTickets

	// bookings
	// bookings[0] = userName + " => " + userEmail //array
	bookings = append(bookings, userName + " => " + userEmail) //slice: flexible and dynamic in nature

	fmt.Printf("%v has booked %v slots for the %v event\n", userName, userTickets, conferenceName)
	fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
	fmt.Printf("First booking value: %v\n", bookings[0])
	fmt.Printf("Tickets booked so far: %v\n", bookings)
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))
	fmt.Printf("Slice Type: %T\n", bookings)
	fmt.Printf("Slice Length: %v\n", len(bookings))

}