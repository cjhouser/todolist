package main

import "fmt"

type taskArray struct {
	uid   int
	tasks []task
}

func (ta *taskArray) create(newTask task) {
	newTask.Id = fmt.Sprint(ta.uid)
	ta.tasks = append(ta.tasks, newTask)
	ta.uid++
}

func (ta taskArray) readOne(id string) (result task) {
	for _, t := range ta.tasks {
		if t.Id == id {
			result = t
			break
		}
	}
	return result
}

func (ta taskArray) read() (result []task) {
	return append(result, ta.tasks...)
}

func (ta taskArray) update(id string, updatedName string) {
	for i := 0; i < len(ta.tasks); i++ {
		if ta.tasks[i].Id == id {
			ta.tasks[i].Name = updatedName
		}
	}
}

func (ta *taskArray) delete(id string) {
	for i, t := range ta.tasks {
		if t.Id == id {
			ta.tasks = append(ta.tasks[:i], ta.tasks[i+1:]...)
		}
	}
}
