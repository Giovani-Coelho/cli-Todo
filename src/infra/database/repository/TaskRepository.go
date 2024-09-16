package repository

import (
	"errors"
	"fmt"

	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database"
	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/entity"
)

type MemoryRepository struct {
	tasks []entity.Task
}

type ITaskRepository interface {
	NewTask(task entity.Task) error
	GetTasks() ([]entity.Task, error)
	DeleteTask(id string) error
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks: []entity.Task{},
	}
}

func (r *MemoryRepository) NewTask(task entity.Task) error {
	stmt := "INSERT INTO task (Name, Done, CreatedAt) VALUES (?, ?, ?)"
	_, err := database.SQL.Exec(stmt, task.Name, task.Done, task.CreatedAt)
	if err != nil {
		return fmt.Errorf("error inserting task: %w", err)
	}
	return nil
}

func (r *MemoryRepository) GetTasks() ([]entity.Task, error) {
	stmt := "SELECT ID, Name, Done, CreatedAt FROM task"
	rows, err := database.SQL.Query(stmt)
	if err != nil {
		return nil, err
	}

	tasks := []entity.Task{}

	for rows.Next() {
		task := entity.Task{}

		err := rows.Scan(&task.ID, &task.Name, &task.Done, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *MemoryRepository) DeleteTask(id string) error {
	stmt := "DELETE FROM task WHERE ID = ?"

	result, err := database.SQL.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}
