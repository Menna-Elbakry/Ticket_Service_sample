package admin

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetAllOperatorTickets(OpId uuid.UUID) {
	newTCT := &model.Ticket{
		OperatorID: OpId,
	}
	database.GetAllOperatorTickets(newTCT)
}
