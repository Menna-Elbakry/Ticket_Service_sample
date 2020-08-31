package admin

import (
	"task/database"
)

//GetUserByRole to select all users with role = customer
func GetUserByRole(role string) {
	newUSR := &database.Users{
		Role: role,
	}
	newUSR.GetUserByRole(role)
}
