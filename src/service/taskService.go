package service

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/repository"
)

type TaskService struct {
	taskRepository repository.ITaskRepository
}

type ITaskService interface {
	NewTask(name string) error
	ListTasks() ([]entity.Task, error)
}

func NewTaskService(taskRepository repository.ITaskRepository) ITaskService {
	return &TaskService{taskRepository}
}

func (tr *TaskService) NewTask(name string) error {
	task := entity.Task{
		Name: name,
	}
	return tr.taskRepository.NewTask(task)
}

func (tr *TaskService) ListTasks() ([]entity.Task, error) {
	return tr.taskRepository.GetTasks()
}
