package admin

import (
	database "task/database/implementation"
	"task/logic"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewUser() {
	hash, _ := logic.HashPassword(Pass)
	newUSR := &model.User{
		ID:       uuid.NewV4(),
		Name:     Name,
		Email:    Email,
		Password: hash,
		Role:     Role,
	}
	database.AddNewUser(newUSR)
}
