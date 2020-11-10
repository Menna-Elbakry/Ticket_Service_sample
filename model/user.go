package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type RoleEnum string

const (
	Admin    RoleEnum = "Admin"
	Operator RoleEnum = "Operator"
	Customer RoleEnum = "Customer"
)

//User table
type User struct {
	ID       uuid.UUID `json:"id,pk,type:uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email,unique"`
	Password string    `json:"password"`
	Role     RoleEnum  `json:"role"`
}

func (u *User) Validate() error {
	if u.ID == uuid.Nil {
		return errors.New("should have userId")
	}
	if u.Name == " " {
		return errors.New("should have Name")
	}
	if u.Email == " " {
		return errors.New("should have Email")
	}
	if u.Password == " " {
		return errors.New("should have Password")
	}
	if u.Role == " " {
		return errors.New("should have Role")
	}
	return nil
}
