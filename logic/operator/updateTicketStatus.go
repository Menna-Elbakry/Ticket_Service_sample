package operator

import (
	"task/database"
)

//UpdateStatus to update the status of the ticket and add opertaor id
func UpdateStatus() {
	newTCT := &database.Tickets{
		TicketID:   Ticket_ID,
		OperatorID: OpID,
	}
	newTCT.UpdateOperatorIDAndStatus()
}
