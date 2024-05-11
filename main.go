package main

import (
	"fmt"
	"log"
	"tableWaiter/party"
	"tableWaiter/restaurant"
)

func main() {
	var potentialParty party.Party

	setupDB() // setup necessary database

	restaurant.Greet()
	restaurant.DeclareHours()
	// restaurant.HandleUserClosedEntranceResponse(time.Now()) // handles responding to the user if not open/taking res
	err := potentialParty.AskForInfo()
	if err != nil {
		log.Fatal("We must gather all of your information to be seated.")
	}

	fmt.Println(potentialParty)
}
