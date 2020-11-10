package admin

import (
	database "task/database/implementation"
	"task/model"
)

func GetAllUsers() ([]model.User, error) {
	return database.GetAllUsers()
}
