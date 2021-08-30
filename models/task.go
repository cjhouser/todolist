package models

import (
	"database/sql"
)

const (
	sqlInsert = `INSERT INTO tasks (title) VALUES(?)`
	sqlAll    = `SELECT id, title FROM tasks`
	sqlSingle = `SELECT id, title FROM tasks WHERE id=?`
	sqlUpdate = `UPDATE tasks SET title=? WHERE id=?`
	sqlDelete = `DELETE FROM tasks WHERE id=?`
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type TaskModel struct {
	DB *sql.DB
}

func (m TaskModel) Insert(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlInsert,
		requestTask.Title)
	return err
}

func (m TaskModel) All() (responseTasks []Task, err error) {
	rows, err := m.DB.Query(sqlAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Title)
		if err != nil {
			return nil, err
		}
		responseTasks = append(responseTasks, task)
	}
	return responseTasks, nil
}

func (m TaskModel) Single(requestTask Task) (responseTask *Task, err error) {
	row := m.DB.QueryRow(
		sqlSingle,
		requestTask.Id)
	responseTask = &Task{}
	err = row.Scan(&responseTask.Id, &responseTask.Title)
	if err != nil {
		return nil, err
	}
	return responseTask, nil
}

func (m TaskModel) Update(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlUpdate,
		requestTask.Title,
		requestTask.Id)
	return err
}

func (m TaskModel) Delete(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlDelete,
		requestTask.Id)
	return err
}
