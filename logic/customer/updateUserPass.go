package customer

import (
	database "task/database/implementation"
	"task/logic"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateUserPass(id uuid.UUID, pass string) (string, error) {
	hash, _ := logic.HashPassword(pass)
	newUSR := &model.User{
		Password: hash,
		ID:       id,
	}
	return database.UpdateUserPassword(newUSR)
}
