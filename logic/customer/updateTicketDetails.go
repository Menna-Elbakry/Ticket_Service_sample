package customer

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateTicketDetails(tId uuid.UUID, details string) (string, error) {
	newTCT := &model.Ticket{
		TicketID: tId,
		Details:  details,
	}
	return database.UpdateTicketDetails(newTCT)
}
