package admin

import (
	"task/database"
	"time"

	uuid "github.com/satori/go.uuid"
)

//AddNewUser to insert new user to database
func AddNewUser() {
	hash, _ := HashPassword(Passwd)
	newUSR := &database.Users{
		ID:        uuid.NewV4(),
		Name:      "lily",
		Email:     "lily@yahoo.com",
		Password:  hash,
		Role:      "customer",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUSR.AddNewUser()
}
