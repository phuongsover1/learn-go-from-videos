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

	//var bookings = [50]string{"Nana", "Nicole", "Peter", "Khang"}
	var bookings [50]string
	bookings[0] = "Nana"
	bookings[1] = "Nicole"

	fmt.Println("first 3rd array value: ")
	println(bookings[0])
	println(bookings[1])
	println(bookings[2])	

	fmt.Println("first 3rd array address value: ")
	println(&bookings)
	println(&bookings[0])
	println(&bookings[1])
	println(&bookings[2])

	fmt.Println("TEST")
	var test []int
	test[0] = 1
	test[1] = 2

	fmt.Println(test[0])
	fmt.Println(test[1]) // error

	var userName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Printf("Enter your firstname: ")
	fmt.Scan(&userName)

	fmt.Printf("Enter you last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email address: ")
	fmt.Scan(&email)
	
	userTickets = 2
	fmt.Printf("User %v %v with email %v booked %v tickets.\n", userName,lastName, email, userTickets)


}
