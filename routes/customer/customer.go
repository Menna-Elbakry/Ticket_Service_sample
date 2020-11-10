package custroute

import (
	"log"
	"math/rand"
	"net/http"
	"task/logic/admin"
	"task/logic/customer"
	"task/queue/publish"
	"time"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

//GetAllCustomerTickets find ticket by it's id
func GetAllCustomerTickets(c echo.Context) error {
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))
	res, err := customer.GetAllCustomerTickets(uID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//GetTicketByID find ticket by it's id
func GetTicketByID(c echo.Context) error {
	id := c.Param("ticketId")
	tID := uuid.Must(uuid.FromString(id))
	res, err := customer.GetTicketById(tID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateTicketDetails find ticket by it's id
func UpdateTicketDetails(c echo.Context) error {

	id := c.Param("ticketId")
	tID := uuid.Must(uuid.FromString(id))

	details := c.QueryParam("details")

	res, err := customer.UpdateTicketDetails(tID, details)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserName find ticket by it's id
func UpdateUserName(c echo.Context) error {

	id := c.Param("userId")
	uId := uuid.Must(uuid.FromString(id))

	name := c.QueryParam("name")

	res, err := customer.UpdateUserName(uId, name)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserPass find ticket by it's id
func UpdateUserPass(c echo.Context) error {

	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))

	pass := c.QueryParam("password")

	res, err := customer.UpdateUserPass(uID, pass)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserMail find ticket by it's id
func UpdateUserMail(c echo.Context) error {

	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))

	mail := c.QueryParam("email")

	res, err := customer.UpdateUserMail(uID, mail)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//DeleteTicket ticket by it's id
func DeleteTicket(c echo.Context) error {

	id := c.Param("ticketId")
	tID := uuid.Must(uuid.FromString(id))

	res, err := customer.DeleteTicket(tID)
	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//DeleteUser user by it's id
func DeleteUser(c echo.Context) error {
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))

	res, err := customer.DeleteUser(uID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//AddUser create new user
func AddUser(c echo.Context) error {

	name := c.QueryParam("name")
	pass := c.QueryParam("pass")
	mail := c.QueryParam("mail")

	res, err := customer.AddNewUser(name, pass, mail)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//AddTicket create new ticket
func AddTicket(c echo.Context) error {
	rand.Seed(time.Now().UTC().UnixNano())

	details := c.QueryParam("details")
	id := c.Param("userId")
	uId := uuid.Must(uuid.FromString(id))

	username, _ := admin.GetUserById(uId)

	tId, Err := customer.AddNewTicket(uId, username, details)
	if Err != nil {
		log.Println("Unsupported method")
	}
	msg, _ := customer.GetTicketById(tId)

	log.Printf(" [x] sending (%v)", msg)
	_, err := publish.Connect(tId)
	publish.FailOnError(err, "Failed to handle Ticket request")

	return c.JSON(http.StatusCreated, tId)
}

//_________________________________________________________//
