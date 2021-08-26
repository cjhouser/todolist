package main

import (
	"log"

	"github.com/CJHouser/tasklist/data"
)

const (
	dbname = "todoapp"
	dbuser = "root"
)

func main() {
	log.Println("Initializing Database")
	db := &data.TodoDB{}
	db.OpenDb(dbname, dbuser)
	err := db.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalf("table not created %v", err)
	}
	log.Println("Listening for requests")
	handleRequests(db)
}
