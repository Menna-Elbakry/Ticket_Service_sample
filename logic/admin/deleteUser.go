package admin

import (
	"task/database"
)

//DeleteUser to delete user
func DeleteUser() {

	newUSR := &database.Users{
		ID: ID,
	}
	newUSR.DeleteUser()
}
