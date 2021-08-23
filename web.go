package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
	log.Println("endpoint: homepage")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: create task")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var t task
	json.Unmarshal(reqBody, &t)
	dataStructure.create(t)
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: tasks")
	json.NewEncoder(w).Encode(dataStructure.read())
}

func returnSingleTask(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: single task")
	vars := mux.Vars(r)
	key := vars["id"]
	json.NewEncoder(w).Encode(dataStructure.readOne(key))
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: update task")
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var t task
	json.Unmarshal(reqBody, &t)
	dataStructure.update(key, t.Name)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: delete task")
	vars := mux.Vars(r)
	key := vars["id"]
	dataStructure.delete(key)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Static
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", returnAllTasks)
	myRouter.HandleFunc("/task", createTask).Methods("POST")
	// Parameterized
	myRouter.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", returnSingleTask)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
