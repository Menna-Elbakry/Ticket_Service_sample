package admin

import "task/database"

//UpdateTicketDetails to update ticket content
func UpdateTicketDetails() {
	newTCT := &database.Tickets{
		TicketID: Ticket_ID,
		Details:  Details,
	}
	newTCT.UpdateTicketDetails()
}
