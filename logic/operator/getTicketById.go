package operator

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetTicketById(tId uuid.UUID) (string, error) {
	newTCT := &model.Ticket{
		TicketID: tId,
	}
	return database.GetTicketById(newTCT)
}
