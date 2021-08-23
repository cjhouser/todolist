package main

var dataStructure = taskArray{uid: 0, tasks: []task{}}

type task struct {
	id   string `json:"id"`
	name string `json:"name"`
}

func main() {
	handleRequests()
}
