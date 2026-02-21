package model

type Status string

const (
	StatusTodo 			Status = "todo"
	StatusInProgress	Status = "in-progress"
	StatusDone			Status = "done"
)

func (s Status) Valid() bool {
	return s == StatusTodo || s == StatusInProgress || s == StatusDone
}