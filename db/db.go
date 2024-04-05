package db

type DB interface {
	Init()
	TableName() string
}

func InitializeTable(d DB) {
	d.Init()
}

func GetTableName(d DB) string {
	return d.TableName()
}
