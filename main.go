package main

import (
	"bufio"
	"fmt"
	"os"
)

var dataStructure = taskArray{tasks: []task{}}

type task struct {
	name string `json:"name"`
}

type ops interface {
	create(string)
	read() []string
	update(int, string)
	delete(int)
}

func run(o ops) {
	var index int
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
			fmt.Print("Update index: ")
			fmt.Scan(&index)
			fmt.Print("Update: ")
			input, _, _ := reader.ReadLine()
			o.update(index, string(input))
		case "delete":
			fmt.Print("Delete index: ")
			fmt.Scan(&index)
			o.delete(index)
		}
	}
}

func main() {
	//run(&dataStructure)
	handleRequests()
}
