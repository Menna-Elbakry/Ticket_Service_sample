package admin

import (
	"task/database"
)

//GetUserByID to select specific User
func GetUserByID() {
	usr := &database.Users{
		ID: ID,
	}
	usr.GetUserByID()
}
