package admin

import (
	database "task/database/implementation"
)

func GetAllUsers() {
	database.GetAllUsers()
}
