package database

import (
	"log"
	"task/database/table"
	"task/model"
	"time"

	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
)

//GetAllTickets to select all tickets exist
func GetAllTickets() ([]model.Ticket, error) {
	var ticketEntity []table.Ticket
	selErr := db.Model(&ticketEntity).Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return nil, selErr
	}
	var tickets []model.Ticket
	for _, tct := range ticketEntity {
		tickets = append(tickets, tct.MapToModule())
		log.Printf("Get Users Successful . \n UserName: %v ,Details: %v,Status: %v\n", tct.UserName, tct.Details, tct.Status)
	}
	return tickets, nil
}

//GetAllUserTicketsById to select specific user tickets
func GetAllUserTicketsById(ticket *model.Ticket) ([]model.Ticket, error) {
	var ticketEntity []table.Ticket
	selErr := db.Model(&ticketEntity).
		Where("user_id=?", ticket.UserID).
		Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return nil, selErr
	}
	var tickets []model.Ticket
	for _, tct := range ticketEntity {
		tickets = append(tickets, tct.MapToModule())
		log.Printf("Get Users Successful . \n ID: %v ,Details: %v,Status: %v", tct.TicketID, tct.Details, tct.Status)
	}
	return tickets, nil
}

//GetAllOperatorTickets to select specific user tickets
func GetAllOperatorTickets(ticket *model.Ticket) ([]model.Ticket, error) {
	var ticketEntity []table.Ticket
	selErr := db.Model(&ticketEntity).
		Where("operator_id=?", ticket.OperatorID).
		Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return nil, selErr
	}
	var tickets []model.Ticket
	for _, tct := range ticketEntity {
		tickets = append(tickets, tct.MapToModule())
		log.Printf("Get Users Successful . \n ID: %v ,Username: %v,Details: %v,Status: %v", tct.TicketID, tct.UserName, tct.Details, tct.Status)
	}
	return tickets, nil
}

//GetTicketById to search for user
func GetTicketById(ticket *model.Ticket) (string, error) {
	var ticketEntity table.Ticket
	getErr := db.Model(&ticketEntity).
		Where("ticket_id=?", ticket.TicketID).
		Returning("user_name,details").
		Select()
	if getErr != nil {
		log.Printf("Error While Getting Ticket By Id Reason:  %v\n", getErr)
		return " ", getErr
	}
	log.Printf("Get Ticket Successful  \n User: %v Details: %v ", ticketEntity.UserName, ticketEntity.Details)
	return ticket.Details, nil
}

//AddNewTicket to database
func AddNewTicket(ticket *model.Ticket) (uuid.UUID, error) {
	_, insertErr := db.Model(table.Ticket{}.Fill(ticket)).Insert()
	if insertErr != nil {
		log.Printf("Error While Adding New Ticket Reason:  %v\n", insertErr)
		return uuid.Nil, insertErr
	}
	log.Printf("New Ticket Added Successfully .\n Ticket ID: %v", ticket.TicketID)
	return ticket.TicketID, nil
}

//DeleteTicket to delete ticket
func DeleteTicket(ticket *model.Ticket) (uuid.UUID, error) {
	var ticketEntity table.Ticket
	_, deleteErr := db.Model(&ticketEntity).
		Where("ticket_id=?", ticket.TicketID).
		Returning("ticket_id").
		Delete()
	if deleteErr != nil {
		log.Printf("Error While Deleting Ticket. Reason:  %v\n", deleteErr)
		return uuid.Nil, deleteErr
	}
	log.Printf("Successfully Deleted Ticket With ID: %s\n", ticket.TicketID)
	return ticket.TicketID, nil
}

//UpdateTicketDetails to update the ticket content
func UpdateTicketDetails(ticket *model.Ticket) (string, error) {
	var ticketEntity table.Ticket
	_, updateErr := db.Model(&ticketEntity).
		Set("details=?,updated_at=?", ticket.Details, time.Now()).
		Returning("ticket_id,user_name,details").
		Where("ticket_id=?", ticket.TicketID).
		Update()
	if updateErr != nil {
		log.Printf("Error While Updating Ticket. Reason:  %v\n", updateErr)
		return " ", updateErr
	}
	log.Printf("Ticket Updated Successfully For Username: %v, TicketID: %v, Details: %v \n ", ticket.UserName, ticket.TicketID, ticket.Details)
	return ticket.Details, nil
}

//UpdateOperatorIDAndStatus to update the User Name and Password
func UpdateOperatorIDAndStatus(ticket *model.Ticket) error {
	var ticketEntity table.Ticket
	_, updateErr := db.Model(&ticketEntity).
		Set("operator_id=? ,status=?,updated_at=?", ticket.OperatorID, ticket.Status, time.Now()).
		Where("ticket_id=?", ticket.TicketID).
		Returning("ticket_id,status").
		Update()
	if updateErr != nil {
		log.Printf("Error While Updating Status Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Status Updated Successfully For Ticket %v, %v\n ", ticket.TicketID, ticket.Status)
	return nil
}

//CreateTicketTable to add tickets table to database
func CreateTicketTable() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&table.Ticket{}, opts)
	if createErr != nil {
		log.Printf("Error While Creating Table Tickets Reason:  %v\n", createErr)
		return createErr
	}
	log.Printf("Table Tickets Created Successfully.\n")
	return nil
}
