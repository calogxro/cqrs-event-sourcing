package cqrs

import (
	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

type CreateUserCommand struct {
	UserId    string
	FirstName string
	LastName  string
}

type UpdateUserCommand struct {
	UserId    string
	Contacts  *utils.Set[d.Contact]
	Addresses *utils.Set[d.Address]
}
