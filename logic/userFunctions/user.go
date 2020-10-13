package logic

import (
	"task/model"

	uuid "github.com/satori/go.uuid"
)

var (
	Pass  string = "operator"
	ID           = uuid.Must(uuid.FromString("06adcb61-510a-40c3-b678-25d5bb512aed"))
	Name  string = "operator"
	Email string = "operator@yahoo.com"
	Role         = model.Operator
)