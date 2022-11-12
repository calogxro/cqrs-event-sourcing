package escqrs

import (
	"github.com/calogxro/cqrs-es/cqrs"
	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/es"
)

type Projector struct {
	repository *cqrs.ReadRepository
}

func NewProjector(repository *cqrs.ReadRepository) *Projector {
	return &Projector{
		repository: repository,
	}
}

func (p *Projector) project(userId string, events []es.IEvent) {
	for _, iEvent := range events {
		if e, ok := iEvent.(es.UserContactAddedEvent); ok {
			p.applyUserContactAddedEvent(userId, e)
		}
		if e, ok := iEvent.(es.UserContactRemovedEvent); ok {
			p.applyUserContactRemovedEvent(userId, e)
		}
		if e, ok := iEvent.(es.UserAddressAddedEvent); ok {
			p.applyUserAddressAddedEvent(userId, e)
		}
		if e, ok := iEvent.(es.UserAddressRemovedEvent); ok {
			p.applyUserAddressRemovedEvent(userId, e)
		}
	}
}

func (p *Projector) applyUserContactAddedEvent(userId string, e es.UserContactAddedEvent) {
	contact := d.NewContact(e.ContactType, e.ContactDetails)
	userContacts := p.repository.GetUserContacts(userId)
	if userContacts == nil {
		userContacts = d.NewUserContacts()
	}
	userContacts.AddContact(contact)
	p.repository.SetUserContacts(userId, userContacts)
}

func (p *Projector) applyUserContactRemovedEvent(userId string, e es.UserContactRemovedEvent) {
	contact := d.NewContact(e.ContactType, e.ContactDetails)
	userContacts := p.repository.GetUserContacts(userId)
	if userContacts != nil {
		userContacts.RemoveContact(contact)
		p.repository.SetUserContacts(userId, userContacts)
	}
}

func (p *Projector) applyUserAddressAddedEvent(userId string, e es.UserAddressAddedEvent) {
	address := d.NewAddress(e.City, e.State)
	userAddresses := p.repository.GetUserAddresses(userId)
	if userAddresses == nil {
		userAddresses = d.NewUserAddresses()
	}
	userAddresses.AddAddress(address)
	p.repository.SetUserAddresses(userId, userAddresses)
}

func (p *Projector) applyUserAddressRemovedEvent(userId string, e es.UserAddressRemovedEvent) {
	address := d.NewAddress(e.City, e.State)
	userAddresses := p.repository.GetUserAddresses(userId)
	if userAddresses != nil {
		userAddresses.RemoveAddress(address)
		p.repository.SetUserAddresses(userId, userAddresses)
	}
}
