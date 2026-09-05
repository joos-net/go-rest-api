package storage

import "tasks-api/internal/models"

// Storage определяет интерфейс для работы с задачами
type Storage interface {
	List() []models.Task
	Create(task models.Task) (models.Task, error)
	Get(id int) (models.Task, bool)
	Update(id int, task models.Task) (models.Task, error)
	Delete(id int) error
}
