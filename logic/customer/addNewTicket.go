package customer

import (
	"task/database"
	"time"

	uuid "github.com/satori/go.uuid"
)

//AddTicket to add new ticket to database
func AddTicket() {
	newTCT := &database.Tickets{
		TicketID:   uuid.NewV4(),
		UserID:     ID,
		UserName:   "",
		Details:    Details,
		Status:     "unsolved",
		OperatorID: uuid.Nil,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	newTCT.AddNewTicket()
}
