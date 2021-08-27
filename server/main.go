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

type env struct {
	todoDB data.TodoDB
}

func main() {
	log.Println("Initializing Database")
	db := &data.TodoDB{}
	err := db.OpenDb(dbname, dbuser)
	if err != nil {
		log.Fatalf("unable to open database %v", err)
	}
	err = db.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalf("table not created %v", err)
	}
	env := &env{
		todoDB: *db,
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	// Middleware
	myRouter.Use(logging)
	// Routes
	myRouter.HandleFunc("/tasks", env.returnAllTodoItems)
	myRouter.HandleFunc("/task", env.createTodoItem).Methods("POST")
	myRouter.HandleFunc("/task/{id}", env.updateTodoItem).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", env.deleteTodoItem).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", env.returnSingleTodoItem)

	log.Println("Listening for requests")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
