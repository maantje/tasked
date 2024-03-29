package action

import (
	"tasked/database"
	"tasked/model"
	"time"
)

type UpdateTaskInput struct {
	ID        string
	Name      string `json:"name" validate:"required"`
	ExpiresAt string `json:"expires_at" validate:"required"`
}

func UpdateTask(input *UpdateTaskInput) (*model.Task, error) {
	task := &model.Task{}

	if err := Validate(input); err != nil {
		return task, err
	}

	if err := database.ORM().First(&task, input.ID).Error; err != nil {
		return task, err
	}

	exp, err := time.Parse(time.DateTime, input.ExpiresAt)

	if err != nil {
		return &model.Task{}, err
	}

	err = database.ORM().Model(&task).Updates(model.Task{
		Name:      input.Name,
		ExpiresAt: exp,
	}).Error

	if err != nil {
		return task, err
	}

	return task, nil
}
