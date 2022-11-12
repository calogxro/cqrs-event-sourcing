package domain

import "github.com/calogxro/cqrs-es/utils"

type UserContacts struct {
	contactsByType map[string]*utils.Set[Contact]
}

func NewUserContacts() *UserContacts {
	return &UserContacts{
		contactsByType: make(map[string]*utils.Set[Contact]),
	}
}

func (c *UserContacts) GetContactsByType(contactType string) *utils.Set[Contact] {
	return c.contactsByType[contactType]
}

func (c *UserContacts) AddContact(contact Contact) {
	if c.contactsByType[contact.contactType] == nil {
		c.contactsByType[contact.contactType] = utils.NewSet[Contact]()
	}
	c.contactsByType[contact.contactType].Add(contact)
}

func (c *UserContacts) RemoveContact(contact Contact) {
	if c.contactsByType[contact.contactType] != nil {
		c.contactsByType[contact.contactType].Remove(contact)
	}
}
