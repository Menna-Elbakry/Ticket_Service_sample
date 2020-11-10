package customer

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func DeleteUser(id uuid.UUID) (uuid.UUID, error) {
	newUSR := &model.User{
		ID: id,
	}
	return database.DeleteUser(newUSR)
}
