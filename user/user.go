package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type User struct {
	lastName  string
	firstName string
	birthDate string
	createdAt time.Time
}

func (u *User) PrintUserFullname() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" && lastName == "" {
		return nil, errors.New("enter firstname or lastname")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// this is similar to inheritence
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin() *Admin {
	return &Admin{
		email:    randomdata.Email(),
		password: randomdata.MacAddress(),
		User: User{
			firstName: randomdata.FirstName(1),
			lastName:  randomdata.LastName(),
			birthDate: randomdata.FullDate(),
			createdAt: time.Now(),
		},
	}
}

func (admin *Admin) GetEmail_Mac() (string, string) {
	return admin.email, admin.password
}
