package es

import (
	"errors"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

type UserService struct {
	repository *EventStore
}

func NewUserService(repository *EventStore) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) createUser(userId, firstName, lastName string) {
	event := NewUserCreatedEvent(userId, firstName, lastName)
	s.repository.AddEvent(userId, event)
}

func (s *UserService) updateUser(userId string, contacts *utils.Set[d.Contact], addresses *utils.Set[d.Address]) error {
	user := RecreateUserState(s.repository, userId)

	// Contacts to remove
	for _, c := range user.GetContacts().Values() {
		if !contacts.Contains(c) {
			event := NewUserContactRemovedEvent(c.GetType(), c.GetDetail())
			s.repository.AddEvent(userId, event)
		}
	}

	// Contacts to add
	for _, c := range contacts.Values() {
		if !user.GetContacts().Contains(c) {
			event := NewUserContactAddedEvent(c.GetType(), c.GetDetail())
			s.repository.AddEvent(userId, event)
		}
	}

	// Addresses to remove
	for _, a := range user.GetAddresses().Values() {
		if !addresses.Contains(a) {
			event := NewUserAddressRemovedEvent(a.GetCity(), a.GetState())
			s.repository.AddEvent(userId, event)
		}
	}

	// Addresses to add
	for _, a := range addresses.Values() {
		if !user.GetAddresses().Contains(a) {
			event := NewUserAddressAddedEvent(a.GetCity(), a.GetState())
			s.repository.AddEvent(userId, event)
		}
	}
	return nil
}

func (s *UserService) getContactByType(userId string, contactType string) (*utils.Set[d.Contact], error) {
	user := RecreateUserState(s.repository, userId)
	if user == nil {
		return nil, errors.New("")
	}
	contacts := utils.NewSet[d.Contact]()
	for _, contact := range user.GetContacts().Values() {
		if contact.GetType() == contactType {
			contacts.Add(contact)
		}
	}
	return contacts, nil
}

func (s *UserService) getAddressByRegion(userId string, state string) (*utils.Set[d.Address], error) {
	user := RecreateUserState(s.repository, userId)
	if user == nil {
		return nil, errors.New("")
	}
	addresses := utils.NewSet[d.Address]()
	for _, address := range user.GetAddresses().Values() {
		if address.GetState() == state {
			addresses.Add(address)
		}
	}
	return addresses, nil
}
