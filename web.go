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
	json.NewEncoder(w).Encode(dataStructure.tasks)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", returnAllTasks)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
