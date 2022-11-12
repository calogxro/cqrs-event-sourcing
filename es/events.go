package es

import (
	"time"

	"github.com/google/uuid"
)

type IEvent interface {
}

type Event struct {
	id      uuid.UUID
	created time.Time
}

func NewEvent() Event {
	return Event{
		id:      uuid.New(),
		created: time.Now(),
	}
}

type UserCreatedEvent struct {
	Event
	userId    string
	firstName string
	lastName  string
}

func NewUserCreatedEvent(userId, firstName, lastName string) UserCreatedEvent {
	return UserCreatedEvent{
		Event:     NewEvent(),
		userId:    userId,
		firstName: firstName,
		lastName:  lastName,
	}
}

type UserContactAddedEvent struct {
	Event
	ContactType    string
	ContactDetails string
}

func NewUserContactAddedEvent(contactType, contactDetails string) UserContactAddedEvent {
	return UserContactAddedEvent{
		Event:          NewEvent(),
		ContactType:    contactType,
		ContactDetails: contactDetails,
	}
}

type UserContactRemovedEvent struct {
	Event
	ContactType    string
	ContactDetails string
}

func NewUserContactRemovedEvent(contactType, contactDetails string) UserContactRemovedEvent {
	return UserContactRemovedEvent{
		Event:          NewEvent(),
		ContactType:    contactType,
		ContactDetails: contactDetails,
	}
}

type UserAddressAddedEvent struct {
	Event
	City  string
	State string
	//postCode string
}

func NewUserAddressAddedEvent(city, state string) UserAddressAddedEvent {
	return UserAddressAddedEvent{
		Event: NewEvent(),
		City:  city,
		State: state,
	}
}

type UserAddressRemovedEvent struct {
	Event
	City  string
	State string
	//postCode string
}

func NewUserAddressRemovedEvent(city, state string) UserAddressRemovedEvent {
	return UserAddressRemovedEvent{
		Event: NewEvent(),
		City:  city,
		State: state,
	}
}
