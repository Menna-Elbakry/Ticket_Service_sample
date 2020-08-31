package main

import (
	log "task/public/login"
)

var (
	mail = "lily@yahoo.com"
	pass = "MyDarkSecret"
)

func main() {
	//Run Once If Table Dosn't Exist
	//database.CreateTables()

	// admin.UpdateUserName()
	log.Login(mail, pass)
}
