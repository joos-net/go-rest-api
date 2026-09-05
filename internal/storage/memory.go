package storage

import (
	"fmt"
	"sync"
	"tasks-api/internal/models"
	"time"
)

// MemoryStorage реализует интерфейс Storage с использованием in-memory
type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[int]models.Task
	idGen int
}

// NewMemoryStorage создаем новое хранилище
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[int]models.Task),
		idGen: 0,
	}
}

// List возвращает все задачи
func (s *MemoryStorage) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result
}

// Create добавляем новую задачу
func (s *MemoryStorage) Create(task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.idGen++
	task.ID = s.idGen
	task.CreatedAt = time.Now().Format(time.RFC3339)
	s.tasks[task.ID] = task
	return task, nil
}

// Get получаем задачу по ID и флаг
func (s *MemoryStorage) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

// Update обновляем существующую задачу по ID
func (s *MemoryStorage) Update(id int, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, fmt.Errorf("задача с ID %d не найдена", id)
	}

	task.ID = id
	if existingTask, exists := s.tasks[id]; exists {
		if task.CreatedAt == "" {
			task.CreatedAt = existingTask.CreatedAt
		}
	}

	s.tasks[id] = task
	return task, nil
}

// Delete удаляем задачу по ID
func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	delete(s.tasks, id)
	return nil
}
