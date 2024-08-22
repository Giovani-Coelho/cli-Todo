package repository

import "github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"

type MemoryRepository struct {
	tasks []entity.Task
}

type ITaskRepository interface {
	NewTask(task entity.Task) error
	GetTasks() ([]entity.Task, error)
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks: []entity.Task{},
	}
}

func (r *MemoryRepository) NewTask(task entity.Task) error {
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *MemoryRepository) GetTasks() ([]entity.Task, error) {
	return r.tasks, nil
}
