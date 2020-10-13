package ticket

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func UpdateOperatorIDAndStatus(tId uuid.UUID, opId uuid.UUID, stat model.StatusEnum) {
	newTCT := &model.Ticket{
		TicketID:   tId,
		OperatorID: opId,
		Status:     stat,
	}
	database.UpdateOperatorIDAndStatus(newTCT)
}
