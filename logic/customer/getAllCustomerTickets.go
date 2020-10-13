package customer

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetAllCustomerTickets(usrId uuid.UUID) {
	newTCT := &model.Ticket{
		UserID: usrId,
	}
	database.GetAllCustomerTickets(newTCT)
}
