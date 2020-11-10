package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func AddNewTicket() (uuid.UUID, error) {
	newTCT := &model.Ticket{
		TicketID:   uuid.NewV4(),
		UserID:     UserID,
		UserName:   UserName,
		Details:    Details,
		Status:     Status,
		OperatorID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
	}
	return database.AddNewTicket(newTCT)

}
