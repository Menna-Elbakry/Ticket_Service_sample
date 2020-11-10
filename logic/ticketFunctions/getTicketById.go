package ticket

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetTicketById(TId uuid.UUID) (string, error) {
	newTCT := &model.Ticket{
		TicketID: TId,
	}
	return database.GetTicketById(newTCT)
}
