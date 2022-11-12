package es

import d "github.com/calogxro/cqrs-es/domain"

func RecreateUserState(store *EventStore, userId string) *d.User {
	var user *d.User
	events := store.getEvents(userId)
	for _, iEvent := range events {
		if e, ok := iEvent.(UserCreatedEvent); ok {
			user = d.NewUser(e.userId, e.firstName, e.lastName)
		}
		if user != nil {
			if e, ok := iEvent.(UserContactAddedEvent); ok {
				contact := d.NewContact(e.ContactType, e.ContactDetails)
				user.GetContacts().Add(contact)
			}
			if e, ok := iEvent.(UserContactRemovedEvent); ok {
				contact := d.NewContact(e.ContactType, e.ContactDetails)
				user.GetContacts().Remove(contact)
			}
			if e, ok := iEvent.(UserAddressAddedEvent); ok {
				address := d.NewAddress(e.City, e.State)
				user.GetAddresses().Add(address)
			}
			if e, ok := iEvent.(UserAddressRemovedEvent); ok {
				address := d.NewAddress(e.City, e.State)
				user.GetAddresses().Remove(address)
			}
		}
	}
	return user
}
