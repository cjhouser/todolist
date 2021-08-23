package main

import (
	"bufio"
	"fmt"
	"os"
)

var dataStructure = taskArray{uid: 0, tasks: []task{}}

type task struct {
	id   string `json:"id"`
	name string `json:"name"`
}

type ops interface {
	create(string)
	read() []string
	update(string, string)
	delete(string)
}

func run(o ops) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("---Tasks---")
		for i, v := range o.read() {
			fmt.Printf("%d. %s\n", i+1, v)
		}
		fmt.Print("Command: ")
		input, _, _ := reader.ReadLine()
		switch string(input) {
		case "create":
			fmt.Print("Create task: ")
			input, _, _ := reader.ReadLine()
			o.create(string(input))
		case "update":
			fmt.Print("Update id: ")
			id, _, _ := reader.ReadLine()
			fmt.Print("Update: ")
			input, _, _ := reader.ReadLine()
			o.update(string(id), string(input))
		case "delete":
			fmt.Print("Delete id: ")
			id, _, _ := reader.ReadLine()
			o.delete(string(id))
		}
	}
}

func main() {
	//run(&dataStructure)
	handleRequests()
}
