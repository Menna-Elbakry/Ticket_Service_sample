package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetUserById(uId uuid.UUID) (string, error) {
	newUSR := &model.User{
		ID: uId,
	}
	return database.GetUserById(newUSR)
}
