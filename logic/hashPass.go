package logic

import (
	"golang.org/x/crypto/bcrypt"
)


//HashPassword to encrypt the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
