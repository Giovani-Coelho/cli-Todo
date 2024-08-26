package taskUseCase

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/repository"
)

var Task ITaskUseCase

type ITaskUseCase interface {
	NewTask(name string) string
	ListTasks() ([]entity.Task, error)
}

type TaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewTaskUseCase(
	taskRepository repository.ITaskRepository,
) ITaskUseCase {
	return &TaskUseCase{taskRepository}
}

func InitDependencies() {
	taskRepository := repository.NewMemoryRepository()
	Task = NewTaskUseCase(taskRepository)
}
