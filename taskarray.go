package main

type taskArray struct {
	tasks []string
}

func (ta *taskArray) create(newTask string) {
	ta.tasks = append(ta.tasks, newTask)
}

func (ta taskArray) read() []string {
	return ta.tasks
}

func (ta taskArray) update(updateIndex int, update string) {
	ta.tasks[updateIndex-1] = update
}

func (ta *taskArray) delete(deleteIndex int) {
	if deleteIndex == len(ta.tasks)-1 {
		ta.tasks = ta.tasks[:len(ta.tasks)-1]
	} else {
		ta.tasks = append(ta.tasks[:deleteIndex-1], ta.tasks[deleteIndex:]...)
	}
}
