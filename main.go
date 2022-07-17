package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTicket, "are still available")
	fmt.Println("Get your ticket here to attend")
}
