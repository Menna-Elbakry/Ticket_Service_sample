package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateUserName(id uuid.UUID, name string) (string, error) {
	newUSR := &model.User{
		ID:   id,
		Name: name,
	}
	return database.UpdateUsername(newUSR)
}
