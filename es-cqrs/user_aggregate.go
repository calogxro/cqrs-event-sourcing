package escqrs

import (
	"github.com/calogxro/cqrs-es/cqrs"
	"github.com/calogxro/cqrs-es/es"
)

type UserAggregate struct {
	writeRepository *es.EventStore
}

func NewUserAggregate(repository *es.EventStore) *UserAggregate {
	return &UserAggregate{
		writeRepository: repository,
	}
}

// Create

func (s *UserAggregate) handleCreateUserCommand(cmd cqrs.CreateUserCommand) []es.IEvent {
	event := es.NewUserCreatedEvent(cmd.UserId, cmd.FirstName, cmd.LastName)
	s.writeRepository.AddEvent(cmd.UserId, event)
	return []es.IEvent{event}
}

// Update

func (s *UserAggregate) handleUpdateUserCommand(cmd cqrs.UpdateUserCommand) ([]es.IEvent, error) {
	user := es.RecreateUserState(s.writeRepository, cmd.UserId)
	events := []es.IEvent{}

	// Contacts to remove
	for _, c := range user.GetContacts().Values() {
		if !cmd.Contacts.Contains(c) {
			event := es.NewUserContactRemovedEvent(c.GetType(), c.GetDetail())
			events = append(events, event)
			s.writeRepository.AddEvent(cmd.UserId, event)
		}
	}

	// Contacts to add
	for _, c := range cmd.Contacts.Values() {
		if !user.GetContacts().Contains(c) {
			event := es.NewUserContactAddedEvent(c.GetType(), c.GetDetail())
			events = append(events, event)
			s.writeRepository.AddEvent(cmd.UserId, event)
		}
	}

	// Addresses to remove
	for _, a := range user.GetAddresses().Values() {
		if !cmd.Addresses.Contains(a) {
			event := es.NewUserAddressRemovedEvent(a.GetCity(), a.GetState())
			events = append(events, event)
			s.writeRepository.AddEvent(cmd.UserId, event)
		}
	}

	// Addresses to add
	for _, a := range cmd.Addresses.Values() {
		if !user.GetAddresses().Contains(a) {
			event := es.NewUserAddressAddedEvent(a.GetCity(), a.GetState())
			events = append(events, event)
			s.writeRepository.AddEvent(cmd.UserId, event)
		}
	}
	return events, nil
}
