package customer

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateUserName(id uuid.UUID, name string) (string, error) {
	newUSR := &model.User{
		Name: name,
		ID:   id,
	}
	return database.UpdateUsername(newUSR)
}
