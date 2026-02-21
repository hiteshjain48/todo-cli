package model

import "errors"

var (
	ErrorTaskNotFound = errors.New("task not found")
	ErrInvalidStatus  = errors.New("invalid status")
	ErrInvalidInput   = errors.New("invalid input")
)

