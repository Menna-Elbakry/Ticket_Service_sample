package main

import "task/routes"

var (
	mail string = "customer@yahoo.com"
	pass string = "customer"
)

func main() {
	routes.Server()

	// http.HandleFunc("/admins", routes.Admins)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// userID, _ := logic.Login(mail, pass)

	// userRole, _ := admin.GetUserRoleById(userID)

	// switch userRole {
	// case model.Admin:
	// 	log.Println("Tickets")
	// 	admin.GetAllTickets()
	// 	log.Println("Users")
	// 	admin.GetAllUsers()

	// case model.Operator:
	// 	consume.Consume()
	// 	log.Println("Tickets")
	// 	operator.GetAllOperatorTickets(userID)

	// case model.Customer:
	// 	rand.Seed(time.Now().UTC().UnixNano())

	// 	tId, _ := customer.AddNewTicket()
	// 	msg, _ := customer.GetTicketById(tId)

	// 	log.Printf(" [x] sending (%v)", msg)
	// 	_, err := publish.Connect(tId)
	// 	publish.FailOnError(err, "Failed to handle Ticket request")

	// 	log.Println("Tickets")
	// 	customer.GetAllCustomerTickets(userID)

	// default:
	// 	log.Println("Error")
	// }

}
