package adroute

import (
	"log"
	"net/http"
	"task/logic/admin"
	"task/model"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// func Admins(w http.ResponseWriter, r *http.Request) {

// 	resp, err := admin.GetUsersByRole(model.Admin)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "Unsupported method %v to %v\n", r.Method, r.URL)
// 		log.Printf("Unsupported method %v to %v\n", r.Method, r.URL)
// 	}

// 	fmt.Fprint(w, resp)

// }

//Users find all users
func Users(c echo.Context) error {

	res, err := admin.GetAllUsers()

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}

//_________________________________________________________//

//GetUserByID find user by it's id
func GetUserByID(c echo.Context) error {
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))
	res, err := admin.GetUserById(uID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//
//GetUserByRole find user by it's id
func GetUserByRole(c echo.Context) error {
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))
	res, err := admin.GetAllUserTickets(uID)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//Admins find all admins
func Admins(c echo.Context) error {

	res, err := admin.GetUsersByRole(model.Admin)

	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)

}

//_________________________________________________________//

//Operators find all operators
func Operators(c echo.Context) error {

	res, err := admin.GetUsersByRole(model.Operator)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}

//_________________________________________________________//

//Customers find all customers
func Customers(c echo.Context) error {

	res, err := admin.GetUsersByRole(model.Customer)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}

//_________________________________________________________//

//Tickets find all Tickets
func Tickets(c echo.Context) error {

	res, err := admin.GetAllTickets()

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)

}

//_________________________________________________________//

//GetAllUserTicketsfind ticket by it's id
func GetAllUserTickets(c echo.Context) error {
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))
	res, err := admin.GetAllUserTickets(uID)

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
	res, err := admin.GetTicketById(tID)

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

	res, err := admin.UpdateTicketDetails(tID, details)
	if err != nil {
		log.Println("Unsupported method")
	}

	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//UpdateUserName find ticket by it's id
func UpdateUserName(c echo.Context) error {

	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))

	name := c.QueryParam("name")

	res, err := admin.UpdateUserName(uID, name)
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

	res, err := admin.UpdateUserPass(uID, pass)
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

	res, err := admin.UpdateUserMail(uID, mail)
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

	res, err := admin.DeleteTicket(tID)
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

	res, err := admin.DeleteUser(uID)

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
	role := c.QueryParam("role")
	var Role model.RoleEnum
	switch role {
	case "admin":
		Role = model.Admin
	case "operator":
		Role = model.Operator
	case "customer":
		Role = model.Customer

	}
	res, err := admin.AddNewUser(name, pass, mail, Role)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//

//AddTicket create new ticket
func AddTicket(c echo.Context) error {

	// username := c.QueryParam("username")
	details := c.QueryParam("details")
	status := c.QueryParam("status")
	var Status model.StatusEnum
	switch status {
	case "processing":
		Status = model.Processing
	case "solved":
		Status = model.Solved
	case "unsolved":
		Status = model.Unsolved
	default:
		Status = model.Unsolved

	}
	id := c.Param("userId")
	uID := uuid.Must(uuid.FromString(id))

	username, err := admin.GetUserById(uID)

	res, err := admin.AddNewTicket(uID, username, details, Status)

	if err != nil {
		log.Println("Unsupported method")
	}
	return c.JSON(http.StatusCreated, res)
}

//_________________________________________________________//
