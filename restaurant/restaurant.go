package restaurant

import (
	"fmt"
	"os"
	"time"
)

const (
	Opens                = "4:00 PM"
	OpensAsInt           = 4
	ReservationsEnd      = "10:00 PM"
	ReservationsEndAsInt = 10
	Closes               = "12:00 PM"
	ClosesAsInt          = 12
	TimeLayout           = "3:04 PM" // Time layout for parsing time strings
	Open                 = 2
	NotTakingRes         = 1
	Closed               = 0

	TableCount         = 12
	MaxSeatingPerTable = 12
)

var TableSections = []string{"A", "B", "C", "D"}

// Greet welcomes the user
func Greet() {
	fmt.Println("Hi, welcome to TBL!")
}

// DeclareHours declares hours of business to the user
func DeclareHours() {
	fmt.Println("We are open from 4 p.m. to 12 p.m.")
	fmt.Println("We stop taking reservations at 10 p.m.")
}

// IsRestaurantOpen checks to see if the user has come at a valid time
func IsRestaurantOpen(curr time.Time) int {
	open, _ := time.Parse(TimeLayout, Opens)
	noRes, _ := time.Parse(TimeLayout, ReservationsEnd)
	closed, _ := time.Parse(TimeLayout, ReservationsEnd)

	if curr.After(noRes) {
		return NotTakingRes
	} else if curr.Before(open) || curr.After(closed) {
		return Closed
	} else {
		return Open
	}
}

// HandleCurrentlyClosedResponse handles response to user if not open/taking users
func HandleCurrentlyClosedResponse(customerEntered time.Time) {
	restaurantState := IsRestaurantOpen(customerEntered)

	if restaurantState < 2 {
		if restaurantState == 1 {
			fmt.Println("Sorry, we are not currently taking reservations. Come back tomorrow and try to get a table!")
		} else {
			fmt.Println("Sorry, we are not open!")
		}
		os.Exit(0)
	}
}
