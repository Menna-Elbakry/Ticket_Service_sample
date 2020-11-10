package customer

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewTicket(uId uuid.UUID, username string, details string) (uuid.UUID, error) {
	newTCT := &model.Ticket{
		TicketID:   uuid.NewV4(),
		UserID:     uId,
		UserName:   username,
		Details:    details,
		Status:     model.Unsolved,
		OperatorID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
	}
	return database.AddNewTicket(newTCT)

}
