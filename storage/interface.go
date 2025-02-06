package storage

import "github.com/Dindonpingpong/todo-cli/storage/model"

type Storer interface {
	LoadTasks() ([]model.Task, error)
	SaveTasks([]model.Task) error
	AddTask(string) error
	RemoveTask(int) error
	GetTasks() ([]model.Task, error)
}