package admin

import (
	"task/database"
)

//UpdateUserName to update user name
func UpdateUserName() {
	hash, _ := HashPassword(Passwd)
	usr := &database.Users{
		ID:       ID,
		Password: hash,
		Name:     Name,
		Email:    Mail,
	}
	usr.UpdateUsername()
}
