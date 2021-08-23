package main

type taskArray struct {
	tasks []task
}

func (ta *taskArray) create(newTask string) {
	ta.tasks = append(ta.tasks, task{name: newTask})
}

func (ta taskArray) read() (result []string) {
	for _, t := range ta.tasks {
		result = append(result, t.name)
	}
	return
}

func (ta taskArray) update(updateIndex int, update string) {
	ta.tasks[updateIndex-1].name = update
}

func (ta *taskArray) delete(deleteIndex int) {
	if deleteIndex == len(ta.tasks)-1 {
		ta.tasks = ta.tasks[:len(ta.tasks)-1]
	} else {
		ta.tasks = append(ta.tasks[:deleteIndex-1], ta.tasks[deleteIndex:]...)
	}
}
