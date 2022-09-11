package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v all still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of tickets:")
	fmt.Scan(&userTickets)
	// & = pointer to the memory address
	// it'ss not possible to assign user's value without the pointer

	remainingTickets -= userTickets

	userTickets = 2

	fmt.Printf("Thank you %v %v for booking %v tikcets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

}
