package admin

import (
	"task/database"
)

//UpdateUserMail to update user email
func UpdateUserMail() {
	usr := &database.Users{
		ID:    ID,
		Email: Email,
	}
	usr.UpdateUserMail()
}
