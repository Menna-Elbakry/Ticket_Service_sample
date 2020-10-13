package database

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

var db = newDatabaseConnection()

//newDatabaseConnection connect to Database
func newDatabaseConnection() *pg.DB {
	opt := &pg.Options{
		User:     "postgres",
		Database: "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
	}
	db := pg.Connect(opt)
	if db == nil {
		log.Printf("Failed To Connect To Database")
		os.Exit(100)
	}
	//log.Printf("Connected Successfully To Database")

	return db
}

//CreateTables to add tables to db if not exist
func CreateTables() {
	CreateTicketTable()
	CreateUserTable()
}
