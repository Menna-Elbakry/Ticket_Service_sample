package customer

import (
	"task/database"
)

//UpdateTicketDetailsByID to update ticket content
func UpdateTicketDetailsByID() {
	newTCT := &database.Tickets{
		TicketID: Ticket_ID,
		Details:  Details,
	}
	newTCT.UpdateTicketDetails()
}
