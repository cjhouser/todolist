package main

import "fmt"

type taskArray struct {
	uid   int
	tasks []task
}

func (ta *taskArray) create(newTask string) {
	ta.tasks = append(ta.tasks, task{Id: fmt.Sprint(ta.uid), Name: newTask})
	ta.uid++
}

func (ta taskArray) readOne(id string) (result string) {
	for _, t := range ta.tasks {
		if t.Id == id {
			result = t.Name
			break
		}
	}
	return result
}

func (ta taskArray) read() (result []string) {
	for _, t := range ta.tasks {
		result = append(result, t.Name)
	}
	return
}

func (ta taskArray) update(id string, updatedName string) {
	for _, t := range ta.tasks {
		if t.Id == id {
			t.Name = updatedName
		}
	}
}

func (ta *taskArray) delete(id string) {
	for i, t := range ta.tasks {
		if t.Id == id {
			if i == len(ta.tasks)-1 {
				ta.tasks = ta.tasks[:len(ta.tasks)-1]
			} else {
				ta.tasks = append(ta.tasks[:i], ta.tasks[i+1:]...)
			}
		}
	}
}
