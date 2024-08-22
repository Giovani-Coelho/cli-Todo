/*
Copyright © 2024 Giovani Coelho
*/
package main

import (
	"github.com/Giovani-Coelho/Todo-CLI/cmd"
	"github.com/Giovani-Coelho/Todo-CLI/src/di"
	controller "github.com/Giovani-Coelho/Todo-CLI/src/infra/controller/task"
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/repository"
	"github.com/Giovani-Coelho/Todo-CLI/src/service"
)

func main() {
	initDependencies()
	cmd.Execute()
}

func initDependencies() {
	taskRepository := repository.NewMemoryRepository()
	taskService := service.NewTaskService(taskRepository)
	di.TaskController = controller.NewTaskController(taskService)
}
