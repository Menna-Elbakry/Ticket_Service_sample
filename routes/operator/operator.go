package optroute

import (
	"log"
	"net/http"
	"task/logic/operator"
	"task/queue/consume"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

//Listen
func Listen(c echo.Context) error {
	id := c.Param("opId")
	opId := uuid.Must(uuid.FromString(id))
	consume.Consume(opId)
	return c.JSON(http.StatusCreated, opId)
}

//_________________________________________________________//

//GetAllOperatorTickets find tickets by it's id
func GetAllOperatorTickets(c echo.Context) error {
	id := c.Param("opId")
	opID := uuid.Must(uuid.FromString(id))
	res, err := operator.GetAllOperatorTickets(opID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//GetTicketByID find ticket by it's id
func GetTicketByID(c echo.Context) error {
	id := c.Param("ticketId")
	tID := uuid.Must(uuid.FromString(id))
	res, err := operator.GetTicketById(tID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserName change username
func UpdateUserName(c echo.Context) error {

	id := c.Param("opId")
	opId := uuid.Must(uuid.FromString(id))

	name := c.QueryParam("name")

	res, err := operator.UpdateUserName(opId, name)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserPass change password
func UpdateUserPass(c echo.Context) error {

	id := c.Param("userId")
	opID := uuid.Must(uuid.FromString(id))

	pass := c.QueryParam("password")

	res, err := operator.UpdateUserPass(opID, pass)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}
