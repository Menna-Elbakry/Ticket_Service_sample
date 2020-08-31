package customer

import (
	"task/database"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	//Role always assigned as customer
	Role string = "customer"
)

//AddUser to insert new user to database
func AddUser() {
	newUSR := &database.Users{
		ID:        uuid.NewV4(),
		Name:      "gigi",
		Email:     "gigi@yahoo.com",
		Password:  "123456",
		Role:      Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUSR.AddNewUser()
}
