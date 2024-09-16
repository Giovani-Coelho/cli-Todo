package taskUseCase

import "fmt"

func (tr *TaskUseCase) DeleteTask(id string) error {
	err := tr.taskRepository.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Task successfully removed")
	return nil
}
