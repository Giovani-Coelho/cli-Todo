package taskUseCase

import "github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"

func (tr *TaskUseCase) ListTasks() ([]entity.Task, error) {
	return tr.taskRepository.GetTasks()
}
