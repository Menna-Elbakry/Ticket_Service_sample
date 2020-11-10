package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewTicket(userId uuid.UUID, username string, details string, status model.StatusEnum) (uuid.UUID, error) {
	newTCT := &model.Ticket{
		TicketID:   uuid.NewV4(),
		UserID:     userId,
		UserName:   username,
		Details:    details,
		Status:     status,
		OperatorID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
	}
	return database.AddNewTicket(newTCT)

}
