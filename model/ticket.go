package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type StatusEnum string

const (
	Solved     StatusEnum = "SOLVED"
	Unsolved   StatusEnum = "UNSOLVED"
	Processing StatusEnum = "Processing"
)

//Ticket table
type Ticket struct {
	TicketID   uuid.UUID  `json:"ticket_id,pk,type:uuid`
	UserID     uuid.UUID  `json:"user_iD"`
	UserName   string     `json:"user_name"`
	Details    string     `json:"details"`
	Status     StatusEnum `json:"status"`
	OperatorID uuid.UUID  `json:"operator_id"`
}

func (t *Ticket) Validate() error {
	if t.TicketID == uuid.Nil {
		return errors.New("should have TicketID")
	}
	if t.UserID == uuid.Nil {
		return errors.New("should have userId")
	}
	if t.UserName == " " {
		return errors.New("should have userName")
	}
	if t.Details == " " {
		return errors.New("should have Details")
	}
	if t.Status == " " {
		return errors.New("should have status")
	}
	if t.OperatorID == uuid.Nil {
		return errors.New("should have operatorId")
	}
	return nil
}
