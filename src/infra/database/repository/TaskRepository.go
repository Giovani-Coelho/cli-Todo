package repository

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database"
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"
)

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
	stmt := "SELECT ID, Name, Done, CreatedAt FROM task"
	rows, err := database.SQL.Query(stmt)
	if err != nil {
		return nil, err
	}

	tasks := []entity.Task{}

	for rows.Next() {
		task := entity.Task{}

		err := rows.Scan(&task.ID, &task.Name, &task.Done, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
