package main

import (
	"fmt"

	"github.com/Fooooooooooox/go_learning_booking_app/foooox"
)

func main() {
	// vertex
	v := foooox.Vertex{3, 4}
	fmt.Println(v.Abs())

	// var conferenceName = "go conference"
	// there are 2 ways to define a variable: var .. = or :==
	conferenceName := "go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T\n", conferenceName)

	fmt.Println(foooox.Bestfoooox())

	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %vtickets\n", conferenceTickets)

	// the number defines how many elements the array can hold
	// what about dynamic size?
	// automatically expand ==> slice
	// string is the datatype of element in that array
	// stuff in {} is the element(you can keep it null)
	// var bookings = [50]string{}
	// bookings[0] = "me"

	// slice
	var bookings []string

	for {
		var userName string
		var userTickets uint
		// ask user for their names

		// validator

		isValidname := len(userName) >= 2

		fmt.Println(isValidname)

		// scan is a pointer so it looks for the memory location(pointer is a concept in c, but many popular languages like java do not have pointers or pointers are not exposed to users)
		fmt.Println("Get your name here to attend")
		fmt.Scan(&userName)
		bookings = append(bookings, userName)
		fmt.Println(bookings)

		// &is used to print memory location of a variable
		// like:
		fmt.Println(&remainingTickets)

		fmt.Println("input the number of tickets you want:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			//  caculation between numbers ==> they must have same type(int and uint won't work)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

			fmt.Printf("there are %v tickets left\n", remainingTickets)

			for _, booking := range bookings {
				print(booking)
			}
		} else {
			fmt.Printf("we only have %v tickets left\n", remainingTickets)
			continue
		}

		if remainingTickets == 0 {
			//end program
			fmt.Println("our conference is booked out...")
			break
		}

	}
}
