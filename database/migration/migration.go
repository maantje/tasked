package migration

import (
	"tasked/database"
	"tasked/model"
)

func Migrate() {
	// Migrate the schema
	if err := database.ORM().AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		panic("Failed to migrate database")
	}
}
