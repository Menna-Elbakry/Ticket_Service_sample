package ticket

import (
	database "task/database/implementation"
	"task/model"
)

func UpdateTicketDetails() {
	newTCT := &model.Ticket{
		TicketID: TID,
		Details: Details,
	}
	database.UpdateTicketDetails(newTCT)
}