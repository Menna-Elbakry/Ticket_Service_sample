package database

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
)

//Tickets table
type Tickets struct {
	tableName  struct{}  `sql:"tickets"`
	TicketID   uuid.UUID `sql:"ticket_id,pk,type:uuid"`
	UserID     uuid.UUID `sql:"user_id,type:uuid"`
	UserName   string    `sql:"user_name"`
	Details    string    `sql:"details"`
	Status     string    `sql:"status"`
	OperatorID uuid.UUID `sql:"operator_id,type:uuid"`
	CreatedAt  time.Time `sql:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at"`
}

//GetAllTickets to select all tickets exist
func (tct *Tickets) GetAllTickets() ([]Tickets, error) {
	var tickets []Tickets
	selErr := db.Model(tct).Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return nil, selErr
	}
	for _, tct := range tickets {
		log.Printf("Get Tickets Successful \n User: %v, Details: %v, Status: %v \n ", tct.UserName, tct.Details, tct.Status)
	}
	return tickets, nil
}

//GetAllCustomerTickets to select specific user tickets
func (tct *Tickets) GetAllCustomerTickets() error {
	var tickets []Tickets
	selErr := db.Model(tct).Where("user_id=?user_id").Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return selErr
	}
	for _, tct := range tickets {
		log.Printf("Get Tickets Successful .%v\n ", tct.UserName)
	}
	return nil
}

//GetAllOperatorTickets to select specific user tickets
func (tct *Tickets) GetAllOperatorTickets() error {
	var tickets []Tickets
	selErr := db.Model(tct).Where("operator_id=?user_id").Select()
	if selErr != nil {
		log.Printf("Error While Getting Tickets Reason:  %v\n", selErr)
		return selErr
	}
	for _, tct := range tickets {
		log.Printf("Get Tickets Successful \n User: %v, Details: %v \n ", tct.UserName, tct.Details)
	}
	return nil
}

//GetTicketByID to search for ticket
func (tct *Tickets) GetTicketByID() error {
	getErr := db.Model(tct).Where("ticket_id=?id").Select()
	if getErr != nil {
		log.Printf("Error While Getting Ticket By Id Reason:  %v\n", getErr)
		return getErr
	}
	log.Printf("Get Ticket By Id Successful For .%v\n ", tct.TicketID)
	return nil
}

//AddNewTicket to database
func (tct *Tickets) AddNewTicket() error {
	insertErr := db.Insert(tct)
	if insertErr != nil {
		log.Printf("Error While Adding New Ticket Reason:  %v\n", insertErr)
		return insertErr
	}
	log.Printf("New Ticket Added Successfully .\n Ticket ID: %s", tct.TicketID)
	return nil
}

//DeleteTicketByID to delete ticket
func (tct *Tickets) DeleteTicketByID() error {
	_, deleteErr := db.Model(tct).Where("id=?id").Delete()
	if deleteErr != nil {
		log.Printf("Error While Deleting Ticket. Reason:  %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Successfully Deleted Ticket With ID: %s\n", tct.TicketID)
	return nil
}

//UpdateTicketDetails to update the ticket content
func (tct *Tickets) UpdateTicketDetails() error {
	_, updateErr := db.Model(tct).Set("details=?details,updated_at=?0", time.Now()).Where("ticket_id=?id").Update()
	if updateErr != nil {
		log.Printf("Error While Updating Ticket. Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Ticket Updated Successfully For User %s\n ", tct.UserName)
	return nil
}

//UpdateOperatorIDAndStatus to update the User Name and Password
func (tct *Tickets) UpdateOperatorIDAndStatus() error {
	_, updateErr := db.Model(tct).Set("operator_id=?operator_id ,status='Solved',updated_at=?0", time.Now()).Where("ticket_id=?ticket_id").Update()
	if updateErr != nil {
		log.Printf("Error While Updating Status Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Status Updated Successfully For Ticket %s\n ", tct.TicketID)
	return nil
}

//CreateTicketTable to add tickets tabe to database
func CreateTicketTable() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Tickets{}, opts)
	if createErr != nil {
		log.Printf("Error While Creating Table Tickets Reason:  %v\n", createErr)
		return createErr
	}
	log.Printf("Table Tickets Created Successfully.\n")
	return nil
}
