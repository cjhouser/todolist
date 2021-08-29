package models

import "database/sql"

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type TodoItemModel struct {
	DB *sql.DB
}
