package customer

import (
	"task/database"
)

//GetAllUserTickets to select all tickets exist
func GetAllUserTickets() {
	newTCT := &database.Tickets{
		UserID: ID,
	}
	newTCT.GetAllCustomerTickets()
}
