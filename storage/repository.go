package storage

import "github.com/hiteshjain48/todo-cli/model"

type Repository interface {
	Load() (model.Tasks, error)
	Save(model.Tasks) error
}