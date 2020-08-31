package database

import (
	"log"

	"time"

	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//Users table
type Users struct {
	tableName struct{}  `sql:"users"`
	ID        uuid.UUID `sql:"id,pk,type:uuid"`
	Name      string    `sql:"name"`
	Email     string    `sql:"email,unique"`
	Password  string    `sql:"password"`
	Role      string    `sql:"role"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

//GetAllUsers to select all users exist
func (usr *Users) GetAllUsers() ([]Users, error) {
	var users []Users
	selErr := db.Model(&users).Select()
	if selErr != nil {
		log.Printf("Error While Getting Users Reason:  %v\n", selErr)
		return nil, selErr
	}
	for _, usr := range users {
		log.Printf("Get Users Successful .ID: %v, Name: %v", usr.ID, usr.Name)
	}
	return users, nil
}

//GetUserByID to search for user
func (usr *Users) GetUserByID() error {
	getErr := db.Model(usr).Where("id=?id").Select()
	if getErr != nil {
		log.Printf("Error While Getting User By Id Reason:  %v\n", getErr)
		return getErr
	}
	log.Printf("Get User By Id Successful For .%v\n ", usr.Name)
	return nil
}

//GetUserByRole to search for user
func (usr *Users) GetUserByRole(role string) ([]Users, error) {
	var users []Users
	getErr := db.Model(&users).Where("role=?", role).Select()
	if getErr != nil {
		log.Printf("Error While Getting Users Reason:  %v\n", getErr)
		return nil, getErr
	}
	for _, usr := range users {
		log.Printf("Get Users Successful .ID: %v, Name: %v", usr.ID, usr.Name)
	}
	return users, nil
}

//AddNewUser to database
func (usr *Users) AddNewUser() error {
	insertErr := db.Insert(usr)
	if insertErr != nil {
		log.Printf("Error While Adding New User Reason:  %v\n", insertErr)
		return insertErr
	}
	log.Printf("New User Added Successfully .\n User Id: %s", usr.ID)
	return nil
}

//DeleteUser to delete user
func (usr *Users) DeleteUser() error {
	_, deleteErr := db.Model(usr).Where("id=?id").Delete()
	if deleteErr != nil {
		log.Printf("Error While Deleting User. Reason:  %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Successfully Deleted User With ID: %s\n", usr.ID)
	return nil
}

//UpdateUserPassword to update the password
func (usr *Users) UpdateUserPassword() error {
	_, updateErr := db.Model(usr).Set("password=?password").Where("id=?id").Update()
	if updateErr != nil {
		log.Printf("Error While Updating Password. Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Password Updated Successfully For Email %v\n ", usr.Name)
	return nil
}

//UpdateUsername to modify the user name
func (usr *Users) UpdateUsername() error {
	_, updateErr := db.Model(usr).Set("Name=?Name").Where("id=?id").Update()
	if updateErr != nil {
		log.Printf("Error While Updating Name. Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Name Updated Successfully For Email %v\n ", usr.Email)
	return nil
}

//UpdateUserMail to modify user email
func (usr *Users) UpdateUserMail() error {
	_, updateErr := db.Model(usr).Set("Email=?Mail").Where("id=?id").Update()
	if updateErr != nil {
		log.Printf("Error While Updating Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Email Updated Successfully For %v\n ", usr.Name)
	return nil
}

//Login func
func (usr *Users) Login(mail, pass string) error {
	getErr := db.Model(usr).Where("Email=?", mail).Select()
	if getErr != nil {
		log.Printf("Error While Login :  %v\n", getErr)
		return getErr
	}
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil
	}
	log.Printf("Successfully Loged ID .%v\n ", usr.ID)

	return nil
}

//CreateUserTable to add users table to database
func CreateUserTable() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Users{}, opts)
	if createErr != nil {
		log.Printf("Error While Creating Table Users Reason:  %v\n", createErr)
		return createErr
	}
	log.Printf("Table Users Created Successfully.\n")
	return nil
}
