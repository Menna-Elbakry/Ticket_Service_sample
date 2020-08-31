package operator

import (
	"task/database"
)

//GetAllUserTickets to select all tickets exist
func GetAllUserTickets() {
	newTCT := &database.Tickets{
		OperatorID: OpID,
	}
	newTCT.GetAllOperatorTickets()
}
