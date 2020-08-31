package operator

import (
	"task/database"
)

//UpdateOperator to update user password and name
func UpdateOperator() {
	hash, _ := HashPassword(Passwd)
	ort := &database.Users{
		ID:       OpID,
		Name:     Name,
		Password: hash,
		Email:    Email,
	}
	ort.UpdateUsername()
	ort.UpdateUserPassword()
	ort.UpdateUserMail()
}
