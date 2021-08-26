package main

import (
	"log"

	"github.com/CJHouser/tasklist/data"
)

const (
	dbname = "todoapp"
	dbuser = "root"
)

// note error handling is omitted for brevity - handle err returns!
func main() {
	log.Println("Initializing Database")
	db := &data.TodoDB{}
	db.OpenDb(dbname, dbuser)
	err := db.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalf("table not created %v", err)
	}
	/*
		// fetch DB items as a typed slice of DBTodoItem
		items, _ := db.SelectAllTodoItems()
		log.Println(len(*items), "items in the database")

		// declare a new DBTodoItem (with embedded TodoItem)
		item := &data.DBTodoItem{&models.TodoItem{}}
		item.Title = "test"

		// insert into the DB
		db.InsertTodoItem(item)
		log.Printf("%s (done: %t) inserted into the database", item.Title, item.Done)

		// demonstrate successful insert
		items, _ = db.SelectAllTodoItems()
		log.Println(len(*items), "items in the database")
		item = (*items)[0]
		log.Printf("%s (done: %t) found in database", item.Title, item.Done)

		// update in DB
		db.UpdateTodoItem(item)
		log.Printf("%s (done: %t) updated in the database", item.Title, item.Done)

		// demonstrate successful update
		items, _ = db.SelectAllTodoItems()
		log.Println(len(*items), "items in the database")
		item = (*items)[0]
		log.Printf("%s (done: %t) found in database", item.Title, item.Done)

		// delete item
		db.DeleteTodoItem(item)

		// demonstrate successful deletion 	items, _ = db.SelectAllTodoItems()
		log.Println(len(*items), "items in the database")
	*/
	log.Println("Listening Request")
	handleRequests(db)
}
