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
	dbname                 = "tasklistapp"
	dbuser                 = "root"
	dbpass                 = "pass"
	sqlCreateTasksTable    = `CREATE TABLE IF NOT EXISTS tasks (id BIGINT PRIMARY KEY AUTO_INCREMENT, title TEXT NOT NULL)`
	sqlCreateAccountsTable = `CREATE TABLE IF NOT EXISTS accounts (id BIGINT PRIMARY KEY AUTO_INCREMENT, username TEXT NOT NULL, password VARCHAR(60) NOT NULL, salt TEXT NOT NULL)`
)

type accountEnvironment struct {
	accounts models.AccountModel
}

type taskEnvironment struct {
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
	if _, err := db.Exec(sqlCreateAccountsTable); err != nil {
		log.Fatalf("unable to create account %v", err)
	}
	log.Println("Connecting to cache")
	mc := memcache.New("localhost:11211")
	// Dependency injection
	accountEnv := &accountEnvironment{
		accounts: models.AccountModel{DB: db},
	}
	taskEnv := &taskEnvironment{
		tasks: models.TaskModel{DB: db, MC: mc},
	}
	// Routes
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users/new", accountEnv.signUp).Methods("POST")
	myRouter.HandleFunc("/tasks", taskEnv.returnAllTasks)
	myRouter.HandleFunc("/task", taskEnv.createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", taskEnv.updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", taskEnv.deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", taskEnv.returnSingleTask)

	// Middleware
	myRouter.Use(logging)
	log.Println("Listening for requests")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
