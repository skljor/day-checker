package models

import (
	"errors"

	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB, task *Task) error {
	if task.Title == "" {
		return errors.New("Title can't be empty")
	}
	return db.Create(task).Error
}

func GetUserTasks(db *gorm.DB, userID uint) ([]Task, error) {
	var tasks []Task
	err := db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func ToogleTaskStatus(db *gorm.DB, taskID uint, isDone bool) error {
	return db.Model(&Task{}).Where("id = ?", taskID).Update("done", isDone).Error
}
