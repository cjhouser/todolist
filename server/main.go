package main

import (
	"log"
	"net/http"

	"github.com/CJHouser/tasklist/data"
	"github.com/gorilla/mux"
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

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", returnAllTodoItems)
	myRouter.HandleFunc("/task", createTodoItem).Methods("POST")
	myRouter.HandleFunc("/task/{id}", updateTodoItem).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTodoItem).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", returnSingleTodoItem)

	log.Println("Listening for requests")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
