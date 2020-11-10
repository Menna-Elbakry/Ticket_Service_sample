package table

import (
	"task/model"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Users table
type User struct {
	tableName struct{}       `sql:"users"`
	ID        uuid.UUID      `sql:"id,pk,type:uuid"`
	Name      string         `sql:"name"`
	Email     string         `sql:"email,unique"`
	Password  string         `sql:"password"`
	Role      model.RoleEnum `sql:"role"`
	CreatedAt time.Time      `sql:"created_at"`
	UpdatedAt time.Time      `sql:"updated_at"`
}

func (u *User) MapToModule() model.User {
	return model.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

func (u User) Fill(user *model.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
