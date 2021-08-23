package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
	log.Println("endpoint: homepage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: tasks")
	json.NewEncoder(w).Encode(dataStructure.read())
}

func returnSingleTask(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: task")
	vars := mux.Vars(r)
	key := vars["id"]
	json.NewEncoder(w).Encode(dataStructure.readOne(key))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", returnAllTasks)
	myRouter.HandleFunc("/task/{id}", returnSingleTask)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
