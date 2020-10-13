package main

import (
	"log"
	"math/rand"
	"task/logic"
	"task/logic/admin"
	"task/logic/customer"
	"task/logic/operator"
	"task/model"
	"task/queue/consume"
	"task/queue/publish"
	"time"
)

var (
	mail string = "customer@yahoo.com"
	pass string = "customer"
)

func main() {

	userID, _ := logic.Login(mail, pass)

	userRole, _ := admin.GetUserRoleById(userID)

	switch userRole {
	case model.Admin:
		log.Println("Tickets")
		admin.GetAllTickets()
		log.Println("Users")
		admin.GetAllUsers()

	case model.Operator:
		consume.Consume()
		log.Println("Tickets")
		operator.GetAllOperatorTickets(userID)

	case model.Customer:
		rand.Seed(time.Now().UTC().UnixNano())

		tId, _ := customer.AddNewTicket()
		msg, _ := customer.GetTicketById(tId)

		log.Printf(" [x] sending (%v)", msg)
		_, err := publish.Connect(tId)
		publish.FailOnError(err, "Failed to handle Ticket request")

		log.Println("Tickets")
		customer.GetAllCustomerTickets(userID)

	default:
		log.Println("Error")
	}

	//logic.Login()
	//logic.AddNewUser()
	//logic.GetUserById()
	//logic.GetAllUsers()
	//logic.UpdateUserPass()
	//logic.GetUsersByRole()
	//logic.UpdateUserName()
	//logic.UpdateUserMail()
	//logic.DeleteUser()

	//ticket.AddNewTicket()
	//ticket.GetAllTickets()
	//ticket.GetTicketById()
	//ticket.GetAllCustomerTickets()
	//ticket.GetAllOperatorTickets()
	//ticket.UpdateOperatorIDAndStatus()
	//ticket.UpdateTicketDetails()
	//ticket.DeleteTicket()

}
