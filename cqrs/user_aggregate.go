package cqrs

import (
	"errors"

	d "github.com/calogxro/cqrs-es/domain"
)

type UserAggregate struct {
	repository *WriteRepository
}

func NewUserAggregate(repository *WriteRepository) *UserAggregate {
	return &UserAggregate{
		repository: repository,
	}
}

// Create

func (s *UserAggregate) handleCreateUserCommand(command CreateUserCommand) *d.User {
	user := d.NewUser(command.UserId, command.FirstName, command.LastName)
	s.repository.addUser(command.UserId, user)
	return user
}

// Update

func (s *UserAggregate) handleUpdateUserCommand(command UpdateUserCommand) (*d.User, error) {
	user := s.repository.getUser(command.UserId)
	if user == nil {
		return nil, errors.New("")
	}
	user.SetContacts(command.Contacts)
	user.SetAddresses(command.Addresses)
	//s.repository.updateUser(command.userId, user)
	return user, nil
}
