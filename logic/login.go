package logic

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func Login(mail string,pass string) (uuid.UUID, error) {
	newUSR := &model.User{
		Email:    mail,
		Password: pass,
	}
	return database.Login(newUSR)
}
