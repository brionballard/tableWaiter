package main

import (
	"fmt"
	"tableWaiter/restaurant"
	"time"
)

func main() {
	setupDB()

	var partyName string
	var partySize int

	customerEntered := time.Now()

	restaurant.Greet()
	restaurant.DeclareHours()
	restaurant.HandleUserClosedEntranceResponse(customerEntered) // handles responding to the user if not open/taking res

	fmt.Println("Can I get the name for your party?")
	fmt.Scanln(&partyName)

	fmt.Println("And how many are in your party?")
	fmt.Scanln(&partySize)

}
