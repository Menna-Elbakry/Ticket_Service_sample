package customer

import (
	"task/model"

	uuid "github.com/satori/go.uuid"
)

var (
	Pass  string = "customer"
	ID           = uuid.Must(uuid.FromString("06adcb61-510a-40c3-b678-25d5bb512aed"))
	Name  string = "customer"
	Email string = "customer@yahoo.com"
	Role         = model.Customer

	TID               = uuid.Must(uuid.FromString("48e225eb-0772-48c0-b2d0-b86c555d4642"))
	UserID            = uuid.Must(uuid.FromString("74270210-e95c-4865-900e-d7ed9c256d5e"))
	UserName   string = "customer"
	Status            = model.Unsolved
	Details    string = "Hello"
	OperatorID        = uuid.Must(uuid.FromString("c6fbdc5e-ee50-4cde-aee5-2983d157b5ba"))
)
