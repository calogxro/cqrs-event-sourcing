package cqrs

import d "github.com/calogxro/cqrs-es/domain"

type ReadRepository struct {
	userContacts  map[string]*d.UserContacts
	userAddresses map[string]*d.UserAddresses
}

func NewReadRepository() *ReadRepository {
	return &ReadRepository{
		userContacts:  make(map[string]*d.UserContacts),
		userAddresses: make(map[string]*d.UserAddresses),
	}
}

// userContacts

func (r *ReadRepository) GetUserContacts(userId string) *d.UserContacts {
	return r.userContacts[userId]
}

func (r *ReadRepository) SetUserContacts(userId string, contacts *d.UserContacts) {
	r.userContacts[userId] = contacts
}

// userAddresses

func (r *ReadRepository) GetUserAddresses(userId string) *d.UserAddresses {
	return r.userAddresses[userId]
}

func (r *ReadRepository) SetUserAddresses(userId string, addresses *d.UserAddresses) {
	r.userAddresses[userId] = addresses
}
