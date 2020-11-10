package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetAllUserTickets(usrId uuid.UUID) ([]model.Ticket, error) {
	newTCT := &model.Ticket{
		UserID: usrId,
	}
	return database.GetAllUserTicketsById(newTCT)
}
