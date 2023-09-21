package action

import (
	"tasked/database"
	"tasked/model"
	"time"
)

type CreateTaskInput struct {
	UserId    uint
	Name      string `json:"name" validate:"required"`
	ExpiresAt string `json:"expires_at" validate:"required"`
}

func CreateTask(input *CreateTaskInput) (*model.Task, error) {
	if err := Validate(input); err != nil {
		return &model.Task{}, err
	}

	exp, err := time.Parse(time.DateTime, input.ExpiresAt)

	if err != nil {
		return &model.Task{}, err
	}

	task := &model.Task{
		Name:      input.Name,
		ExpiresAt: exp,
		UserId:    int(input.UserId),
	}

	if err := database.ORM().Create(&task).Error; err != nil {
		return &model.Task{}, err
	}

	return task, nil
}
