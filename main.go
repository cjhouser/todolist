package main

var dataStructure = taskArray{uid: 0, tasks: []task{}}

type task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	handleRequests()
}
