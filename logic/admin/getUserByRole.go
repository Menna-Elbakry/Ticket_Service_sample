package admin

import (
	database "task/database/implementation"
	"task/model"
)

func GetUsersByRole() {
	newUSR := &model.User{
		Role: Role,
	}
	database.GetUsersByRole(newUSR)
}
