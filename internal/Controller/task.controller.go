package controller

import (
	"errors"
	model "github/1AdrianM/go_inventario/internal/Model"
	repository "github/1AdrianM/go_inventario/internal/Repository"
)

func CreateTask(task *model.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	if err := repository.CreateTask(task); err != nil {
		return errors.New("something bad Happend, not able create task ")
	}
	return nil
}
func GetAllTask() ([]*model.Task, error) {
	tasks, err := repository.GetTasks()

	if err != nil {
		return nil, errors.New(" not possible to fetch tasks data")
	}
	return tasks, nil
}
func GetTaskById(id uint) (*model.Task, error) {
	if id == 0 {
		return nil, errors.New("id is 0")
	}

	task, err := repository.GetTaskById(id)
	if err != nil {
		return nil, errors.New("not possible to fetch task data")
	}
	return task, nil
}

func UpdateTask(id uint, task *model.Task) error {
	if id == 0 || task == nil {
		return errors.New("not possible to find ID or task")
	}

	if err := repository.UpdateTask(id, task); err != nil {
		return errors.New("not possible to update task")
	}
	return nil
}

func DeleteTask(id uint) error {
	if id == 0 {
		return errors.New("id is 0")
	}
	if err := repository.DeleteTask(id); err != nil {
		return errors.New("not possible to delete task")
	}
	return nil
}
