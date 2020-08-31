package log

import "task/database"

//Login
func Login(mail, pass string) {

	log := &database.Users{}
	log.Login(mail, pass)

}
