package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	sqlTaskInsert = `INSERT INTO tasks (title) VALUES(?)`
	sqlTaskAll    = `SELECT id, title FROM tasks`
	sqlTaskSingle = `SELECT id, title FROM tasks WHERE id=?`
	sqlTaskUpdate = `UPDATE tasks SET title=? WHERE id=?`
	sqlTaskDelete = `DELETE FROM tasks WHERE id=?`
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type TaskModel struct {
	DB *sql.DB
	MC *memcache.Client
}

func (m TaskModel) TaskInsert(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlTaskInsert,
		requestTask.Title)
	return err
}

func (m TaskModel) TaskAll() (responseTasks []Task, err error) {
	rows, err := m.DB.Query(sqlTaskAll)
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

func (m TaskModel) TaskSingle(requestTask Task) (responseTask *Task, err error) {
	row := m.DB.QueryRow(
		sqlTaskSingle,
		requestTask.Id)
	responseTask = &Task{}
	err = row.Scan(&responseTask.Id, &responseTask.Title)
	if err != nil {
		return nil, err
	}
	return responseTask, nil
}

func (m TaskModel) TaskUpdate(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlTaskUpdate,
		requestTask.Title,
		requestTask.Id)
	return err
}

func (m TaskModel) TaskDelete(requestTask Task) error {
	_, err := m.DB.Exec(
		sqlTaskDelete,
		requestTask.Id)
	return err
}

func (m TaskModel) TaskSingleCache(requestTask Task) (responseTask *Task, err error) {
	fetchedItem, err := m.MC.Get(fmt.Sprintf("%d_task", requestTask.Id))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fetchedItem.Value, &responseTask)
	if err != nil {
		return nil, err
	}
	return responseTask, nil
}
