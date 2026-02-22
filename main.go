package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/hiteshjain48/todo-cli/model"
	"github.com/hiteshjain48/todo-cli/service"
	"github.com/hiteshjain48/todo-cli/storage"
)

func main() {
	args := os.Args[1:]
	nArgs := len(args)
	if nArgs == 0 {
		fmt.Printf("Please enter a valid command...\n")
		os.Exit(1)
	}
	svc := service.NewTaskService(storage.JSONRepository{Path: "tasks.json"})

	switch args[0] {
	case "add":
		if nArgs != 2 {
			exitf("Invalid number of arguments")
		}
		task, err := svc.Add(args[1])
		exitOnErr(err)
		fmt.Printf("Task added (id=%d). \n", task.ID)
	case "update":
		if nArgs != 3 {
			exitf("Invalid number of arguments")
		}
		id := parseID(args[1])
		err := svc.Update(id, args[2])
		exitOnErr(err)
		fmt.Printf("Task updated.\n")
	case "delete":
		if nArgs != 2 {
			exitf("Invalid number of arguments")
		}
		id := parseID(args[1])
		err := svc.Delete(id)
		exitOnErr(err)
		fmt.Printf("Task deleted.\n")

	case "mark-in-progress":
		if nArgs != 2 {
			exitf("Invalid number of arguments")
		}
		id := parseID(args[1])
		err := svc.SetStatus(id, model.StatusInProgress)
		exitOnErr(err)
		fmt.Printf("Marked as in progress.\n")

	case "mark-done":
		if nArgs != 2 {
			exitf("Invalid number of arguments")
		}
		id := parseID(args[1])
		err := svc.SetStatus(id, model.StatusDone)
		exitOnErr(err)
		fmt.Printf("Marked as done.\n")

	case "list":
		if nArgs > 2 {
			exitf("Invalid number of arguments")
		}
		var filter *model.Status
		if nArgs == 2 {
			s := model.Status(args[1])
			if !s.Valid() {
				exitf("Invalid status: %s", args[1])
			}
			filter = &s
		}
		items, err := svc.List((filter))
		exitOnErr(err)
		for _, t := range items {
			fmt.Printf("%d | %s | %s | created=%s | updated=%s\n", t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"))
		}

	default:
		fmt.Printf("Please enter a valid command...\n")
		return
	}
}

func exitf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}

func exitOnErr(err error) {
	if err == nil {
		return
	}
	switch {
	case errors.Is(err, model.ErrorTaskNotFound):
		exitf("Task not found")
	case errors.Is(err, model.ErrInvalidStatus):
		exitf("Invalid status")
	case errors.Is(err, model.ErrInvalidInput):
		exitf("Invalid input")
	default:
		exitf("Error: %v", err)
	}
}

func parseID(raw string) int {
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		exitf("Invalid id: %s", raw)
	}
	return id
}