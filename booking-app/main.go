package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets,"tickets and", remainingTickets, "are still available")
	fmt.Println("Get you tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
