package main

import (
	"container/list"
)

type ops interface {
	create(string)
	read() []interface{}
	update(int, string)
	delete(int)
}

type taskList struct {
	tasks *list.List
}

func (tl taskList) create(newTask string) {
	tl.tasks.PushBack(newTask)
}

func (tl taskList) read() (result []interface{}) {
	for task := tl.tasks.Front(); task != nil; task = task.Next() {
		result = append(result, task.Value)
	}
	return
}

func (tl taskList) update(updateIndex int, update string) {
	task := tl.tasks.Front()
	for index := 1; index < updateIndex; index++ {
		task = task.Next()
	}
	task.Value = update
}

func (tl taskList) delete(deleteIndex int) {
	task := tl.tasks.Front()
	for index := 1; index < deleteIndex; index++ {
		task = task.Next()
	}
	tl.tasks.Remove(task)
}
