package main

import (
	"tableWaiter/db"
	"tableWaiter/table"
)

// Init all data deps
func setupDataDependacies() {
	var tableDB = &table.TableDB{}
	db.InitializeTable(tableDB)

}
