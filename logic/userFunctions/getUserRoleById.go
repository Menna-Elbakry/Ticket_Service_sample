package logic

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetUserRoleById(Id uuid.UUID) (model.RoleEnum, error) {
	newUSR := &model.User{
		ID: Id,
	}
	return database.GetUserRoleById(newUSR)
}
