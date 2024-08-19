package service

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/database/entity"
	"github.com/Giovani-Coelho/Todo-CLI/src/database/repository"
)

type TaskService struct {
	repo repository.ITaskRepository
}

func NewTaskService(repo repository.ITaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(name string) error {
	task := entity.Task{
		Name: name,
	}
	return s.repo.AddTask(task)
}

func (s *TaskService) ListTasks() ([]entity.Task, error) {
	return s.repo.GetTasks()
}
