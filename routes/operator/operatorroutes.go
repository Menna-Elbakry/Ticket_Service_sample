package optroute

import (
	"github.com/labstack/echo/v4"
)

func Operator() {

	e := echo.New()

	e.GET("/operator/consume/:opId", Listen)

	e.GET("/operator/getticket/:ticketId", GetTicketByID)
	e.GET("/operator/operatortickets/:userId", GetAllOperatorTickets)
	e.PUT("/operator/updateusername/:userId", UpdateUserName)
	e.PUT("/operator/updateuserpassword/:userId", UpdateUserPass)

	e.Logger.Fatal(e.Start(":8080"))

}
