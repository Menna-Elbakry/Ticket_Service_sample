package logic

import (
	database "task/database/implementation"
	"task/model"
)

func GetUserById() {
	newUSR := &model.User{
		ID: ID,
	}
	database.GetUserById(newUSR)
}
