package admin

import (
	database "task/database/implementation"
	"task/model"
)

func UpdateUserMail() {
	newUSR := &model.User{
		Email: Email,
		ID:    ID,
	}
	database.UpdateUserMail(newUSR)
}
