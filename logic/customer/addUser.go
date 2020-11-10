package customer

import (
	database "task/database/implementation"
	"task/logic"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewUser(name string, pass string, mail string) (uuid.UUID, error) {
	hash, _ := logic.HashPassword(pass)
	newUSR := &model.User{
		ID:       uuid.NewV4(),
		Name:     name,
		Email:    mail,
		Password: hash,
		Role:     model.Customer,
	}
	return database.AddNewUser(newUSR)
}
