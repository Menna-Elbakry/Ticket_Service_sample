package admin

import (
	"task/database"
)

//GetAllUsers to select all users
func GetAllUsers() {
	allUsr := &database.Users{}
	allUsr.GetAllUsers()
}
