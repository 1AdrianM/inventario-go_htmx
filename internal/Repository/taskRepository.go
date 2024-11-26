package repository

import (
	"errors"
	db "github/1AdrianM/go_inventario/internal/Db"
	model "github/1AdrianM/go_inventario/internal/Model"
)

func CreateTask(task *model.Task) error {
	if task.Title == "" || task.Text == "" {
		return errors.New("title and text are required")
	}
	if err := db.Db.Create(&task).Error; err != nil {
		return errors.New("not able to create task")
	}

	return nil
}
func UpdateTask(id uint, UpdateTask *model.Task) error {
	var task model.Task
	if err := db.Db.First(&task, id).Error; err != nil {
		return errors.New("task not found")
	}
	if UpdateTask.Title != "" {
		UpdateTask.Title = task.Title
	}
	if UpdateTask.Text != "" {
		UpdateTask.Text = task.Title
	}
	UpdateTask.Completed = task.Completed
	if err := db.Db.Save(&UpdateTask).Error; err != nil {
		return errors.New("not able to update task")
	}
	return nil
}
func DeleteTask(id uint) error {
	if id <= 0 {
		return errors.New("id is required")
	}
	if err := db.Db.Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return errors.New("not able to delete task")
	}
	return nil
}
func GetTaskById(id uint) (*model.Task, error) {
	var task model.Task
	if id <= 0 {
		return nil, errors.New("id is required")
	}
	if err := db.Db.First(&task, id).Error; err != nil {
		return nil, errors.New("not able to get task")
	}
	return &task, nil
}
func GetTasks() ([]*model.Task, error) {
	var tasks []*model.Task
	if err := db.Db.Find(&tasks).Error; err != nil {
		return nil, errors.New("not able to get tasks")
	}
	return tasks, nil
}
