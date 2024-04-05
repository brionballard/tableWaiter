package main

import (
	"fmt"
)

func main() {
	setupDataDependacies()

	var partyName string
	var partySize int

	fmt.Println("Hi, welcome to TBL!")
	fmt.Println("Can I get the name for your party?")
	fmt.Scanln(&partyName)

	fmt.Println("And how many are in your party?")
	fmt.Scanln(&partySize)
}
