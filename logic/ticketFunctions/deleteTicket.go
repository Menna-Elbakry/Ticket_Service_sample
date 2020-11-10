package ticket

import (
	database "task/database/implementation"
	"task/model"
)

func DeleteTicket() {
	newTCT := &model.Ticket{
		TicketID: TID,
	}
	database.DeleteTicket(newTCT)
}