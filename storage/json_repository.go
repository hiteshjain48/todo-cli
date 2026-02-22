package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hiteshjain48/todo-cli/model"
)

type JSONRepository struct {
	Path string
}

func (r JSONRepository) Load() (model.Tasks, error) {
	var tasks model.Tasks

	data, err := os.ReadFile(r.Path)
	if err != nil {
		if os.IsNotExist(err) {
			tasks.Tasks = []model.Task{}
			return tasks, nil
		}
		return tasks, fmt.Errorf("read tasks file: %w", err)
	}

	if len(data) == 0 {
		tasks.Tasks = []model.Task{}
		return tasks, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return tasks, fmt.Errorf("unmarshal tasks: %w", err)
	}

	if tasks.Tasks == nil {
		tasks.Tasks = []model.Task{}
	}
	return tasks, nil
}

func (r JSONRepository) Save(tasks model.Tasks) error {
	if tasks.Tasks == nil {
		tasks.Tasks = []model.Task{}
	}

	payload, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	dir := filepath.Dir(r.Path)
	tmp, err := os.CreateTemp(dir, "tasks-*.json")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmp.Name()

	if _, err := tmp.Write(payload); err != nil {
		tmp.Close()
		_ = os.Remove(tmpPath)
		return fmt.Errorf("close temp file: %w", err)
	}

	if err := os.Rename(tmpPath, r.Path); err != nil {
		_ = os.Remove(tmpPath)
		return fmt.Errorf("replace tasks file: %w", err)
	}

	return nil
}

