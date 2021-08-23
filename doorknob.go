package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var tasks []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("What do?: ")
		text, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		switch string(text) {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Printf("exit\nhelp\ncreate\nread\nupdate\ndelete\n")
		case "create":
			fmt.Print("New task?: ")
			newTask, _, err := reader.ReadLine()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			tasks = append(tasks, string(newTask))
			for index, task := range tasks {
				fmt.Printf("%d. %s\n", index+1, task)
			}
		case "read":
			for index, task := range tasks {
				fmt.Printf("%d. %s\n", index+1, task)
			}
		case "update":
			var updateIndex int
			fmt.Print("Which index to update?: ")
			_, err := fmt.Scanf("%d", &updateIndex)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			fmt.Print("What is the new name of the task?: ")
			taskNameUpdate, _, err := reader.ReadLine()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			tasks[updateIndex-1] = string(taskNameUpdate)
			for index, task := range tasks {
				fmt.Printf("%d. %s\n", index+1, task)
			}
		case "delete":
			var deleteIndex int
			fmt.Print("Which index to delete?: ")
			_, err := fmt.Scanf("%d", &deleteIndex)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			if len(tasks) == 1 {
				tasks = tasks[:0]
			} else {
				tasks = append(tasks[:deleteIndex-1], tasks[deleteIndex:]...)
			}
			for index, task := range tasks {
				fmt.Printf("%d. %s\n", index+1, task)
			}
		default:
			fmt.Println("What?")
		}
	}
}
