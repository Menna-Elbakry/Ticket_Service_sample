package operator

import (
	database "task/database/implementation"
	"task/model"

	uuid "github.com/satori/go.uuid"
)

func GetAllOperatorTickets(OpId uuid.UUID) ([]model.Ticket,error){
	newTCT := &model.Ticket{
		OperatorID: OpId,
	}
	return database.GetAllOperatorTickets(newTCT)
}
