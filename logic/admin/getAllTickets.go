package admin

import (
	database "task/database/implementation"
)

func GetAllTickets() {
	database.GetAllTickets()
}
