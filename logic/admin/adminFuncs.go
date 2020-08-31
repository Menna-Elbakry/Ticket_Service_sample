package admin

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	Email  string = "fifi@yahoo.com"
	Passwd string = "MyDarkSecret"
	Name   string = "Menna"
	Mail   string = "lili@yahoo.com"
	ID            = uuid.Must(uuid.FromString("fb0f4444-32e1-4a13-8f07-89647469abd6"))
	//ID               = uuid.Must(uuid.FromString("bb0372b2-7e86-470a-91f0-2c65a415e5c5"))
	OpID             = uuid.Must(uuid.FromString("2e12da87-c95f-4ac8-be09-579f9cf43743"))
	Ticket_ID        = uuid.Must(uuid.FromString("d41797fa-7e7c-480b-a073-53dff747f18c"))
	Details   string = "details"
)

//HashPassword to encrypt the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
