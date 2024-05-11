package main

import (
	"tableWaiter/db"
	"tableWaiter/table"
)

// Init all data deps
func setupDB() {
	var tableDB = &table.TableDB{}
	db.InitializeTable(tableDB)
}
