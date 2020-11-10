package operator

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateOperatorIDAndStatus(tId uuid.UUID, opId uuid.UUID, stat model.StatusEnum) error{
	newTCT := &model.Ticket{
		TicketID:   tId,
		OperatorID: opId,
		Status:     stat,
	}
	return database.UpdateOperatorIDAndStatus(newTCT)
}
