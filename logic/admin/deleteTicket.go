package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func DeleteTicket(tId uuid.UUID) (uuid.UUID, error) {
	newTCT := &model.Ticket{
		TicketID: tId,
	}
	return database.DeleteTicket(newTCT)
}
