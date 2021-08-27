package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/CJHouser/tasklist/data"
	"github.com/CJHouser/tasklist/models"
	"github.com/gorilla/mux"
)

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
		log.Printf("FAIL: return all todoItems: %v\n", err)
	}
	json.NewEncoder(w).Encode(items)
}

func returnSingleTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: single todoItem")
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	item := data.DBTodoItem{&models.TodoItem{
		Id:    key,
		Title: "",
	}}
	todoItem, err := db.SelectSingleTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: return single todoItem: %v\n", err)
	}
	json.NewEncoder(w).Encode(todoItem)
}

func updateTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: update todoItem")
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	json.Unmarshal(reqBody, &item)
	item.Id = key
	err := db.UpdateTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: update todoItem: %v\n", err)
	}
}

func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: delete todoItem")
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	item := data.DBTodoItem{&models.TodoItem{
		Id:    key,
		Title: "",
	}}
	err := db.DeleteTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: delete todoItem: %v\n", err)
	}
}
