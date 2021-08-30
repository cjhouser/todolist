package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CJHouser/tasklist/models"
	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbname              = "tasklistapp"
	dbuser              = "root"
	dbpass              = "pass"
	sqlCreateTasksTable = `CREATE TABLE IF NOT EXISTS tasks (id BIGINT PRIMARY KEY AUTO_INCREMENT, title TEXT NOT NULL)`
)

type env struct {
	tasks models.TaskModel
}

func main() {
	log.Println("Connecting to database")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	if err != nil {
		log.Fatalf("unable to open database %v", err)
	}
	if _, err := db.Exec(sqlCreateTasksTable); err != nil {
		log.Fatalf("unable to create table %v", err)
	}
	log.Println("Connecting to cache")
	mc := memcache.New("localhost:11211")
	// Dependency injection
	env := &env{
		tasks: models.TaskModel{DB: db, MC: mc},
	}
	// Routes
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", env.returnAllTasks)
	myRouter.HandleFunc("/task", env.createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", env.updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", env.deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", env.returnSingleTask)
	// Middleware
	myRouter.Use(logging)
	log.Println("Listening for requests")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
