package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hiteshjain48/todo-cli/tasks"
)

func main() {
	args := os.Args[1:]
	nArgs := len(args)
	if nArgs == 0 {
		fmt.Println("Please enter a valid command...")
		return
	}

	switch args[0] {
	case "add":
		if nArgs != 2 {
			fmt.Println("Invaild number of arguments...")
			return
		}
		fmt.Println("Adding task...")
		tasks.AddTask(args[1])
	case "update":
		if nArgs != 3 {
			fmt.Println("Invaild number of arguments...")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error during fetching id... ", err)
			return
		}
		fmt.Println("Updating task...")
		tasks.UpdateTask(id, args[2])
	case "delete":
		if nArgs != 2 {
			fmt.Println("Invaild number of arguments...")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error during fetching id... ", err)
			return
		}
		fmt.Println("Deleting task...")
		tasks.DeleteTask(id)
	case "mark-in-progress":
		if nArgs != 2 {
			fmt.Println("Invaild number of arguments...")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error during fetching id... ", err)
			return
		}
		fmt.Println("Marking in progress...")
		tasks.MarkProgress(id)
	case "mark-done":
		if nArgs != 2 {
			fmt.Println("Invaild number of arguments...")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error during fetching id... ", err)
			return
		}
		fmt.Println("Marking as done...")
		tasks.MarkDone(id)
	case "list":
		if nArgs > 2 {
			fmt.Println("Invalid number of arguments")
			return
		}
		fmt.Println("Listing task...")
		switch nArgs {
		case 1:
			tasks.ListTasks()
		case 2:
			tasks.ListTasks(args[1])
		}
	default:
		fmt.Println("Please enter a valid command...")
		return
	}
}
