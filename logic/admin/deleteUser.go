package admin

import (
	database "task/database/implementation"
	"task/model"
)

func DeleteUser() {
	newUSR := &model.User{
		ID: ID,
	}
	database.DeleteUser(newUSR)
}
