package customer

import (
	"task/database"
)

//DeleteTicketByID to remove spicific ticket from database
func DeleteTicketByID() {
	newTCT := &database.Tickets{
		TicketID: Ticket_ID,
	}
	newTCT.DeleteTicketByID()
}
