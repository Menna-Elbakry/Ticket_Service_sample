package adroute

import (
	"github.com/labstack/echo/v4"
)

func Admin() {

	e := echo.New()

	e.GET("/admin/users", Users)
	e.GET("/admin/getuser/:userId", GetUserByID)
	e.GET("/admin/admins", Admins)
	e.GET("/admin/operators", Operators)
	e.GET("/admin/customers", Customers)
	e.GET("/admin/tickets", Tickets)
	e.GET("/admin/getticket/:ticketId", GetTicketByID)
	e.GET("/admin/usertickets/:userId", GetAllUserTickets)

	e.PUT("/admin/updateticketdetails/:ticketId", UpdateTicketDetails)
	e.PUT("/admin/updateusername/:userId", UpdateUserName)
	e.PUT("/admin/updateuserpassword/:userId", UpdateUserPass)
	e.PUT("/admin/updateusermail/:userId", UpdateUserMail)

	e.DELETE("/admin/deleteticket/:ticketId", DeleteTicket)
	e.DELETE("/admin/deleteuser/:userId", DeleteUser)

	e.POST("/admin/adduser", AddUser)
	e.POST("/admin/addticket/:userId", AddTicket)

	e.Logger.Fatal(e.Start(":8080"))

}
