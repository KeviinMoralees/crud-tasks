package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Task struct {
	Id          string
	Name        string
	Description string
}

func NewTask(task Task) Task {
	return Task{
		Id:          uuid.New().String(),
		Name:        task.Name,
		Description: task.Description,
	}
}

func (t *Task) Print() {
	fmt.Println(t)
}
