package admin

import (
	"task/database"
)

//GetAllTickets to select all tickets exist
func GetAllTickets() {
	newTCT := &database.Tickets{}
	newTCT.GetAllTickets()
}
