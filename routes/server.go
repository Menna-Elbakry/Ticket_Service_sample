package routes

import (
	"log"
	"net/http"
	"task/logic"
	"task/logic/admin"
	"task/model"
	adroute "task/routes/admin"
	custroute "task/routes/customer"
	optroute "task/routes/operator"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	Mail := c.QueryParam("mail")
	Pass := c.QueryParam("pass")

	userId, err := logic.Login(Mail, Pass)

	if err != nil {
		log.Println("Unsupported method")
	}
	role, _ := admin.GetUserRoleById(userId)

	switch role {
	case model.Admin:
		adroute.Admin()
	case model.Operator:
		optroute.Operator()
	case model.Customer:
		custroute.Customer()

	}
	return c.JSON(http.StatusCreated, userId)

}

//Server to connect to http
func Server() {

	e := echo.New()

	e.GET("/login", Login)

	//Admin Route
	e.GET("/admin/users", adroute.Users)
	e.GET("/admin/getuser/:userId", adroute.GetUserByID)
	e.GET("/admin/admins", adroute.Admins)
	e.GET("/admin/operators", adroute.Operators)
	e.GET("/admin/customers", adroute.Customers)
	e.GET("/admin/tickets", adroute.Tickets)
	e.GET("/admin/getticket/:ticketId", adroute.GetTicketByID)
	e.GET("/admin/usertickets/:userId", adroute.GetAllUserTickets)

	e.PUT("/admin/updateticketdetails/:ticketId", adroute.UpdateTicketDetails)
	e.PUT("/admin/updateusername/:userId", adroute.UpdateUserName)
	e.PUT("/admin/updateuserpassword/:userId", adroute.UpdateUserPass)
	e.PUT("/admin/updateusermail/:userId", adroute.UpdateUserMail)

	e.DELETE("/admin/deleteticket/:ticketId", adroute.DeleteTicket)
	e.DELETE("/admin/deleteuser/:userId", adroute.DeleteUser)

	e.POST("/admin/adduser", adroute.AddUser)
	e.POST("/admin/addticket/:userId", adroute.AddTicket)

	//Operator Route
	e.GET("/operator/consume/:opId", optroute.Listen)
	e.GET("/operator/getticket/:ticketId", optroute.GetTicketByID)
	e.GET("/operator/operatortickets/:userId", optroute.GetAllOperatorTickets)

	e.PUT("/operator/updateusername/:userId", optroute.UpdateUserName)
	e.PUT("/operator/updateuserpassword/:userId", optroute.UpdateUserPass)

	//Customer Route
	e.GET("/customer/getticket/:ticketId", custroute.GetTicketByID)
	e.GET("/customer/customertickets/:userId", custroute.GetAllCustomerTickets)

	e.PUT("/customer/updateticketdetails/:ticketId", custroute.UpdateTicketDetails)
	e.PUT("/customer/updateusername/:userId", custroute.UpdateUserName)
	e.PUT("/customer/updateuserpassword/:userId", custroute.UpdateUserPass)
	e.PUT("/customer/updateusermail/:userId", custroute.UpdateUserMail)

	e.DELETE("/customer/deleteticket/:ticketId", custroute.DeleteTicket)
	e.DELETE("/customer/deleteuser/:userId", custroute.DeleteUser)

	e.POST("/customer/adduser", custroute.AddUser)
	e.POST("/customer/addticket/:userId", custroute.AddTicket)

	e.Logger.Fatal(e.Start(":8080"))

}
