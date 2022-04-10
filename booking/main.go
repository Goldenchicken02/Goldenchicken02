package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("%v tickets in total\n", conferenceTickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)

	bookings := []string{}
	for {
		var userFname string
		var userLname string
		var userEmail string
		var userTickets uint

		fmt.Println("Registration System")
		fmt.Println("Firstname: ")
		fmt.Scan(&userFname)

		fmt.Println("Lastname: ")
		fmt.Scan(&userLname)

		fmt.Println("Email: ")
		fmt.Scan(&userEmail)

		fmt.Println("Ticket Amount: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userFname+" "+userLname)

		fmt.Printf("User %v %v booked %v tickets. Get your confirmation email at %v.\n", userFname, userLname, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		fmt.Printf("The bookings list is %v\n", bookings)
	}
}
