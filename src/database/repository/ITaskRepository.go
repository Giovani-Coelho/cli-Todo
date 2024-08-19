package repository

import "github.com/Giovani-Coelho/Todo-CLI/src/database/entity"

type ITaskRepository interface {
	AddTask(task entity.Task) error
	GetTasks() ([]entity.Task, error)
}
