/*
Copyright © 2024 Giovani Coelho
*/
package main

import (
	"github.com/Giovani-Coelho/Todo-CLI/cmd"
	taskUseCase "github.com/Giovani-Coelho/Todo-CLI/src/useCase/task"
)

func main() {
	taskUseCase.InitDependencies()
	cmd.Execute()
}
