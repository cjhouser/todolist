package main

import (
	"container/list"
	"fmt"
)

type ops interface {
	create()
	read() []string
	update()
	delete()
}

type taskList struct {
	tasks list.List
}

func (tl taskList) create(newTask string) {
	tl.tasks.PushBack(newTask)
}

func (tl taskList) read() (result []string) {
	index := 1
	for task := tl.tasks.Front(); task != nil; task.Next() {
		fmt.Printf("%d. %s", index, task.Value)
		index++
	}
	return
}

func (tl taskList) update(updateIndex int, update string) {
	task := tl.tasks.Front()
	for index := 1; index < updateIndex; index++ {
		task.Next()
	}
	task.Value = update
}

func (tl taskList) delete(deleteIndex int) {
	task := tl.tasks.Front()
	for index := 1; index < deleteIndex; index++ {
		task.Next()
	}
	tl.tasks.Remove(task)
}
