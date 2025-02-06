package inmemory

import (
	"sync"
	
	storageModel "github.com/Dindonpingpong/todo-cli/storage/model"
	"github.com/Dindonpingpong/todo-cli/storage"
)

var _ storage.Storer = (*Storage)(nil)

type Storage struct {
	mu sync.Mutex
	tasks []storageModel.Task
	nextID int64
}

var Store = Storage{
	tasks:  []storageModel.Task{},
	nextID: 1,
}

func (s *Storage) LoadTasks() ([]storageModel.Task, error) {
	return nil, nil
}

func (s *Storage) SaveTasks([]storageModel.Task) error {
	return nil
}

func (s *Storage) AddTask(text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks = append(s.tasks, storageModel.Task{ID: s.nextID, Text: text})
	s.nextID++
	return nil
}

func (s *Storage) GetTasks() ([]storageModel.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.tasks, nil
}

func (s *Storage) RemoveTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
		}
	}

	return nil
}