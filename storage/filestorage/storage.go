package filestorage

import (
	"fmt"
	"os"
	"sync"
	"encoding/json"
	
	storageModel "github.com/Dindonpingpong/todo-cli/storage/model"
	"github.com/Dindonpingpong/todo-cli/storage"
)

const taskFile = "tasks.json"

var _ storage.Storer = (*Storage)(nil)

type Storage struct {
	mu sync.Mutex
	tasks []storageModel.Task
}

func NewStorage() *Storage {
	s := &Storage{}
	s.LoadTasks()

	return s
}

func (s *Storage) LoadTasks() ([]storageModel.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(taskFile)
	if err != nil {
		s.tasks = []storageModel.Task{}
		return s.tasks, nil
	}

	if err := json.Unmarshal(file, &s.tasks); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return s.tasks, nil
}

func (s *Storage) SaveTasks(tasks []storageModel.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(taskFile, data, 0644); err != nil {
		return err
	}

	s.tasks = tasks

	return nil
}

func (s *Storage) AddTask(text string) error {
	newTask := storageModel.Task{ID: len(s.tasks) + 1, Text: text}
	s.tasks = append(s.tasks, newTask)

	return s.SaveTasks(s.tasks)
}

func (s *Storage) RemoveTask(id int) error {
	if id < 1 || id > len(s.tasks) {
		return fmt.Errorf("invalid task ID")
	}

	s.tasks = append(s.tasks[:id-1], s.tasks[id:]...)

	return s.SaveTasks(s.tasks)
}


func (s *Storage) GetTasks() ([]storageModel.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.tasks, nil
}
