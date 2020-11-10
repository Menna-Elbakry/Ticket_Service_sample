package logic

import (
	database "task/database/implementation"
	"task/model"
)

func UpdateUserName() {
	newUSR := &model.User{
		Name: Name,
		ID:       ID,
	}
	database.UpdateUsername(newUSR)
}
