package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/CJHouser/tasklist/models"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
)

func (env *taskEnvironment) createTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)
	err := env.TaskInsert(task)
	if err != nil {
		log.Printf("createTask: %v\n", err)
	}
}

func (env *taskEnvironment) returnAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := env.TaskAll()
	if err != nil {
		log.Printf("returnAllTasks: %v\n", err)
	}
	json.NewEncoder(w).Encode(tasks)
}

func (env *taskEnvironment) returnSingleTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	requestTask := models.Task{Id: key}
	responseTask, err := env.TaskSingleCache(requestTask)
	if err != nil {
		log.Printf("returnSingleTask %v\n", err)
		// Get from DB on cache error and miss
		responseTask, err = env.TaskSingle(requestTask)
		if err != nil {
			log.Printf("returnSingleTask %v\n", err)
			// Internal Server Error
		}
		responseTaskBytes, err := json.Marshal(responseTask)
		if err != nil {
			log.Printf("returnSingleTask %v\n", err)
		}
		responseTaskItem := memcache.Item{
			Key:        fmt.Sprintf("%d_task", key),
			Value:      responseTaskBytes,
			Expiration: 5,
		}
		env.MC.Set(&responseTaskItem)
	}
	json.NewEncoder(w).Encode(responseTask)
}

func (env *taskEnvironment) updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestTask models.Task
	json.Unmarshal(reqBody, &requestTask)
	requestTask.Id = key
	err := env.TaskUpdate(requestTask)
	if err != nil {
		log.Printf("updateTask %v\n", err)
	}
}

func (env *taskEnvironment) deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	requestTask := models.Task{Id: key}
	err := env.TaskDelete(requestTask)
	if err != nil {
		log.Printf("deleteTask %v\n", err)
	}
}
