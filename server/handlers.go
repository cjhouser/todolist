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

func (env *env) createTodoItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	json.Unmarshal(reqBody, &item)
	err := env.todoDB.InsertTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: create todoItem: %v\n", err)
	}
}

func (env *env) returnAllTodoItems(w http.ResponseWriter, r *http.Request) {
	items, err := env.todoDB.SelectAllTodoItems()
	if err != nil {
		log.Printf("FAIL: return all todoItems: %v\n", err)
	}
	json.NewEncoder(w).Encode(items)
}

func (env *env) returnSingleTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	item := data.DBTodoItem{&models.TodoItem{
		Id:    key,
		Title: "",
	}}
	todoItem, err := env.todoDB.SelectSingleTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: return single todoItem: %v\n", err)
	}
	json.NewEncoder(w).Encode(todoItem)
}

func (env *env) updateTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item data.DBTodoItem
	json.Unmarshal(reqBody, &item)
	item.Id = key
	err := env.todoDB.UpdateTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: update todoItem: %v\n", err)
	}
}

func (env *env) deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	item := data.DBTodoItem{&models.TodoItem{
		Id:    key,
		Title: "",
	}}
	err := env.todoDB.DeleteTodoItem(&item)
	if err != nil {
		log.Printf("FAIL: delete todoItem: %v\n", err)
	}
}
