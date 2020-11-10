package operator

import (
	database "task/database/implementation"
	"task/logic"
	"task/model"
)

func UpdateUserPass() {
	hash, _ := logic.HashPassword(Pass)
	newUSR := &model.User{
		Password: hash,
		ID:       ID,
	}
	database.UpdateUserPassword(newUSR)
}
