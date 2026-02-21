package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/hiteshjain48/todo-cli/model"
)

var fileName = "tasks.json"

func AddTask(task string) {
	tk := model.Task{
		Id:          0,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	err := addHandler(fileName, tk)
	if err != nil {
		fmt.Println("error adding task ", err)
		return
	}
	fmt.Println("Task added.")
}

func UpdateTask(id int, task string) {
	err := updateHandler(fileName, id, task)
	if err != nil {
		fmt.Println("error updating task ", err)
		return
	}
	fmt.Println("Task updated.")
}

func DeleteTask(id int) {
	err := deleteHandler(fileName, id)
	if err != nil {
		fmt.Println("error deleting task ", err)
		return
	}
	fmt.Println("Task deleted.")
}

func MarkProgress(id int) {
	err := statusHandler(fileName, id, "in-progress")
	if err != nil {
		fmt.Println("error marking task ", err)
		return
	}
	fmt.Println("Marked as in progress.")
}

func MarkDone(id int) {
	err := statusHandler(fileName, id, "done")
	if err != nil {
		fmt.Println("error marking task ", err)
		return
	}
	fmt.Println("Marked as done.")
}

func ListTasks(ids ...string) {

	if len(ids) > 0 {
		err := listHandler(fileName, ids[0])
		if err != nil {
			fmt.Println("error listing task ", err)
			return
		}
		fmt.Println("Listing task with status ", ids[0])
	} else {
		err := listHandler(fileName, "")
		if err != nil {
			fmt.Println("error listing task ", err)
			return
		}
		fmt.Println("Listing all tasks...")
	}

}

func addHandler(filename string, tk model.Task) error {
	byteValue, err := os.ReadFile(filename)
	var taskList model.Tasks
	if err != nil {
		if os.IsNotExist(err) {
			taskList = model.Tasks{Tasks: []model.Task{}}
		} else {
			return fmt.Errorf("error reading file %w", err)
		}
	} else if len(byteValue) > 0 {
		if err := json.Unmarshal(byteValue, &taskList); err != nil {
			return fmt.Errorf("error unmarshalling %w", err)
		}
	}
	id := taskList.GetNextId()
	tk.Id = id
	taskList.Tasks = append(taskList.Tasks, tk)

	updatedTaskList, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling %w", err)
	}

	err = os.WriteFile(filename, updatedTaskList, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil
}

func updateHandler(filename string, id int, tk string) error {
	byteValue, err := os.ReadFile(filename)
	var taskList model.Tasks
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Cant update without adding")
		} else {
			return fmt.Errorf("error reading file %w", err)
		}
	} else if len(byteValue) > 0 {
		if err = json.Unmarshal(byteValue, &taskList); err != nil {
			return fmt.Errorf("error unmarshaling %w", err)
		}
	}
	found := false
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == id {
			taskList.Tasks[i].Description = tk
			taskList.Tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}
	updatedTaskList, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling %w", err)
	}
	err = os.WriteFile(filename, updatedTaskList, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil

}

func deleteHandler(filename string, id int) error {
	byteValue, err := os.ReadFile(filename)
	var taskList model.Tasks
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Cant delete without adding")
		} else {
			return fmt.Errorf("error reading file %w", err)
		}
	} else if len(byteValue) > 0 {
		if err = json.Unmarshal(byteValue, &taskList); err != nil {
			return fmt.Errorf("error unmarshaling %w", err)
		}
	}
	var index = -1
	for i, t := range taskList.Tasks {
		if t.Id == id {
			index = i
			break
		}
	}
	if index != -1 {
		taskList.Tasks = slices.Delete(taskList.Tasks, index, index+1)
	} else {
		return fmt.Errorf("task with id %d not found", id)
	}
	updatedTaskList, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling %w", err)
	}
	err = os.WriteFile(filename, updatedTaskList, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil
}

func statusHandler(filename string, id int, status string) error {
	byteValue, err := os.ReadFile(filename)
	var taskList model.Tasks
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Cant update without adding")
		} else {
			return fmt.Errorf("error reading file %w", err)
		}
	} else if len(byteValue) > 0 {
		if err = json.Unmarshal(byteValue, &taskList); err != nil {
			return fmt.Errorf("error unmarshaling %w", err)
		}
	}

	found := false
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == id {
			taskList.Tasks[i].Status = status
			taskList.Tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}

	updatedTaskList, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling %w", err)
	}
	err = os.WriteFile(filename, updatedTaskList, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil
}

func listHandler(filename string, status string) error {
	byteValue, err := os.ReadFile(filename)
	var taskList model.Tasks
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Cant list without adding")
		} else {
			return fmt.Errorf("error reading file %w", err)
		}
	} else if len(byteValue) > 0 {
		if err = json.Unmarshal(byteValue, &taskList); err != nil {
			return fmt.Errorf("error unmarshaling %w", err)
		}
	}

	for _, t := range taskList.Tasks {
		if status == "" {
			fmt.Println(t)
			continue
		}
		if t.Status == status {
			fmt.Println(t)
		}
	}
	return nil
}
