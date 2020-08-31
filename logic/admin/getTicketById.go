package admin

import (
	"task/database"
)

//GetTicketByID to select specific ticket
func GetTicketByID() {
	newTCT := &database.Tickets{
		TicketID: Ticket_ID,
	}
	newTCT.GetTicketByID()
}
