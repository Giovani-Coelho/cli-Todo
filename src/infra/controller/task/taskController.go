package controller

import "github.com/Giovani-Coelho/Todo-CLI/src/service"

type TaskController struct {
	taskService service.ITaskService
}

func NewTaskController(
	taskService service.ITaskService,
) *TaskController {
	return &TaskController{taskService}
}

func (tc *TaskController) NewTask(name string) string {
	err := tc.taskService.NewTask(name)

	if err != nil {
		return "Error on create a new task!"
	}

	return "New task has been created! New Task: " + name
}
