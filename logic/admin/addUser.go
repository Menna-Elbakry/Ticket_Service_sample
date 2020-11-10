package admin

import (
	database "task/database/implementation"
	"task/logic"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewUser(name string, pass string, mail string, role model.RoleEnum) (uuid.UUID, error) {
	hash, _ := logic.HashPassword(pass)
	newUSR := &model.User{
		ID:       uuid.NewV4(),
		Name:     name,
		Email:    mail,
		Password: hash,
		Role:     role,
	}
	return database.AddNewUser(newUSR)
}
