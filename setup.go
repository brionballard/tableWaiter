package main

import (
	"tableWaiter/db"
	"tableWaiter/table"
)

// setupDB all necessary databases and data for the program to run
func setupDB() {
	var tableDb table.TableDb
	db.InitializeTable(&tableDb)
}
