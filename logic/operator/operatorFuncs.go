package operator

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	Email     string = "fifi@yahoo.com"
	Passwd    string = "123456"
	Name      string = "Menna"
	OpID             = uuid.Must(uuid.FromString("2e12da87-c95f-4ac8-be09-579f9cf43743"))
	Ticket_ID        = uuid.Must(uuid.FromString("d41797fa-7e7c-480b-a073-53dff747f18c"))
)

//HashPassword to encrypt the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
