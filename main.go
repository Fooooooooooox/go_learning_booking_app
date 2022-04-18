package main

import (
	"fmt"

	"github.com/Fooooooooooox/foooox"

	"rsc.io/quote"
)

func main() {
	var conferenceName = "go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T", conferenceName)

	fmt.Println(foooox.Bestfoooox())
	fmt.Println(quote.Go())

	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %vtickets\n", conferenceTickets)
	fmt.Printf("there are %v tickets left\n", remainingTickets)

	var userName string
	var userTickets uint
	// ask user for their names

	// scan is a pointer so it looks for the memory location(pointer is a concept in c, but many popular languages like java do not have pointers or pointers are not exposed to users)
	fmt.Println("Get your name here to attend")
	fmt.Scan(&userName)

	// &is used to print memory location of a variable
	// like:
	fmt.Println(&remainingTickets)

	fmt.Println("input the number of tickets you want:")
	fmt.Scan(&userTickets)

	// caculation between numbers ==> they must have same type(int and uint won't work)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

	fmt.Printf("there are %v tickets left\n", remainingTickets)
}
