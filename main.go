package main

import (
	"log"
	"tableWaiter/party"
	"tableWaiter/restaurant"
)

func main() {
	var potentialParty party.Party

	setupDB() // setup necessary database

	restaurant.Greet()
	restaurant.DeclareHours()
	//restaurant.HandleCurrentlyClosedResponse(time.Now())

	err := potentialParty.AskForInfo()
	if err != nil {
		log.Fatal("We must gather all of your information to be seated.")
	}

	potentialParty.Seat()
}
