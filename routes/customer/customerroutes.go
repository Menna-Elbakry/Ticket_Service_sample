package custroute

import (
	"github.com/labstack/echo/v4"
)

func Customer() {

	e := echo.New()

	e.GET("/customer/getticket/:ticketId", GetTicketByID)
	e.GET("/customer/customertickets/:userId", GetAllCustomerTickets)

	e.PUT("/customer/updateticketdetails/:ticketId", UpdateTicketDetails)
	e.PUT("/customer/updateusername/:userId", UpdateUserName)
	e.PUT("/customer/updateuserpassword/:userId", UpdateUserPass)
	e.PUT("/customer/updateusermail/:userId", UpdateUserMail)

	e.DELETE("/customer/deleteticket/:ticketId", DeleteTicket)
	e.DELETE("/customer/deleteuser/:userId", DeleteUser)

	e.POST("/customer/adduser", AddUser)
	e.POST("/customer/addticket/:userId", AddTicket)

	e.Logger.Fatal(e.Start(":8080"))

}
