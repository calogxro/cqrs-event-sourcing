package domain

import "github.com/calogxro/cqrs-es/utils"

type User struct {
	id        string
	firstName string
	lastName  string
	contacts  *utils.Set[Contact]
	addresses *utils.Set[Address]
}

func NewUser(userId, firstName, lastName string) *User {
	return &User{
		id:        userId,
		firstName: firstName,
		lastName:  lastName,
		contacts:  utils.NewSet[Contact](),
		addresses: utils.NewSet[Address](),
	}
}

func (user *User) GetId() string {
	return user.id
}

func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) GetContacts() *utils.Set[Contact] {
	return user.contacts
}

func (user *User) SetContacts(contacts *utils.Set[Contact]) {
	user.contacts = contacts
}

func (user *User) GetAddresses() *utils.Set[Address] {
	return user.addresses
}

func (user *User) SetAddresses(addresses *utils.Set[Address]) {
	user.addresses = addresses
}
