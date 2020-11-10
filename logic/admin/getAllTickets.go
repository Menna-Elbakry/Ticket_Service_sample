package admin

import (
	database "task/database/implementation"
	"task/model"
)

func GetAllTickets() ([]model.Ticket, error) {
	return database.GetAllTickets()
}
