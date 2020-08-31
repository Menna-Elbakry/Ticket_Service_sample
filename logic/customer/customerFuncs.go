package customer

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	Email     string = "fifi@yahoo.com"
	Passwd    string = "123456"
	Name      string = "Menna"
	Mail      string = "gigi@yahoo.com"
	ID               = uuid.Must(uuid.FromString("fb0f4444-32e1-4a13-8f07-89647469abd6"))
	Ticket_ID        = uuid.Must(uuid.FromString("d41797fa-7e7c-480b-a073-53dff747f18c"))
	Details   string = "details"
)

//HashPassword to encrypt the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
