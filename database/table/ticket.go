package table

import (
	"task/model"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Ticket table
type Ticket struct {
	tableName  struct{}         `sql:"tickets"`
	TicketID   uuid.UUID        `sql:"ticket_id,pk,type:uuid"`
	UserID     uuid.UUID        `sql:"user_id,type:uuid"`
	UserName   string           `sql:"user_name"`
	Details    string           `sql:"details"`
	Status     model.StatusEnum `sql:"status"`
	OperatorID uuid.UUID        `sql:"operator_id,type:uuid"`
	CreatedAt  time.Time        `sql:"created_at"`
	UpdatedAt  time.Time        `sql:"updated_at"`
}

func (t *Ticket) MapToModule() model.Ticket {
	return model.Ticket{
		TicketID:   t.TicketID,
		UserID:     t.UserID,
		UserName:   t.UserName,
		Details:    t.Details,
		Status:     t.Status,
		OperatorID: t.OperatorID,
	}
}

func (t Ticket) Fill(ticket *model.Ticket) *Ticket {
	return &Ticket{
		TicketID:   ticket.TicketID,
		UserID:     ticket.UserID,
		UserName:   ticket.UserName,
		Details:    ticket.Details,
		Status:     ticket.Status,
		OperatorID: ticket.OperatorID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
