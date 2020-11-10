package admin

import (
	database "task/database/implementation"
	"task/model"
)

func GetUsersByRole(role model.RoleEnum) ([]model.User, error){
	newUSR := &model.User{
		Role: role,
	}
	return database.GetUsersByRole(newUSR)
}
