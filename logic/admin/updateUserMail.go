package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateUserMail(id uuid.UUID, mail string) (string, error) {
	newUSR := &model.User{
		Email: mail,
		ID:    id,
	}
	return database.UpdateUserMail(newUSR)
}
