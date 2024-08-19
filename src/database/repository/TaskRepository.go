package repository

import "github.com/Giovani-Coelho/Todo-CLI/src/database/entity"

type MemoryRepository struct {
	tasks []entity.Task
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks: []entity.Task{},
	}
}

func (r *MemoryRepository) AddTask(task entity.Task) error {
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *MemoryRepository) GetTasks() ([]entity.Task, error) {
	return r.tasks, nil
}
