package db

type DB interface {
	Init()
}

// Initialize specific table
func InitializeTable(d DB) {
	d.Init()
}
