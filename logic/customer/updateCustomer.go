package customer

import (
	"task/database"
)

//UpdateCustomer to update user name and email
func UpdateCustomer() {
	hash, _ := HashPassword(Passwd)
	cst := &database.Users{
		ID:       ID,
		Name:     Name,
		Password: hash,
		Email:    Mail,
	}
	cst.UpdateUsername()
	cst.UpdateUserPassword()
	cst.UpdateUserMail()
}
