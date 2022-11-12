package cqrs

import (
	d "github.com/calogxro/cqrs-es/domain"
)

type Projector struct {
	repository *ReadRepository
}

func NewProjector(repository *ReadRepository) *Projector {
	return &Projector{
		repository: repository,
	}
}

func (p *Projector) project(user *d.User) {
	// userContacts
	userContacts := d.NewUserContacts()
	if user.GetContacts() != nil {
		for _, contact := range user.GetContacts().Values() {
			userContacts.AddContact(contact)
		}
	}
	p.repository.SetUserContacts(user.GetId(), userContacts)

	// userAddresses
	userAddresses := d.NewUserAddresses()
	if user.GetAddresses() != nil {
		for _, address := range user.GetAddresses().Values() {
			userAddresses.AddAddress(address)
		}
	}
	p.repository.SetUserAddresses(user.GetId(), userAddresses)
}
