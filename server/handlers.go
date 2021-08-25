package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/CJHouser/tasklist/data"
	"github.com/gorilla/mux"
)

var db *data.TodoDB

func createTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: create todoItem")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	json.Unmarshal(reqBody, &item)
	err := db.InsertTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: create todoItem: %v\n", err)
	}
}

func returnAllTodoItems(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: todoItems")
	items, err := db.SelectAllTodoItems()
	if err != nil {
		log.Printf("FAIL: create todoItem: %v\n", err)
	}
	json.NewEncoder(w).Encode(items)
}

/*
func returnSingleTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: single todoItem")
	vars := mux.Vars(r)
	key := vars["id"]
	db.
	json.NewEncoder(w).Encode(dataStructure.readOne(key))
}
*/

func updateTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: update todoItem")
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	json.Unmarshal(reqBody, &item)
	item.Id, _ = strconv.Atoi(key)
	err := db.UpdateTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: update todoItem: %v\n", err)
	}
}

func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: delete todoItem")
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	item.Id, _ = strconv.Atoi(key)
	json.Unmarshal(reqBody, &item)
	err := db.DeleteTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: create todoItem: %v\n", err)
	}
}

func handleRequests(dbp *data.TodoDB) {
	db = dbp
	myRouter := mux.NewRouter().StrictSlash(true)
	// Static
	myRouter.HandleFunc("/tasks", returnAllTodoItems)
	myRouter.HandleFunc("/task", createTodoItem).Methods("POST")
	// Parameterized
	myRouter.HandleFunc("/task/{id}", updateTodoItem).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTodoItem).Methods("DELETE")
	//myRouter.HandleFunc("/task/{id}", returnSingleTodoItem)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
