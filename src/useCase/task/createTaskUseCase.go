package taskUseCase

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"
)

func (tr *TaskUseCase) NewTask(name string) string {
	task := entity.Task{Name: name}
	err := tr.taskRepository.NewTask(task)

	if err != nil {
		return "Error on create a new task!"
	}

	return "New task has been created! Task: " + name
}
