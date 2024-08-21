package usecase

import (
	"fmt"

	"github.com/Giovani-Coelho/Todo-CLI/src/app/task/services"
	"github.com/Giovani-Coelho/Todo-CLI/src/database/repository"
)

type TaskUseCase struct {
	taskService *services.TaskService
}

func NewTaskUseCase() *TaskUseCase {
	taskRepository := repository.NewMemoryRepository()
	taskService := services.NewTaskService(taskRepository)

	return &TaskUseCase{
		taskService: taskService,
	}
}

func (us *TaskUseCase) CreateTask(args []string) string {
	err := us.taskService.AddTask(args[0])
	if err != nil {
		fmt.Println("Error:", err)
	}

	return "New task has been created!"
}
