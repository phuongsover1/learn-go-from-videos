package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get you tickets here to attend")

	//var bookings = [50]string{"Nana", "Nicole", "Peter", "Khang"}
	bookings := []string{}

	bookings = append(bookings, "Nana")
	bookings = append(bookings, "Nicole")

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
	// bookings[0] = "Nana"
	// bookings[1] = "Nicole"

	// fmt.Println("first 3rd array value: ")
	// println(bookings[0])
	// println(bookings[1])
	// println(bookings[2])

	// fmt.Println("first 3rd array address value: ")
	// println(&bookings)
	// println(&bookings[0])
	// println(&bookings[1])
	// println(&bookings[2])

	// fmt.Println("The array size is", len(bookings))
	// fmt.Printf("The type of array: %T\n", bookings)

	// fmt.Println("TEST")
	// var test []int
	// test[0] = 1
	// test[1] = 2

	// fmt.Println(test[0])
	// fmt.Println(test[1]) // error

	for remainingTickets > 0 {
		var userName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Printf("Enter your firstname: ")
		fmt.Scan(&userName)

		fmt.Printf("Enter you last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter you email address: ")
		fmt.Scan(&email)

		fmt.Print("How many ticket do you want to book: ")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("The remaining tickets are %v. So you cannot book %v tickets\n", remainingTickets, userTickets)
			continue
		}
		fmt.Printf("User %v %v with email %v booked %v tickets.\n", userName, lastName, email, userTickets)
		fmt.Printf("You will receive a confirmation email at %v \n", email)

		bookings = append(bookings, userName+" "+lastName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)
		fmt.Printf("These are all our bookings: %v\n", bookings)
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
		}
	}

}
