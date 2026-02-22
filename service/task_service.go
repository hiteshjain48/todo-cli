package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/hiteshjain48/todo-cli/model"
	"github.com/hiteshjain48/todo-cli/storage"
)

type TaskService struct {
	repo storage.Repository
}

func NewTaskService(repo storage.Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Add(description string) (model.Task, error) {
	description = strings.TrimSpace(description)

	if description == "" {
		return model.Task{}, model.ErrInvalidInput
	}

	list, err := s.repo.Load()
	if err != nil {
		return model.Task{}, err
	}

	now := time.Now()
	task := model.Task{
		ID:				list.GetNextID(),
		Description: 	description,
		Status: 		model.StatusTodo,
		CreatedAt:  	now,
		UpdatedAt: 		now,	
	}

	list.Tasks = append(list.Tasks, task)
	if err := s.repo.Save(list); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *TaskService) Update(id int, description string) error {
	description = strings.TrimSpace(description)
	if id <= 0 || description == "" {
		return model.ErrInvalidInput
	}

	list, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range list.Tasks {
		if list.Tasks[i].ID == id {
			list.Tasks[i].Description = description
			list.Tasks[i].UpdatedAt = time.Now()
			return s.repo.Save(list)
		}
	}

	return fmt.Errorf("%w: id=%d", model.ErrorTaskNotFound, id)
}

func (s *TaskService) Delete(id int) error {
	if id <= 0 {
		return model.ErrInvalidInput
	}

	list, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range list.Tasks {
		if list.Tasks[i].ID == id {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			return s.repo.Save(list)
		}
	}

	return fmt.Errorf("%w: id=%d", model.ErrorTaskNotFound, id)
}

func (s *TaskService) SetStatus(id int, status model.Status) error {
	if id <= 0 {
		return model.ErrInvalidInput
	}
	if !status.Valid() {
		return model.ErrInvalidStatus
	}

	list, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range list.Tasks{
		if list.Tasks[i].ID == id {
			list.Tasks[i].Status = status
			list.Tasks[i].UpdatedAt = time.Now()
			return s.repo.Save(list)
		}
	}

	return fmt.Errorf("%w: id=%d", model.ErrorTaskNotFound, id)
}

func (s *TaskService) List(filter *model.Status) ([]model.Task, error) {
	list, err := s.repo.Load()
	if err != nil {
		return nil, err
	}

	if filter == nil {
		return list.Tasks, nil
	}

	if !filter.Valid() {
		return nil, model.ErrInvalidStatus
	}

	out := make([]model.Task, 0, len(list.Tasks))

	for _, t := range list.Tasks {
		if t.Status == *filter {
			out = append(out, t)
		}
	}
	return out, nil
}