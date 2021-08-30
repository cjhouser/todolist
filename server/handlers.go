package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/CJHouser/tasklist/models"
	"github.com/gorilla/mux"
)

func (env *env) createTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)
	err := env.tasks.Insert(task)
	if err != nil {
		log.Printf("createTask: %v\n", err)
	}
}

func (env *env) returnAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := env.tasks.All()
	if err != nil {
		log.Printf("returnAllTasks: %v\n", err)
	}
	json.NewEncoder(w).Encode(tasks)
}

func (env *env) returnSingleTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	requestTask := models.Task{Id: key}
	responseTask, err := env.tasks.Single(requestTask)
	if err != nil {
		log.Printf("returnSingleTask %v\n", err)
	}
	json.NewEncoder(w).Encode(responseTask)
}

func (env *env) updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestTask models.Task
	json.Unmarshal(reqBody, &requestTask)
	requestTask.Id = key
	err := env.tasks.Update(requestTask)
	if err != nil {
		log.Printf("updateTask %v\n", err)
	}
}

func (env *env) deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	requestTask := models.Task{Id: key}
	err := env.tasks.Delete(requestTask)
	if err != nil {
		log.Printf("deleteTask %v\n", err)
	}
}
