package admin

import (
	"task/database"
)

//UpdateUserPassword to update user password
func UpdateUserPassword() {
	hash, _ := HashPassword(Passwd)
	newUSR := &database.Users{
		ID:       ID,
		Password: hash,
	}
	newUSR.UpdateUserPassword()
}
